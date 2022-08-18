package filestorage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var uploader *

type response struct {
	URL string `json:"url"`
}

type Client struct {
	endpoint   string
	httpClient *http.Client
}

func NewClient(endpoint string) *Client {
	client := &http.Client{}

	return &Client{
		httpClient: client,
		endpoint:   endpoint,
	}
}

func (c *Client) getPresignedURL(token string) (string, error) {
	// the line bellow needs go 1.19
	uploadUrl, _ := url.JoinPath(c.endpoint, token)
	req, err := http.NewRequest("GET", uploadUrl, nil)
	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var r response
	err = json.NewDecoder(res.Body).Decode(&r)

	return r.URL, err
}

// Expects a complete file path. It will open and upload the file to s3
// Returns the name/token of the file that was uploaded
func (c *Client) Upload(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	token := fmt.Sprintf("%s_%d", file.Name(), time.Now().Unix())
	signedURL, err := c.getPresignedURL(token)
	if err != nil {
		return "", err
	}

	buff := &bytes.Buffer{}
	w := multipart.NewWriter(buff)
	part, err := w.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPut, signedURL, buff)
	if err != nil {
		return "", err
	}

	fInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fSize := strconv.FormatInt(fInfo.Size(), 10)

	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Set("Content-Type", "application/pdf")
	req.Header.Set("Content-Length", fSize)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("upload request failed with status code %d", res.StatusCode)
	}

	return token, nil
}
