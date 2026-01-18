package config

import "time"

type ServerConfig struct {
	APIPrefix        string        `mapstructure:"api_prefix"`
	Port             int           `mapstructure:"port"`
	WriteTimeout     time.Duration `mapstructure:"write_timeout"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout"`
	IdleTimeout      time.Duration `mapstructure:"idle_timeout"`
	MaxHeaderBytes   int           `mapstructure:"max_header_bytes"`
	AllowOrigins     []string      `mapstructure:"allow_origins"`
	AllowMethods     []string      `mapstructure:"allow_methods"`
	AllowHeaders     []string      `mapstructure:"allow_headers"`
	ExposeHeaders    []string      `mapstructure:"expose_headers"`
	AllowCredentials bool          `mapstructure:"allow_credentials"`
	MaxAge           time.Duration `mapstructure:"max_age"`
}

type JWTConfig struct {
	AccessName       string        `mapstructure:"access_name"`
	RefreshName      string        `mapstructure:"refresh_name"`
	GuestName        string        `mapstructure:"guest_name"`
	SecretKey        string        `mapstructure:"secret_key"`
	AccessExpiresIn  time.Duration `mapstructure:"access_expires_in"`
	RefreshExpiresIn time.Duration `mapstructure:"refresh_expires_in"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Encoding   string `mapstructure:"encoding"`
	OutputPath string `mapstructure:"output_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type PostgreSQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl_mode"`
	DBName   string `mapstructure:"db_name"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	UseSSL   bool   `mapstructure:"use_ssl"`
}

type MinIOConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Region          string `mapstructure:"region"`
	Bucket          string `mapstructure:"bucket"`
	PublicRead      bool   `mapstructure:"public_read"`
	UseSSL          bool   `mapstructure:"use_ssl"`
	PublicDomain    string `mapstructure:"public_domain"`
}

type SuperUserConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Vhost    string `mapstructure:"vhost"`
	UseSSL   bool   `mapstructure:"use_ssl"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Log        LogConfig        `mapstructure:"log"`
	PostgreSQL PostgreSQLConfig `mapstructure:"postgresql"`
	Redis      RedisConfig      `mapstructure:"redis"`
	MinIO      MinIOConfig      `mapstructure:"minio"`
	SuperUser  SuperUserConfig  `mapstructure:"super_user"`
	RabbitMQ   RabbitMQ         `mapstructure:"rabbitmq"`
	SMTPConfig SMTPConfig       `mapstructure:"smtp"`
}
