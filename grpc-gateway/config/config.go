package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server  ServerConfig
	Gateway GatewayConfig
	Redis   RedisConfig
	Limiter Limiter
	Logger  Logger
	Metrics Metrics
	Jaeger  Jaeger
	Minio   Minio
	Casbin  Casbin
}

type ServerConfig struct {
	Name         string
	AppVersion   string
	Port         string
	Mode         string
	JwtSecretKey string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Gateway config struct
type GatewayConfig struct {
	AuthServiceEndpoint   string
	UserServiceEndpoint   string
	SafetyServiceEndpoint string
}

// Redis config
type RedisConfig struct {
	Address      string
	Password     string
	DB           int
	MinIdleConns int
	PoolSize     int
	PoolTimeout  int
}

// Limiter
type Limiter struct {
	Limit string
}

// Logger config
type Logger struct {
	Encoding string
	Level    string
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// Jaeger
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

// Casbin Config
type Casbin struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string
	PostgresTZ       string
	URL              string
}

// Minio config
type Minio struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
	NewBucket       bool
	Bucket          string
	Location        string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	c := new(Config)

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	// customize env
	serverPort := os.Getenv("PORT")
	if len(serverPort) > 0 {
		cfg.Server.Port = serverPort
	}

	serverMode := os.Getenv("MODE")
	if len(serverMode) > 0 {
		cfg.Server.Mode = serverMode
	}

	serverJwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if len(serverJwtSecretKey) > 0 {
		cfg.Server.JwtSecretKey = serverJwtSecretKey
	}

	gatewayAuthServiceEndpoint := os.Getenv("AUTH_SERVICE_ENDPOINT")
	if len(gatewayAuthServiceEndpoint) > 0 {
		cfg.Gateway.AuthServiceEndpoint = gatewayAuthServiceEndpoint
	}

	gatewayUserServiceEndpoint := os.Getenv("USER_SERVICE_ENDPOINT")
	if len(gatewayUserServiceEndpoint) > 0 {
		cfg.Gateway.UserServiceEndpoint = gatewayUserServiceEndpoint
	}

	gatewaySafetyServiceEndpoint := os.Getenv("SAFETY_SERVICE_ENDPOINT")
	if len(gatewaySafetyServiceEndpoint) > 0 {
		cfg.Gateway.SafetyServiceEndpoint = gatewaySafetyServiceEndpoint
	}

	redisAddress := os.Getenv("REDIS_ADDRESS")
	if len(redisAddress) > 0 {
		cfg.Redis.Address = redisAddress
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if len(redisPassword) > 0 {
		cfg.Redis.Password = redisPassword
	}

	redisDB := os.Getenv("REDIS_DB")
	if len(redisDB) > 0 {
		db, err := strconv.Atoi(redisDB)
		if err != nil {
			return nil, err
		}
		cfg.Redis.DB = db
	}

	casbinPostgresHost := os.Getenv("CASBIN_PGHOST")
	if len(casbinPostgresHost) > 0 {
		cfg.Casbin.PostgresHost = casbinPostgresHost
	}

	casbinPostgresPort := os.Getenv("CASBIN_PGPORT")
	if len(casbinPostgresPort) > 0 {
		cfg.Casbin.PostgresPort = casbinPostgresPort
	}

	casbinPostgresUser := os.Getenv("CASBIN_PGUSER")
	if len(casbinPostgresUser) > 0 {
		cfg.Casbin.PostgresUser = casbinPostgresUser
	}

	casbinPostgresPassword := os.Getenv("CASBIN_PGPASSWORD")
	if len(casbinPostgresPassword) > 0 {
		cfg.Casbin.PostgresPassword = casbinPostgresPassword
	}

	casbinPostgresTZ := os.Getenv("CASBIN_PGTZ")
	if len(casbinPostgresTZ) > 0 {
		cfg.Casbin.PostgresTZ = casbinPostgresTZ
	}

	casbinURL := os.Getenv("CASBIN_URL")
	if len(casbinURL) > 0 {
		cfg.Casbin.URL = casbinURL
	}

	metricsURL := os.Getenv("METRICS_URL")
	if len(metricsURL) > 0 {
		cfg.Metrics.URL = metricsURL
	}

	jaegerHost := os.Getenv("JAEGER_HOST")
	if len(jaegerHost) > 0 {
		cfg.Jaeger.Host = jaegerHost
	}

	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	if len(minioEndpoint) > 0 {
		cfg.Minio.Endpoint = minioEndpoint
	}

	minioAccessKeyId := os.Getenv("MINIO_ACCESSKEYID")
	if len(minioAccessKeyId) > 0 {
		cfg.Minio.AccessKeyId = minioAccessKeyId
	}

	minioSecretAccessKey := os.Getenv("MINIO_SECRETACCESSKEY")
	if len(minioSecretAccessKey) > 0 {
		cfg.Minio.SecretAccessKey = minioSecretAccessKey
	}

	minioNewBucket := os.Getenv("MINIO_NEWBUCKET")
	if len(minioNewBucket) > 0 {
		newBucket, err := strconv.ParseBool(minioNewBucket)
		if err != nil {
			return nil, err
		}
		cfg.Minio.NewBucket = newBucket
	}

	minioBucket := os.Getenv("MINIO_BUCKET")
	if len(minioBucket) > 0 {
		cfg.Minio.Bucket = minioBucket
	}

	minioLocation := os.Getenv("MINIO_LOCATION")
	if len(minioLocation) > 0 {
		cfg.Minio.Location = minioLocation
	}

	return cfg, nil
}
