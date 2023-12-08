package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	d "example/go-nba-api/data"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Use the centralized error handler middleware
	router.Use(errorHandlerMiddleware)

	// Use the middleware to set the data in the context
	router.Use(setDataMiddleware)

	// Get the PORT environment variable, or use 8080 as the default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Routes
	router.GET("/v1/nba/:team/:year", getTeamYear)
	router.GET("/v1/nba/:team/:year/roster", getRoster)
	router.GET("/v1/nba/:team/:year/schedule", getSchedule)

	// Run the server
	router.Run(":" + port)
	fmt.Printf("Server is running on :%s\n", port)
}

func readJson() (d.Data, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("allTeamData.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
        return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result d.Data
	{
	}
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

func errorHandlerMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic, log the error, etc.
			fmt.Println("Recovered from panic:", r)
			c.IndentedJSON(http.StatusInternalServerError, "Server Error")
		}
	}()

	c.Next()
}

func setDataMiddleware(c *gin.Context) {
	result, err := readJson()
	if err != nil {
		panic(err)
	}

	// Set the result data in the context
	c.Set("resultData", result)

	c.Next()
}


func getTeamYear(c *gin.Context) {
	team := c.Param("team")
	year := c.Param("year")

	result := c.MustGet("resultData").(d.Data)[team][year]

    if result.Roster.PlayerMap == nil {
		errorMsg := fmt.Sprintf("Team Data was not found for the %s - %s", year, team)
		fmt.Println(errorMsg)
		c.IndentedJSON(http.StatusNotFound, errorMsg)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func getRoster(c *gin.Context) {
	team := c.Param("team")
	year := c.Param("year")

	result := c.MustGet("resultData").(d.Data)[team][year].Roster

	if result.PlayerMap == nil {
		errorMsg := fmt.Sprintf("Roster Data was not found for the %s - %s", year, team)
		fmt.Println(errorMsg)
		c.IndentedJSON(http.StatusNotFound, errorMsg)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func getSchedule(c *gin.Context) {
	team := c.Param("team")
	year := c.Param("year")

	result := c.MustGet("resultData").(d.Data)[team][year].Schedule

    if result.GameMap == nil {
		errorMsg := fmt.Sprintf("Schedule Data was not found for the %s - %s", year, team)
		fmt.Println(errorMsg)
		c.IndentedJSON(http.StatusNotFound, errorMsg)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
