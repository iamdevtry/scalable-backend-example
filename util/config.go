package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read bt viper form a config file or environment variables
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	//s3 config
	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
	S3Region     string `mapstructure:"S3_REGION"`
	S3APIKey     string `mapstructure:"S3_API_KEY"`
	S3Secret     string `mapstructure:"S3_SECRET"`
	S3Domain     string `mapstructure:"S3_DOMAIN"`
	SysSecretKey string `mapstructure:"SYSTEM_SECRET_KEY"`
}

// LoadConfig reads configurations from file or enviroment varivable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
