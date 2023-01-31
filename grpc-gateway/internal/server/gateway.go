package server

import (
	"context"
	"encoding/json"
	"gateway/pb"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) RunGateway() error {
	errCh := make(chan error)

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	//* Register gRPC server endpoint
	//* Note: Make sure the gRPC server is running properly and accessible
	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.AuthServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterRoleServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterMethodServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterServiceServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterPolicyServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.UserServiceEndpoint, opts)
	if err != nil {
		return err
	}
	// TODO(Bagus): register safety service
	err = pb.RegisterAttendanceServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterCertificateServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterCheckServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterOfficeServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterScheduleServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pb.RegisterWorkspaceServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Gateway.SafetyServiceEndpoint, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	//* mount the gRPC HTTP gateway to the root
	mux.Handle("/", grpcMux)

	if s.cfg.Server.Mode != "production" {
		//* mount a path to expose the generated OpenAPI specification on disk
		mux.HandleFunc(
			"/swagger-ui/swagger.json",
			func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "./doc/swagger/gateway.swagger.json")
			},
		)

		//* mount the Swagger UI that uses the OpenAPI specification path above
		mux.Handle(
			"/swagger-ui/",
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./doc/swagger-ui"))),
		)
	}

	//* Upload Image
	mux.HandleFunc(
		"/images",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// ============= Middleware Start =============
			// ============= RateLimiter Start =============
			key := s.limiter.Limiter().GetIPKey(r)

			limiterContext, err := s.limiter.Limiter().Get(ctx, key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Add("X-RateLimit-Limit", strconv.FormatInt(limiterContext.Limit, 10))
			w.Header().Add("X-RateLimit-Remaining", strconv.FormatInt(limiterContext.Remaining, 10))
			w.Header().Add("X-RateLimit-Reset", time.Unix(limiterContext.Reset, 0).Format(time.RFC3339))

			if limiterContext.Reached {
				http.Error(
					w,
					"The method of UploadImage's limit has been exceeded. Please try again later.",
					http.StatusTooManyRequests,
				)
				return
			}
			// ============= RateLimiter End =============
			// ============= Auth Start =============
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Request unauthenticated with Bearer", http.StatusUnauthorized)
				return
			}
			splits := strings.SplitN(authHeader, " ", 2)
			if len(splits) < 2 {
				http.Error(w, "Bad authorization string", http.StatusUnauthorized)
				return
			}
			if (!strings.EqualFold(splits[0], "Bearer")) || (!strings.EqualFold(splits[0], "bearer")) {
				http.Error(w, "Request unauthenticated with Bearer", http.StatusUnauthorized)
				return
			}

			accessToken := splits[1]
			isExpire := true
			payload, err := s.tokenMaker.VerifyToken(accessToken, isExpire)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			userId := payload.Get("user_id")
			role := payload.Get("role")
			service, method := "pb.ImageService", "UploadImage"

			allowed, err := s.casbin.Enforce(s.ctx, role, service, method)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "Method not allowed", http.StatusUnauthorized)
				return
			}
			// ============= Auth End =============
			// ============= Middleware End =============

			span, ctx := opentracing.StartSpanFromContext(s.ctx, "image.UploadImage")
			defer span.Finish()

			// Maximum upload of 2 MB images
			// r.Body = http.MaxBytesReader(w, r.Body, 2<<20)
			if err := r.ParseMultipartForm(2 << 20); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			uploadedImage, handler, err := r.FormFile("image")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer uploadedImage.Close()

			objectName := "images/" + uuid.New().String() + "-" + url.PathEscape(handler.Filename)
			fileSize := handler.Size
			mimeType := handler.Header.Get("Content-Type")

			mime := strings.Split(mimeType, "/")
			if mime[0] != "image" {
				http.Error(w, "The format file is not valid.", http.StatusBadRequest)
				return
			}

			info, err := s.minio.PutObject(
				ctx,
				s.cfg.Minio.Bucket,
				objectName,
				uploadedImage,
				fileSize,
				minio.PutObjectOptions{
					ContentType: mimeType,
					UserMetadata: map[string]string{
						"user_id": userId,
						"role":    role,
					},
				},
			)
			if err != nil {
				s.logger.Errorf("minio.PutObject: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(info)
			if err != nil {
				s.logger.Errorf("json.NewEncoder: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	//* Get Image
	mux.HandleFunc(
		"/images/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// ============= Middleware Start =============
			// ============= RateLimiter Start =============
			key := s.limiter.Limiter().GetIPKey(r)

			limiterContext, err := s.limiter.Limiter().Get(ctx, key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Add("X-RateLimit-Limit", strconv.FormatInt(limiterContext.Limit, 10))
			w.Header().Add("X-RateLimit-Remaining", strconv.FormatInt(limiterContext.Remaining, 10))
			w.Header().Add("X-RateLimit-Reset", time.Unix(limiterContext.Reset, 0).Format(time.RFC3339))

			if limiterContext.Reached {
				http.Error(
					w,
					"The method of UploadImage's limit has been exceeded. Please try again later.",
					http.StatusTooManyRequests,
				)
				return
			}
			// ============= RateLimiter End =============
			// ============= Auth Start =============
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Request unauthenticated with Bearer", http.StatusUnauthorized)
				return
			}
			splits := strings.SplitN(authHeader, " ", 2)
			if len(splits) < 2 {
				http.Error(w, "Bad authorization string", http.StatusUnauthorized)
				return
			}
			if (!strings.EqualFold(splits[0], "Bearer")) || (!strings.EqualFold(splits[0], "bearer")) {
				http.Error(w, "Request unauthenticated with Bearer", http.StatusUnauthorized)
				return
			}

			accessToken := splits[1]
			isExpire := true
			payload, err := s.tokenMaker.VerifyToken(accessToken, isExpire)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			role := payload.Get("role")
			service, method := "pb.ImageService", "GetImageId"

			allowed, err := s.casbin.Enforce(s.ctx, role, service, method)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "Method not allowed", http.StatusUnauthorized)
				return
			}
			// ============= Auth End =============
			// ============= Middleware End =============

			span, ctx := opentracing.StartSpanFromContext(s.ctx, "image.GetImageId")
			defer span.Finish()

			id := strings.TrimPrefix(r.URL.Path, "/images/")
			if id == "" {
				http.Error(w, "No image name, invalid request.", http.StatusBadRequest)
				return
			}

			id = url.PathEscape(id)
			objectName := strings.Join([]string{"images", id}, "/")

			object, err := s.minio.GetObject(
				ctx,
				s.cfg.Minio.Bucket,
				objectName,
				minio.GetObjectOptions{},
			)
			if err != nil {
				s.logger.Errorf("minio.GetObject: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer object.Close()

			objInfo, err := object.Stat()
			if err != nil {
				s.logger.Errorf("minio.GetObject.Stat: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for k, v := range objInfo.Metadata {
				w.Header().Add(k, v[0])
			}

			w.Header().Add("Accept-Ranges", "bytes")
			w.Header().Add("ETag", objInfo.ETag)
			w.Header().Add("Content-Length", strconv.FormatInt(objInfo.Size, 10))
			w.Header().Add("Last-Modified", objInfo.LastModified.Format("Thu, 19 Jan 2023 15:35:58 GMT"))

			http.ServeContent(w, r, objInfo.Key, time.Now(), object)
		},
	)

	handler := s.cors.Handler(mux)

	server := &http.Server{
		Addr:         s.cfg.Server.Port,
		Handler:      handler,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	go func() {
		s.logger.Infof("Gateway available URL: %s", s.cfg.Server.Port)
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-s.ctx.Done():
		c := make(chan bool)
		go func() {
			defer close(c)
			errCh <- server.Shutdown(s.ctx)
		}()
		select {
		case <-c:
		case <-time.After(5 * time.Second):
		}
		s.logger.Info("Gateway Exited Properly")
		return nil
	case err := <-errCh:
		return err
	}
}
