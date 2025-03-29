package main

import (
	"instant_messaging/router"
	"instant_messaging/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081")

}
