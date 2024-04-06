package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// createAd now stores the advertisement in Redis
func createAd(c *gin.Context) {
    var newAd Advertisement
    if err := c.BindJSON(&newAd); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Serialize the Advertisement object to JSON
    adJSON, err := json.Marshal(newAd)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize ad"})
        return
    }

    // Save the JSON string in Redis
    err = rdb.Set(ctx, "ad:"+newAd.Title, adJSON, 0).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save ad in Redis"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "created"})
}

// listAds retrieves a list of advertisements from Redis
func listAds(c *gin.Context) {
    // Parse query parameters
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	age, _ := strconv.Atoi(c.Query("age"))
	gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	// Fetch all advertisement keys (consider optimization for large datasets)
	keys, err := rdb.Keys(ctx, "ad:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch keys from Redis"})
		return
	}

    var ads []Advertisement
    for _, key := range keys {
        adJSON, err := rdb.Get(ctx, key).Result()
        if err != nil {
            continue // Skip on error
        }

        var ad Advertisement
        err = json.Unmarshal([]byte(adJSON), &ad)
        if err != nil {
            continue // Skip on error
        }
        ads = append(ads, ad)
    }

    // Filter the advertisements based on the query parameters
    filteredAds := FilterAds(ads, offset, limit, age, gender, country, platform)

	c.JSON(http.StatusOK, gin.H{"items": filteredAds})
}
