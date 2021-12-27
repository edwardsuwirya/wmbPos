package main

import "github.com/edwardsuwirya/wmbPos/config"

func main() {
	appConfig := config.NewConfig()
	appConfig.RunMigration()
	appConfig.StartEngine()
}
