package main

import (
	r "ldchat/router"
	"ldchat/utils"
)

func main() {
	utils.SystemInit()
	r := r.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
