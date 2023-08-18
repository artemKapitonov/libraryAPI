package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

type Service interface {
	CreateUser(user models.User) (int, error)
	ParseToken(token string) (int, error)
	GenerateToken(username, password string) (string, error)
	CreateBook(book *models.Book, userID int) (int, string, error)
	MyBooks(userID int) ([]models.BookResponse, error)
	AllBooks() ([]models.BookResponse, error)
	DeleteBook(userID, bookID int) error
	BookByID(bookID int) (models.BookResponse, error)
	UpdateBook(book *models.Book, userID, bookID int) error
}

func (h *Handler) myBooks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var books []models.BookResponse

	books, err = h.Service.MyBooks(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)

}

func (h *Handler) allBooks(c *gin.Context) {
	var books []models.BookResponse

	books, err := h.Service.AllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)

}

func (h *Handler) bookByID(c *gin.Context) {

	param := c.Param("id")
	bookID, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	book, err := h.Service.BookByID(bookID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) createBook(c *gin.Context) {
	userID, err := getUserID(c)
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
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unautorized user")
		return
	}

	param := c.Param("id")
	bookID, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
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

	if err := h.Service.UpdateBook(book, userID, bookID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "success",
	})
}

func (h *Handler) deleteByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "unautorized user")
		return
	}

	param := c.Param("id")
	bookID, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.Service.DeleteBook(userID, bookID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "success",
	})
}
