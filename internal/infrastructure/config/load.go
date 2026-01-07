package config

import "github.com/spf13/viper"

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.BindEnv("server.api_prefix", "SV_API_PREFIX")
	viper.BindEnv("server.port", "SV_PORT")
	viper.BindEnv("server.write_timeout", "SV_WRITE_TIMEOUT")
	viper.BindEnv("server.read_timeout", "SV_READ_TIMEOUT")
	viper.BindEnv("server.idle_timeout", "SV_IDLE_TIMEOUT")
	viper.BindEnv("server.allow_origins", "SV_ALLOW_ORIGINS")
	viper.BindEnv("server.allow_methods", "SV_ALLOW_METHODS")
	viper.BindEnv("server.allow_headers", "SV_ALLOW_HEADERS")
	viper.BindEnv("server.expose_headers", "SV_EXPOSE_HEADERS")
	viper.BindEnv("server.allow_credentials", "SV_ALLOW_CREDENTIALS")
	viper.BindEnv("server.max_age", "SV_MAX_AGE")
	viper.BindEnv("server.max_header_bytes", "SV_MAX_HEADER_BYTES")

	viper.BindEnv("postgresql.host", "PG_HOST")
	viper.BindEnv("postgresql.port", "PG_PORT")
	viper.BindEnv("postgresql.user", "PG_USER")
	viper.BindEnv("postgresql.password", "PG_PASSWORD")
	viper.BindEnv("postgresql.ssl_mode", "PG_SSL_MODE")
	viper.BindEnv("postgresql.db_name", "PG_DB_NAME")

	viper.BindEnv("log.level", "LOG_LEVEL")
	viper.BindEnv("log.encoding", "LOG_ENCODING")
	viper.BindEnv("log.output_path", "LOG_OUTPUT_PATH")
	viper.BindEnv("log.max_size", "LOG_MAX_SIZE")
	viper.BindEnv("log.max_backups", "LOG_MAX_BACKUPS")
	viper.BindEnv("log.max_age", "LOG_MAX_AGE")
	viper.BindEnv("log.compress", "LOG_COMPRESS")

	viper.BindEnv("minio.endpoint", "MIN_ENDPOINT")
	viper.BindEnv("minio.access_key_id", "MIN_ACCESS_KEY_ID")
	viper.BindEnv("minio.secret_access_key", "MIN_SECRET_ACCESS_KEY")
	viper.BindEnv("minio.bucket", "MIN_BUCKET")
	viper.BindEnv("minio.region", "MIN_REGION")
	viper.BindEnv("minio.public_read", "MIN_PUBLIC_READ")
	viper.BindEnv("minio.use_ssl", "MIN_USE_SSL")

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
