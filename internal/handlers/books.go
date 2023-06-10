package handlers

import (
	"fmt"
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
	fmt.Println(userID)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unautorized user")
		return
	}

	var input models.Book

	input.Author = c.Request.FormValue("author")

	input.Title = c.Request.FormValue("title")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Can't get file from form: %s", err.Error()))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Can't open file: %s", err.Error()))
		return
	}

	input.File = file

	book := &input

	id, href, err := h.Service.CreateBook(book, userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Can't create book: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"book": href,
	})
}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteByID(c *gin.Context) {

}
