package inflections

import (
	"fmt"

	"github.com/jinzhu/inflection"
)

func ToPlural(word string) {
	pluralWd := inflection.Plural(word)

	fmt.Printf("Singular: %s, Plural: %s\n", word, pluralWd)
}

func ToSingular(word string) {
	singularWd := inflection.Singular(word)

	fmt.Printf("Plural: %s, Singular: %s\n", word, singularWd)
}