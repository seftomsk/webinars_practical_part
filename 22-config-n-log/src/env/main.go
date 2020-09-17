package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
	Host string
}

func main() {
	httpPort, err := strconv.Atoi(os.Getenv("SHORTENER_PORT"))
	if err != nil {
		log.Fatal("SHORTENER_PORT is not defined")
	}

	shortenerHost := os.Getenv("SHORTENER_HOST")
	if shortenerHost == "" {
		log.Fatal("SHORTENER_HOST is not defined")
	}

	config := Config{httpPort, shortenerHost}
	fmt.Printf("%+v\n", config)
}
