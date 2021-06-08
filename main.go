package main

import (
	entry "YOUR_APP_NAME.YOUR_MODULE_NAME/YOUR_FUNCTION_NAME"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}