package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ducktordanny/cubeit/backend/services/auth"
	"github.com/ducktordanny/cubeit/backend/services/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	port string
	db   *sql.DB
}

func NewAPIServer(port string, db *sql.DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

func (server *APIServer) Run() error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			fmt.Sprintf("http://127.0.0.1:%s", server.port),
			fmt.Sprintf("http://localhost:%s", server.port),
		},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        8 * time.Hour,
	}))
	subrouter := router.Group("/api/v1")
	subrouter.GET("", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello CubeIt API!",
		})
	})

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	authHandler := auth.NewHandler(userStore)
	authHandler.RegisterRoutes(subrouter)

	return router.Run(fmt.Sprintf(":%s", server.port))
}
