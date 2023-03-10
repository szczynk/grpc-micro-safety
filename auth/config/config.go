package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server       ServerConfig
	Postgres     PostgresConfig
	Redis        RedisConfig
	Kafka        KafkaConfig
	RefreshToken RefreshToken
	AccessToken  AccessToken
	Logger       Logger
	Gateway      Gateway
	Casbin       Casbin
	Metrics      Metrics
	Jaeger       Jaeger
	Limiter      Limiter
}

// Server config struct
type ServerConfig struct {
	Name              string
	AppVersion        string
	Port              string
	Mode              string
	JwtSecretKey      string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

// Logger config
type Logger struct {
	Encoding string
	Level    string
}

// Postgres config
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TZ       string
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

// Kafka config
type KafkaConfig struct {
	Brokers []string
	GroupID string
}

// RefreshToken config
type RefreshToken struct {
	MaxAge time.Duration
}

// AccessToken config
type AccessToken struct {
	Expire time.Duration
}

// Server config struct
type Gateway struct {
	Port string
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

	postgresHost := os.Getenv("PGHOST")
	if len(postgresHost) > 0 {
		cfg.Postgres.Host = postgresHost
	}

	postgresPort := os.Getenv("PGPORT")
	if len(postgresPort) > 0 {
		cfg.Postgres.Port = postgresPort
	}

	postgresUser := os.Getenv("PGUSER")
	if len(postgresUser) > 0 {
		cfg.Postgres.User = postgresUser
	}

	postgresPassword := os.Getenv("PGPASSWORD")
	if len(postgresPassword) > 0 {
		cfg.Postgres.Password = postgresPassword
	}

	postgresDBName := os.Getenv("PGDBNAME")
	if len(postgresDBName) > 0 {
		cfg.Postgres.DBName = postgresDBName
	}

	postgresTZ := os.Getenv("PGTZ")
	if len(postgresTZ) > 0 {
		cfg.Postgres.TZ = postgresTZ
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

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if len(kafkaBrokers) > 0 {
		brokers := strings.Split(kafkaBrokers, ",")
		cfg.Kafka.Brokers = brokers
	}

	return cfg, nil
}
