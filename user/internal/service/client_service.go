package service

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientService struct {
	baseURL string
}

func NewClientService(baseURL string) *ClientService {
	return &ClientService{baseURL: baseURL}
}

func (s *ClientService) ForwardRequest(c *gin.Context, path string) {
	// Create a new request
	url := s.baseURL + path
	method := c.Request.Method

	// Copy the request body
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Create the new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy headers
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to forward request"})
		return
	}
	defer resp.Body.Close()

	// Copy the response
	for k, v := range resp.Header {
		c.Header(k, v[0])
	}
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
