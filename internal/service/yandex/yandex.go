package yandex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

type UploadURLresponse struct {
	OperationID string `json:"operation_id"`
	Href        string `json:"href"`
	Method      string `json:"method"`
	Tempalted   bool   `json:"templated"`
}

type GetBookResponse struct {
	Name     string `json:"name"`
	FilePath string `json:"file"`
}

func ConvertBookToBookResp(book *models.Book) models.BookResponse {
	bookResp := models.BookResponse{
		ID:     book.ID,
		Author: book.Author,
		Title:  book.Title,
		Path:   book.Path,
	}

	return bookResp
}

func UploadFileToYandexDisk(book *models.Book, token string, userID int) (string, error) {

	var uploadURLResponse UploadURLresponse

	bookResp := ConvertBookToBookResp(book)

	fileName := GenerateFileName(bookResp, userID)

	url := fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources/upload?path=app:/%s", fileName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &uploadURLResponse); err != nil {
		return "", err
	}

	req, err = http.NewRequest("PUT", uploadURLResponse.Href, book.File)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/binary")

	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	filePath, err := GetFilePath(fileName, token)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func GenerateFileName(book models.BookResponse, userID int) string {
	fileName := fmt.Sprintf("%s__%s__%s.fb2", book.Author, book.Title, strconv.Itoa(userID))
	return fileName
}

func GetFilePath(fileName, token string) (string, error) {
	url := fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources?path=app:/%s", fileName)

	var bookResp GetBookResponse

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bookResp); err != nil {
		return "", err
	}

	return bookResp.FilePath, nil
}

func DeleteFile(book models.BookResponse, token string, userID int) error {

	fileName := GenerateFileName(book, userID)

	url := fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources?path=app:/%s", fileName)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
