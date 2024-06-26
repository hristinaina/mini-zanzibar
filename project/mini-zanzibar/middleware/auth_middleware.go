package middleware

import (
	"encoding/json"
	"io/ioutil"
	"mini-zanzibar/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	apiKeys    map[string]string // Map to store client name to API key
	logService *services.LogService
}

func NewMiddleware(logService *services.LogService) (Middleware, error) {
	// Initialize log service
	mw := Middleware{
		logService: logService,
	}

	// Load API keys from file
	apiKeys, err := loadAPIKeys(os.Getenv("KEYS_FILE_PATH"))
	if err != nil {
		mw.logService.Error("Failed to load API keys from file: " + err.Error())
		return Middleware{}, err
	}
	mw.apiKeys = apiKeys

	mw.logService.Info("Middleware initialized successfully")
	return mw, nil
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
		mw.logService.Info("Incoming request for API key authentication")

		clientName := c.GetHeader("Client-Name")
		apiKey := c.GetHeader("X-API-KEY")

		// Check if clientName and apiKey are provided
		if clientName == "" || apiKey == "" {
			mw.logService.Error("Client name or API key not provided")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Client name and API key required"})
			c.Abort()
			return
		}

		// Check if the client name exists in the loaded API keys
		expectedApiKey, exists := mw.apiKeys[clientName]
		if !exists || apiKey != expectedApiKey {
			mw.logService.Error("Invalid client name or API key")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid client name or API key"})
			c.Abort()
			return
		}

		mw.logService.Info("API key authentication successful")
		c.Next()
	}
}
