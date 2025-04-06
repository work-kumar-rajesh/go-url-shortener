package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service"
)

// URLHandler holds the URLService dependency.
type URLHandler struct {
	urlService service.URLService
}

// NewURLHandler creates a new URLHandler with the given URLService.
func NewURLHandler(s service.URLService) *URLHandler {
	return &URLHandler{urlService: s}
}

// Shorten handles POST requests to shorten URLs.
func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	// Validate input
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Call the service to shorten the URL
	shortCode, err := h.urlService.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not shorten URL"})
		return
	}

	// Return the shortened URL
	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortCode})
}

// Redirect handles GET /:code requests.
func (h *URLHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("code")
	originalURL, err := h.urlService.ResolveURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	// Redirect to the original URL
	c.Redirect(http.StatusFound, originalURL)
}
