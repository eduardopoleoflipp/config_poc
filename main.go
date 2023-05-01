package main

import (
	AppConfig "config/poc/app_config"
	PlatformConfig "config/poc/platform_config"
)

func main() {
	config := AppConfig.LoadConfig()
	platformConfig := PlatformConfig.LoadConfig()

	println("\n")
	println("########APP CONFIG########")
	AppConfig.PrintConfig(config)
	println("\n")
	println("########PLATFORM CONFIG########")
	PlatformConfig.PrintConfig(platformConfig)
}
