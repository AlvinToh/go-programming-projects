package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/middleware"
	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/model"
	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/app/route"
	"github.com/alvintoh/go-programming-projects/go-bookstore/internal/platform/database"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	setupConfig()
	connectDatabase()
}

func setupConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./resource")
	viper.AddConfigPath("../../resource")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if selectedConfigEnv := os.Getenv("CONFIG_ENV"); selectedConfigEnv != "" {
		viper.SetConfigName(selectedConfigEnv)
		viper.MergeInConfig()
	}
}

func connectDatabase() {
	database.Connect()
	db := database.GetDB()
	model.SetDB(db)

	if err := db.AutoMigrate(&model.Book{}); err != nil {
		panic("Failed to auto migrate Book: " + err.Error())
	}
}

func main() {
	r := mux.NewRouter()
	setupRoutes(r)
	startServer(r)
}

func setupRoutes(r *mux.Router) {
	r.PathPrefix("/swagger/ui/").Handler(http.StripPrefix("/swagger/ui/", http.FileServer(http.Dir("./swagger/ui"))))
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))
	r.Use(middleware.LoggingMiddleware)
	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
}

func startServer(r *mux.Router) {
	addr := viper.GetString("service.host") + ":" + viper.GetString("service.port")
	log.Println("Starting server at " + addr)

	server := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
