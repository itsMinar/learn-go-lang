package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is Required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP Port is Required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP Port must be a number")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}

func GetConfig() Config {
	loadConfig()
	return configurations
}
