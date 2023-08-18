package handlers

import (
	"github.com/gin-gonic/gin"
	"gitnub.com/artemKapitonov/libraryAPI/internal/service"
)

type Handler struct {
	Service
}

func New(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	books := router.Group("/books", h.userIdentity)
	{
		books.GET("/", h.allBooks)

		books.GET("/:id", h.bookByID)

		books.GET("/my", h.myBooks)

		books.POST("/", h.createBook)

		books.PUT("/:id", h.updateBook)

		books.DELETE("/:id", h.deleteByID)
	}

	return router
}
