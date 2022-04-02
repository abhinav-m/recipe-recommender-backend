package main

import (
	"net/http"
	"recipe-recommender-backend/data"

	"github.com/gin-gonic/gin"
)

/*
gin.Context is the most important part of Gin.
 It carries request details, validates and serializes JSON, and more.
*/
func getAllRecipes(c *gin.Context) {
	getRecipeData := data.CsvReaderClosure("/Users/abhinav-m/Work/recipe-recommender-backend/data/all_recipes.csv")
	recipes := getRecipeData()
	c.IndentedJSON(http.StatusOK,recipes)
}

func getRecipe(c *gin.Context) {
	getRecipeData := data.CsvReaderClosure("/Users/abhinav-m/Work/recipe-recommender-backend/data/all_recipes.csv")
	recipes := getRecipeData()

	id := c.Param("id")

	for _, r := range recipes {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK,r)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Recipe not found"})
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())
	 
	router.GET("/recipes",getAllRecipes)
	router.GET("/recipe/:id",getRecipe)
	router.Run("localhost:8080")
}