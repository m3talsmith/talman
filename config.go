package main

import (
	"fmt"
	"os"
	"strings"
)

const NAME = "talman"
const VERSION = "0.0.1"

var host, port, url string

func initConfig() {
	host = getEnv("host", "localhost")
	port = getEnv("port", "3000")
	url = fmt.Sprintf("%s:%s", host, port) // combined to make it easier to ListenAndServe
}

// getEnv allows you to fetch a value from the os environment or return a default value
func getEnv(key, defaultValue string) string {
	env := fmt.Sprintf("%s_%s", NAME, strings.ToUpper(key))
	if v := os.Getenv(env); v != "" {
		return v
	}
	return defaultValue
}
