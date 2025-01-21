package main

import (
	"fmt"
	"log"

	"github.com/jdkato/prose/v2"
	"github.com/knuppe/vader"
	"github.com/syedazeez337/GoNLP/analysis"
	"github.com/syedazeez337/GoNLP/input"
)

func main() {

	// Example text to extract noun phrases from
	text := `Hello world! This is a test text. Is this working?
		         It has multiple sentences... Some with ellipsis.
		         And some with exclamation marks! Let's see how it works.`

	// Create a new document using prose
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	// parts of speech tagging
	for _, token := range doc.Tokens() {
		fmt.Printf("Word- %8s, Tag- %s\n", token.Text, token.Tag)
	}

	// Extract noun phrases based on part-of-speech tags
	var nounPhrases []string
	var currentPhrase []string

	// Loop through the tokens (words) in the document
	for _, tok := range doc.Tokens() {
		// Check for nouns (singular/plural) or proper nouns
		if tok.Tag == "NN" || tok.Tag == "NNS" || tok.Tag == "NNP" || tok.Tag == "NNPS" || tok.Tag == "DT" || tok.Tag == "JJ" {
			currentPhrase = append(currentPhrase, tok.Text)
		} else {
			// If we encounter a non-noun, store the current noun phrase and reset
			if len(currentPhrase) > 0 {
				nounPhrases = append(nounPhrases, fmt.Sprintf("%s", currentPhrase))
				currentPhrase = nil
			}
		}
	}

	// If the last phrase is a noun phrase, add it to the list
	if len(currentPhrase) > 0 {
		nounPhrases = append(nounPhrases, fmt.Sprintf("%s", currentPhrase))
	}

	// Print the extracted noun phrases
	fmt.Println("Extracted Noun Phrases:")
	for _, np := range nounPhrases {
		fmt.Println(np)
	}

	// sentiment analysis
	v, err := vader.NewVader("en.zip")
	if err != nil {
		panic(err)
	}

	sentiment := v.PolarityScores(text)

	fmt.Println(sentiment)

	// positive example
	positiveExample := []string{
		"I love programming in Go!",
		"Go is a great language for backend developement",
		"TextBlob is an amazing tool for NLP",
	}

	// negative example
	negativeExample := []string{
		"I hate bugs in my code.",
		"Debugging is so frustating.",
		"Errors are annoying while programming.",
	}

	nb := analysis.NewNaiveBayesCalssifier()
	nb.Train(positiveExample, negativeExample)

	fmt.Print("Get an input from the user: ")
	testText := input.GetInput()
	result := nb.Predict(testText)

	fmt.Println("Prediction for test:", testText)
	fmt.Println("Sentiment:", result)

	/*
		// detect a language
		testText := "Bonjour tout le monde!"

		translate.DetectLang(testText)

		inflections.ToPlural("Cat")
		inflections.ToSingular("Rabbits")
	*/

}
