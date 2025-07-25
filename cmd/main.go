package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gofax-store/internal/config"
	"gofax-store/internal/migrations"
	"gofax-store/pkg/bootstrap"
	"gofax-store/pkg/utils"
	"log"
	"os"
	"time"
)

func init() {
	location, _ := time.LoadLocation("Asia/Tashkent")
	time.Local = location

}

func main() {

	// LOAD ENVIRONMENTS
	LoadEnv()

	// GIN MODE
	gin.SetMode(os.Getenv("GIN_MODE"))

	// GIN DEFAULT
	r := gin.Default()

	//LOGGER
	r.Use(bootstrap.SetupErrorLogger())

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Database connection
	config.DBConnect()
	db := config.GetDB()
	utils.SetDB(db)

	// Database migrations
	if err := migrations.MigrateAll(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Routes
	config.RegisterRoutes(r)

	// Start server
	port := ":" + os.Getenv("PROJECT_PORT")
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env variables...")
		log.Fatal(err)
	}
}
