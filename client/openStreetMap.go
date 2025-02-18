package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Udehlee/go-Ride/models"
	"github.com/gin-gonic/gin"
)

type OSMClient struct {
	Client    *http.Client
	UserAgent string
	baseURL   string
}

func NewOSMClient(userAgent, baseURL string) *OSMClient {
	return &OSMClient{
		Client:    &http.Client{},
		UserAgent: userAgent,
		baseURL:   baseURL,
	}
}

// CurrentAddr returns the address of the latitude and longitude provided
func (osm *OSMClient) CurrentAddr(c *gin.Context, lat, lon float64) (*models.OSMResponse, error) {
	osm.Client.Timeout = time.Second * 10
	url := fmt.Sprintf("%s/reverse?lat=%f&lon=%f&format=json", osm.baseURL, lat, lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating request: %v", err),
		})
		return nil, err
	}

	req.Header.Set("User-Agent", osm.UserAgent)

	resp, err := osm.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("request failed: %v", err),
		})
		return nil, err
	}
	defer resp.Body.Close()

	return osm.OSMJsonResponse(c, resp.Body)
}

// OSMJsonResponse decodes the JSON response from OpenStreetMap
func (osm *OSMClient) OSMJsonResponse(c *gin.Context, body io.Reader) (*models.OSMResponse, error) {
	var OSMresp models.OSMResponse

	if err := json.NewDecoder(body).Decode(&OSMresp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to read location details"})
		return nil, err
	}

	return &OSMresp, nil
}
