package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	apiKeys map[string]string // Map to store client name to API key
}

func NewMiddleware() (Middleware, error) {
	//todo load file from env
	apiKeys, err := loadAPIKeys(os.Getenv("KEYS_FILE_PATH"))
	if err != nil {
		return Middleware{}, err
	}
	return Middleware{apiKeys: apiKeys}, nil
}

func loadAPIKeys(configFile string) (map[string]string, error) {
	// Read the entire file content
	byteValue, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	// Parse JSON data
	var config struct {
		Clients map[string]struct {
			APIKey string `json:"apiKey"`
		} `json:"clients"`
	}

	if err := json.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}

	// Transform into a simpler map[string]string for easier access
	apiKeys := make(map[string]string)
	for clientName, clientData := range config.Clients {
		apiKeys[clientName] = clientData.APIKey
	}

	return apiKeys, nil
}

func (mw Middleware) ApiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientName := c.GetHeader("Client-Name")
		apiKey := c.GetHeader("X-API-KEY")

		// Check if clientName and apiKey are provided
		if clientName == "" || apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Client name and API key required"})
			c.Abort()
			return
		}

		// Check if the client name exists in the loaded API keys
		expectedApiKey, exists := mw.apiKeys[clientName]
		if !exists || apiKey != expectedApiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid client name or API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
