package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// type Recipe struct {
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Category string `json:"category"`
// 	Ingredients []string `json:"ingredients"`
// 	Calories string `json:"calories"`
// 	CookTime string `json:"time_in_mins"`
// 	Rating float32 `json:"rating"`
// 	SourScore int8 `json:"sour_score"`
// 	SaltScore int8 `json:"salt_score"`
// 	SweetScore int8 `json:"sweet_score"`
// 	BitterScore int8 `json:"bitter_score"`
// }


type mapType map[string]interface{}

func createRecipeList(data []byte) mapType {
	var parsedJson mapType
	

	if err := json.Unmarshal(data,&parsedJson);err !=nil{
		log.Fatal(err)
	}

	return parsedJson

}



// Export functions in capital letters
func JsonReader(jsonPath string) func() mapType {
	f, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}
	
	recipeData := createRecipeList(byteValue)


	return func() mapType {
		return recipeData
	}
}

type USER_RATING struct {
	USER string `json:"user" binding:"required"`
	RECIPE_ID string  `json:"recipe_id" binding:"required"`
	RATING string  `json:"rating_score" binding:"required"`
}


// Export functions in capital letters
func JsonArrayReader(jsonPath string) func() []USER_RATING{
	f, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}
	
	var existingRatings []USER_RATING


	if err := json.Unmarshal(byteValue,&existingRatings);err !=nil{
		log.Fatal(err)
	}


	return func() []USER_RATING {
		return existingRatings
	}
}




func RecommendationRatingSaver(ratingFilePath string, userRating USER_RATING ) bool {

	f,err := os.OpenFile(ratingFilePath,os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)


	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var existingRatings []USER_RATING

	byteValue, _ := ioutil.ReadAll(f)

	_ = json.Unmarshal(byteValue,&existingRatings)
	
	existingRatings  = append(existingRatings,userRating)

	file, _ := json.MarshalIndent(existingRatings, "", " ")

	_ = ioutil.WriteFile(ratingFilePath,file,0644)

	return true



}