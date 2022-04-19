package main

import (
	"net/http"
	"os"
	"recipe-recommender-backend/data"

	"github.com/gin-gonic/gin"
)

/*
gin.Context is the most important part of Gin.
 It carries request details, validates and serializes JSON, and more.
*/
func getAllRecipes(c *gin.Context) {
	getRecipeData := data.JsonReader("/Users/abhinav-m/Work/recipe-recommender-backend/all_recipes.json")
	recipes := getRecipeData()
	c.IndentedJSON(http.StatusOK,recipes)
}

func getAllRecommendations(c *gin.Context) {
	getRecipeData := data.JsonReader("/Users/abhinav-m/Work/recipe-recommender-backend/recommendation_data.json")
	recipeRecommendations := getRecipeData()
	c.IndentedJSON(http.StatusOK,recipeRecommendations)

}

func getRecommendation(c *gin.Context) {
	getRecipeData := data.JsonReader("/Users/abhinav-m/Work/recipe-recommender-backend/recommendation_data.json")
	recipeRecommendations := getRecipeData()
	
	id := c.Param("id")

	if recipe,ok := recipeRecommendations[id]; ok {
		c.IndentedJSON(http.StatusOK,recipe)
		return 
	}
	
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Recipe Recommendation not found"})
}

func getRecipe(c *gin.Context) {
	getRecipeData := data.JsonReader("/Users/abhinav-m/Work/recipe-recommender-backend/all_recipes.json")
	recipes := getRecipeData()

	id := c.Param("id")

	if recipe,ok := recipes[id]; ok {
		c.IndentedJSON(http.StatusOK,recipe)
		return 
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
	router.GET("/recommendations",getAllRecommendations)
	router.GET("/recommendation/:id",getRecommendation)

	port := os.Getenv("PORT")

	if port == ""  {
		port = "8080"
	}

	// router.GET("/predict-recipe/:id",predictRecipe)
	router.Run("localhost:"+port)
}