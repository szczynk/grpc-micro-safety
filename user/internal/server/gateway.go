package server

import (
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func (s *Server) RunGateway() error {
	errCh := make(chan error)

	// ctx, cancel := context.WithCancel(s.ctx)
	// defer cancel()

	//* Register gRPC server endpoint
	//* Note: Make sure the gRPC server is running properly and accessible
	grpcMux := runtime.NewServeMux(
		runtime.WithHealthzEndpoint(s.healthClient),
	)
	// opts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }

	// err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Server.Port, opts)
	// if err != nil {
	// 	return err
	// }
	// err = pb.RegisterRoleServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Server.Port, opts)
	// if err != nil {
	// 	return err
	// }
	// err = pb.RegisterMethodServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Server.Port, opts)
	// if err != nil {
	// 	return err
	// }
	// err = pb.RegisterServiceServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Server.Port, opts)
	// if err != nil {
	// 	return err
	// }
	// err = pb.RegisterPolicyServiceHandlerFromEndpoint(ctx, grpcMux, s.cfg.Server.Port, opts)
	// if err != nil {
	// 	return err
	// }

	mux := http.NewServeMux()
	//* mount the gRPC HTTP gateway to the root
	mux.Handle("/", grpcMux)

	//* mount a path to expose the generated OpenAPI specification on disk
	// mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./doc/swagger/user.swagger.json")
	// })

	//* mount the Swagger UI that uses the OpenAPI specification path above
	// mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./doc/swagger-ui"))))

	server := &http.Server{
		Addr:         s.cfg.Gateway.Port,
		Handler:      mux,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	go func() {
		s.logger.Infof("Gateway available URL: %s", s.cfg.Gateway.Port)
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
