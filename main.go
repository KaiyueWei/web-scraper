package main

import (
	"log" // Logging 
	"net/http" // HTTP server
	"os" // Environment Variables

	"github.com/gin-contrib/cors" // CORS middleware for GIN
	"github.com/gin-gonic/gin" // Gin Web framework
	"github.com/joho/godotenv" // Load .env files
)

func main(){
	// Load the environment
	godotenv.Load(".env")
	portString := os.Getenv("PORT") // Get the Port value
	if portString == ""{
		log.Fatal("Port is not found in the environment")
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
	}

	// Create and configure the server
	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server starting on port %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}