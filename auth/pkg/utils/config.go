package utils

import (
	"fmt"
	"os"
	"strings"
)

// Get config path for local or docker
func GetConfigPath(configPathFromEnv string) string {
	var configPath string
	if configPathFromEnv != "" {
		configPath = configPathFromEnv
	} else {
		getwd, _ := os.Getwd()
		if getwd == "/" {
			getwd = strings.TrimPrefix(getwd, "/")
		}
		configPath = fmt.Sprintf("%s/config/config-local.yml", getwd)
	}

	return configPath
}
