package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type recipe struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Ingredients []string `json:"ingredients"`
	Image string `json:"image"`
	Description string `json:"Description"`
	CuisineType string `json:"CuisineType"`
}

var recipes  = []recipe {
	{ID:"1", Title:"Chicken ghee me", Ingredients: []string{"Chicken","ghee","onions","tomato"},Image:"https://www.simplyrecipes.com/thmb/xNRMdPJcmR20G5gcwBjndiMxYBk=/736x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2010__06__tandoori-chicken-horiz-a-1600-a92053df1c764ee1beaa91ae6383dcfd.jpg" ,Description :"Chicken and ghee",CuisineType:"Indian"},
	{ID:"2", Title:"Chicken ghee me", Ingredients: []string{"Chicken","ghee","onions","tomato"},Image:"https://www.simplyrecipes.com/thmb/xNRMdPJcmR20G5gcwBjndiMxYBk=/736x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2010__06__tandoori-chicken-horiz-a-1600-a92053df1c764ee1beaa91ae6383dcfd.jpg" ,Description :"Chicken and ghee",CuisineType:"Indian"},
	{ID:"3", Title:"Chicken ghee me", Ingredients: []string{"Chicken","ghee","onions","tomato"},Image:"https://www.simplyrecipes.com/thmb/xNRMdPJcmR20G5gcwBjndiMxYBk=/736x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2010__06__tandoori-chicken-horiz-a-1600-a92053df1c764ee1beaa91ae6383dcfd.jpg" ,Description :"Chicken and ghee",CuisineType:"Indian"},
}

/*
gin.Context is the most important part of Gin.
 It carries request details, validates and serializes JSON, and more. 
*/
func getAllRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,recipes)
}

func getRecipe(c *gin.Context) {

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