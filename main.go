package main

import (
	"fmt"
	"log"
	"words_counter/services"
	"words_counter/services/random_text_api"
	"words_counter/services/static_text"
)

func main() {
	text1, err := services.NewCounter(static_text.GetText)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(text1)
	fmt.Println()

	text2, err := services.NewCounter(random_text_api.GetText)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(text2)
	fmt.Println()
}
