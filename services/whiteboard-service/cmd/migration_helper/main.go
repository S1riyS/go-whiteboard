package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/S1riyS/go-whiteboard/whiteboard-service/internal/config"
)

// main is a helper to write database DSN from .yaml config to file
// provided by command line flag `env-dsn-path`.
//
// It is used in Taskfile to work with migrations.
func main() {
	// Read path to config
	var configPath string
	flag.StringVar(&configPath, "config-path", "", "path to config file")

	// Read path to store .env.tmp
	var envDsnPath string
	flag.StringVar(&envDsnPath, "env-dsn-path", "", "path to store file")
	flag.Parse()

	// Load DSN from config
	cfg := config.MustLoad(configPath)
	dsn := cfg.Database.DSN()
	_ = dsn

	// Create .env.tmp
	f, err := os.Create(envDsnPath)
	if err != nil {
		log.Fatalf("cannot create .env.dsn file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("DSN=%s\n", dsn)); err != nil {
		log.Fatalf("cannot write to .env.dsn file: %v", err)
	}
}
