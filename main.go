package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper" //untuk baca env
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DbConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv() //baca dari env secara otomatis
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DbConn: viper.GetString("DB_CONN"),
	}

	//setup db
	db, err := database.InitDB(config.DbConn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	//setup routes
	http.HandleFunc("/api/products", productHandler.HandleProducts)

	// /health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Starting server on :" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
