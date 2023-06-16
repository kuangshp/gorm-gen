package main

import (
	"gorm_gen/initialize"
	_ "gorm_gen/utils"
)

func main() {
	router := initialize.Routers()
	router.Run(":9000")
}
