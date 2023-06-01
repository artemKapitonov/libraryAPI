package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	token = "y0_AgAAAABkKL5FAAn-dAAAAADkcRkYSFksQS8jThWFdB0hQqqUwKa1Kjo"
)

func uploadFileToYandexDisk(filePath string, destinationFolder string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(file)
	defer file.Close()
	//https://cloud-api.yandex.net/v1/disk/resources/upload? path=
	url := fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources/upload?path=/%s/%s", destinationFolder, file.Name())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	uploadURL := string(body)
	req, err = http.NewRequest("PUT", uploadURL, file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 201
}

func main() {
	filePath := "/home/kapitonov/go-work/src/github.com/artemKapitonov/libraryAPI/internal/api/hello.txt"
	destinationFolder := "/Library"
	if uploadFileToYandexDisk(filePath, destinationFolder) {
		fmt.Println("File uploaded successfully")
	} else {
		fmt.Println("Error uploading file")
	}
}
