package analysis

import (
	"fmt"
	"strings"
)

// implementation of naive bayes analysis
type NaiveBayesClassifier struct {
	positiveWords  map[string]int
	negativeWords  map[string]int
	totalPositive  int
	totalNegative  int
	vocabularySize int
}

// new instance of naive bayes classifier
func NewNaiveBayesCalssifier() *NaiveBayesClassifier {
	return &NaiveBayesClassifier{
		positiveWords: make(map[string]int),
		negativeWords: make(map[string]int),
	}
}

// train method on positive and negative words
func (nb *NaiveBayesClassifier) Train(posExamples, negExamples []string) {
	// positive examples
	for _, text := range posExamples {
		words := tokenize(text)
		for _, word := range words {
			nb.positiveWords[word]++
			nb.totalPositive++
			// nb.totalWordCount++
		}
	}

	// train on negative examples
	for _, text := range negExamples {
		words := tokenize(text)
		for _, word := range words {
			nb.negativeWords[word]++
			nb.totalNegative++
			// nb.totalWordCount++
		}
	}

	// calculate vocabulary size
	nb.vocabularySize = len(mergeKeys(nb.positiveWords, nb.negativeWords))
}

// merge key function
func mergeKeys(m1, m2 map[string]int) map[string]bool {
	keys := make(map[string]bool)

	for key := range m1 {
		keys[key] = true
	}

	for key := range m2 {
		keys[key] = true
	}

	return keys
}

// tokenize function
func tokenize(text string) []string {
	// convert lowercase and split the text into words
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

// calculate the probability of being positive or negative
func (nb *NaiveBayesClassifier) Predict(text string) string {
	words := tokenize(text)

	totalWordCount := nb.totalPositive + nb.totalNegative
	positiveProb := float64(nb.totalPositive) / float64(totalWordCount)
	negativeProb := float64(nb.totalNegative) / float64(totalWordCount)

	fmt.Println(positiveProb, negativeProb)

	posScore := positiveProb
	negScore := negativeProb

	alpha := 1.0

	for _, word := range words {
		posWordProb := 
		(float64(nb.positiveWords[word]) + alpha) / (float64(nb.totalPositive) + alpha * float64(nb.vocabularySize))
		negWordProb := 
		(float64(nb.negativeWords[word]) + alpha) / (float64(nb.totalNegative) + alpha * float64(nb.vocabularySize))

		posScore *= posWordProb
		negScore *= negWordProb
	}

	fmt.Println(posScore, negScore)

	if posScore > negScore {
		return "Positive"
	}

	return "Negative"
}
