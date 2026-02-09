package main

import (
	"database/sql"
	"log"      // Logging
	"net/http" // HTTP server
	"os"       // Environment Variables

	"github.com/KaiyueWei/rssagg/internal/database"
	"github.com/gin-contrib/cors" // CORS middleware for GIN
	"github.com/gin-gonic/gin"    // Gin Web framework
	"github.com/joho/godotenv"    // Load .env files
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load the environment
	godotenv.Load(".env")

	// Port
	portString := os.Getenv("PORT") // Get the Port value
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	// Database
	dbURL := os.Getenv("DB_URL") // Get the Database URL
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	queries := database.New(conn)
	apiCfg := apiConfig{
		DB: queries,
	}

	// router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := router.Group("/v1")
	{
		v1.GET("/healthz", gin.WrapF(handlerReadiness))
		v1.GET("/err", gin.WrapF(handlerErr))
		v1.GET("/users", gin.WrapF(apiCfg.middlewareAuth(apiCfg.handlerGetUser)))
		v1.POST("/users", gin.WrapF(apiCfg.handlerCreateUser))
		v1.POST("/feeds", gin.WrapF(apiCfg.middlewareAuth(apiCfg.handlerCreateFeed)))
		v1.GET("/feeds", gin.WrapF(apiCfg.handlerGetFeeds))
	}

	// Create and configure the server
	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server starting on port %v", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
