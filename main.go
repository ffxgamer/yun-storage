package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Database struct {
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	DBFile      string `json:"db_file"`
	TablePrefix string `json:"table_prefix"`
	SslMode     string `json:"ssl_mode"`
}

type Scheme struct {
	Https    bool   `json:"https"`
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}

type Log struct {
	Enable     bool   `json:"enable"`
	Name       string `json:"name"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

type Config struct {
	Force          bool     `json:"force"`
	Address        string   `json:"address"`
	Port           int      `json:"port"`
	JwtSecret      string   `json:"jwt_secret"`
	TokenExpiresIn int      `json:"token_expires_in"`
	SiteUrl        string   `json:"site_url"`
	Cdn            string   `json:"cdn"`
	Database       Database `json:"database"`
	Scheme         Scheme   `json:"scheme"`
	TempDir        string   `json:"temp_dir"`
	Log            Log      `json:"log"`
}

func main() {
	dbType := os.Getenv("DB_TYPE")
	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbTablePrefix := os.Getenv("DB_TABLE_PREFIX")
	dbSslMode := os.Getenv("DB_SSL_MODE")
	configCdn := os.Getenv("CDN")
	configPortStr := os.Getenv("PORT")
	configSiteUrl := os.Getenv("SITE_URL")

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatalf("failed to parse DB_PORT: %v", err)
	}

	configPort, err := strconv.Atoi(configPortStr)
	if err != nil {
		log.Fatalf("failed to parse PORT: %v", err)
	}

	config := Config{
		Force:          false,
		Address:        "0.0.0.0",
		Port:           configPort,
		JwtSecret:      "random generated",
		TokenExpiresIn: 48,
		SiteUrl:        configSiteUrl,
		Cdn:            configCdn,
		Database:       Database{Type: dbType, Host: dbHost, Port: dbPort, User: dbUser, Password: dbPassword, Name: dbName, TablePrefix: dbTablePrefix, SslMode: dbSslMode},
		Scheme:         Scheme{Https: false, CertFile: "", KeyFile: ""},
		TempDir:        "data/temp",
		Log:            Log{Enable: true, Name: "log/log.log", MaxSize: 10, MaxBackups: 5, MaxAge: 28, Compress: false},
	}

	configBody, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}

	err = os.WriteFile("/opt/alist/data/config.json", configBody, 0644)
	if err != nil {
		log.Fatalf("failed to write config file: %v", err)
	}

	fmt.Println("config file generated successfully")
}
