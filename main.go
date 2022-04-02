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

func main() {
	router := gin.Default()
	 
	router.GET("/recipes",getAllRecipes)
	router.GET("/recipe/:id",getRecipe)
	router.Run("localhost:8080")
}