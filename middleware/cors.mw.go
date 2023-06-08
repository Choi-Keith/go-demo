package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Origin, X-Requested-With, Content-Type, Content-Length, Accept, Authorization"},
		AllowCredentials: true,
		MaxAge:           7200,
	})
}
