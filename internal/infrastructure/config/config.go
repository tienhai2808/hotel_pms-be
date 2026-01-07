package config

import "time"

type Config struct {
	Server struct {
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
	} `mapstructure:"server"`

	Log struct {
		Level      string `mapstructure:"level"`
		Encoding   string `mapstructure:"encoding"`
		OutputPath string `mapstructure:"output_path"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAge     int    `mapstructure:"max_age"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"log"`

	PostgreSQL struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"ssl_mode"`
		DBName   string `mapstructure:"db_name"`
	} `mapstructure:"postgresql"`

	MinIO struct {
		Endpoint        string `mapstructure:"endpoint"`
		AccessKeyID     string `mapstructure:"access_key_id"`
		SecretAccessKey string `mapstructure:"secret_access_key"`
		Region          string `mapstructure:"region"`
		Bucket          string `mapstructure:"bucket"`
		PublicRead      bool   `mapstructure:"public_read"`
		UseSSL          bool   `mapstructure:"use_ssl"`
	} `mapstructure:"minio"`
}
