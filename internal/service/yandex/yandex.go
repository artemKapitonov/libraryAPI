package yandex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

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

func UploadFileToYandexDisk(book *models.Book, token string) (string, error) {

	var uploadURLResponse UploadURLresponse

	fileName := GenerateFileName(book)

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

func GenerateFileName(book *models.Book) string {
	currentTime := time.Now().Unix()

	fileName := fmt.Sprintf("%s__%s__%s.fb2", book.Author, book.Title, strconv.Itoa(int(currentTime)))
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

func DeleteBook(fileName, token string) error {
	if fileName == "" {
		return errors.New("Null name of file")
	}

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
