package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type mapType map[string]interface{}



type Recipe struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Ingredients []string `json:"ingredients"`
	Calories string `json:"calories"`
	CookTime string `json:"time_in_mins"`
	Rating float32 `json:"rating"`
	SourScore int8 `json:"sour_score"`
	SaltScore int8 `json:"salt_score"`
	SweetScore int8 `json:"sweet_score"`
	BitterScore int8 `json:"bitter_score"`
}

func formatAndSplit(s string) []string{
	s = strings.TrimSuffix(s,"]")
	s = strings.TrimPrefix(s,"[")
	s = strings.ReplaceAll(s,"'","")
	return strings.Split(s,",")
}

func createRecipeList(data []byte) []Recipe {
	var parsedJson mapType
	

	if err := json.Unmarshal(data,&parsedJson);err !=nil{
		log.Fatal(err)
	}


}

func jsonReaderClosure(jsonPath string) func() []Recipe {
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


	return func() []Recipe {
		return recipeData
	}
}