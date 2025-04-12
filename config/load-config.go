package config

import (
	"log/slog"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func loadConfig() error {
	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		slog.Warn(".env not found,that's okay!")
	}

	viper.AutomaticEnv()

	config = &Config{
		Version:     viper.GetString("VERSION"),
		Mode:        Mode(viper.GetString("MODE")),
		ServiceName: viper.GetString("SERVICE_NAME"),
		HttpPort:    viper.GetInt("HTTP_PORT"),
		//HealthCheckRoute: viper.GetString("HEALTH_CHECK_ROUTE"),
		//ReadRedisURL:     viper.GetString("READ_REDIS_URL"),
		//WriteRedisURL:    viper.GetString("WRITE_REDIS_URL"),
		// CacheExpirationTimeInSecond: viper.GetInt("CACHE_EXPIRATION_TIME_IN_SECOND"),
		//EnableRedisTLSMode: viper.GetBool("ENABLE_REDIS_TLS_MODE"),
		// JwtSecret:                   viper.GetString("JWT_SECRET"),
		//MongoReadDB: &MongoReadDB{
		// 	MONGO_URI:           viper.GetString("MONGO_URI"),
		// 	Database:            viper.GetString("MONGO_DATABASE"),
		// 	User:                viper.GetString("MONGO_USER"),
		// 	Password:            viper.GetString("MONGO_PASSWORD"),
		// 	ReplicaSet:          viper.GetString("MONGO_REPLICA_SET"),
		// 	AuthSource:          viper.GetString("MONGO_AUTH_SOURCE"),
		// 	SSL:                 viper.GetBool("MONGO_SSL"),
		// 	MaxPoolSize:         viper.GetUint64("MONGO_MAX_POOL_SIZE"),
		// 	MinPoolSize:         viper.GetUint64("MONGO_MIN_POOL_SIZE"),
		// 	MaxConnIdleTimeInMs: viper.GetInt("MONGO_MAX_CONN_IDLE_TIME_MS"),
		// },
		// MongoWriteDB: &MongoWriteDB{
		// 	MONGO_URI:           viper.GetString("MONGO_URI"),
		// 	Database:            viper.GetString("MONGO_DATABASE"),
		// 	User:                viper.GetString("MONGO_USER"),
		// 	Password:            viper.GetString("MONGO_PASSWORD"),
		// 	ReplicaSet:          viper.GetString("MONGO_REPLICA_SET"),
		// 	AuthSource:          viper.GetString("MONGO_AUTH_SOURCE"),
		// 	SSL:                 viper.GetBool("MONGO_SSL"),
		// 	MaxPoolSize:         viper.GetUint64("MONGO_MAX_POOL_SIZE"),
		// 	MinPoolSize:         viper.GetUint64("MONGO_MIN_POOL_SIZE"),
		// 	MaxConnIdleTimeInMs: viper.GetInt("MONGO_MAX_CONN_IDLE_TIME_MS"),
		// },

		// Apm: &Apm{
		// 	ServiceName: viper.GetString("APM_SERVICE_NAME"),
		// 	ServerURL:   viper.GetString("APM_SERVER_URL"),
		// 	SecretToken: viper.GetString("APM_SECRET_TOKEN"),
		// 	Environment: viper.GetString("APM_ENVIRONMENT"),
		// },
	}

	v := validator.New()
	if err = v.Struct(config); err != nil {
		exit(err)
	}

	return nil
}
