package util

import "github.com/spf13/viper"

// Config stores all the cofigurations of the application
// The values are read by Viper from a config file or env variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

/**
 * LoadConfig loads configuration from the specified path using viper.
 * It sets the configuration file name to "app" and type to "env" (json, xml).
 * It automatically loads environment variables.
 * If successful, it unmarshals the configuration into the provided struct pointer.
 * @param path The path to the configuration file.
 * @return config The loaded configuration.
 * @return err Any error that occurred during the loading process.
 */
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
