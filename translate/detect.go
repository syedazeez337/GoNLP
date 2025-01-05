package translate

import (
	"fmt"

	"github.com/abadojack/whatlanggo"
)

func DetectLang(text string) {
	info := whatlanggo.Detect(text)

	fmt.Println("Detected language: ", info.Lang)
	fmt.Println("Confidence: ", info.Confidence)
}