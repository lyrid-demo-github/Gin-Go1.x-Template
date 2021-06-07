package main

import (
	entry "$APP_NAME.$MODULE_NAME/$FUNCTION_NAME"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}