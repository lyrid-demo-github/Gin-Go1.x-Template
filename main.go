package main

import (
	"github.com/joho/godotenv"
	entry "go1x_gin.template/entry"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}
