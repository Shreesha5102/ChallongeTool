package configs

import (
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/constants"
	"github.com/spf13/viper"
)

const (
	httpPort = "HTTP_PORT"
)

func init() {
	viper.SetConfigName("ChallongeTool")
	viper.SetConfigType("json")

	viper.SetDefault(httpPort, constants.HttpPort)

	viper.AutomaticEnv()
}

// Func to return the http Port
func GetHTTPPort() string {
	return viper.GetString(httpPort)
}
