package main

import (
	"go-gpt/config"
	"go-gpt/internal/router"
)

func main() {
	r := router.SetRouters()
	err := r.Run(":" + config.GLOBAL_CONFIG.PORT)
	if err != nil {
		panic(err)
	}
}
