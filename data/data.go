package data

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type Recipe struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Ingredients []string `json:"ingredients"`
	Calories string `json:"calories"`
}

func formatAndSplit(s string) []string{
	s = strings.TrimSuffix(s,"]")
	s = strings.TrimPrefix(s,"[")
	s = strings.ReplaceAll(s,"'","")
	return strings.Split(s,",")
}

func createRecipeList(data [][]string) []Recipe {
	var recipeList []Recipe
	for i, line := range data {
		// Omit header details for csv files
		if i  > 0 {
			var record Recipe
			for j, field := range line {
				switch j {
				case 4:
					record.ID = field
				case 6:
					record.Title = field
				case 1:
					record.Category = field
				case 2:
					record.Ingredients = formatAndSplit(field)
				case 0:
					record.Calories = field
				}
			}
			if record.Title != "Recipe webpage is unconventional" {
				recipeList = append(recipeList,record)
			}
		
		}
	}

	return recipeList
}

func CsvReaderClosure(csvPath string) func() []Recipe {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	recipeData := createRecipeList(data)


	return func() []Recipe {
		return recipeData
	}
}