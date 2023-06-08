package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

func (h *Handler) allBooks(c *gin.Context) {
}

func (h *Handler) bookByID(c *gin.Context) {

}

func (h *Handler) createBook(c *gin.Context) {

	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unautorized user")
	}

	var input models.Book

	input.Author = c.Request.FormValue("author")

	input.Title = c.Request.FormValue("title")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	input.File = file

	h.Service.CreateBook(input, userID)
}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteByID(c *gin.Context) {

}
