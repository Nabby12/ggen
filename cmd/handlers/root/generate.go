package root

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

var keys = map[string]string{
	"l": "lower",
	"n": "numeric",
	"s": "symbol",
	"u": "upper",
}

var letters = map[string]string{
	keys["l"]: "abcdefghijklmnopqrstuvwxyz",
	keys["u"]: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	keys["n"]: "0123456789",
	keys["s"]: "!@#$%^&*",
}

func Generate(count int, noLowercase bool, noNumeric bool, noSymbol bool, noUppercase bool) (string, error) {
	if count == 0 {
		return "", errors.New("word count ('-c' option) must be greater than '0'")
	}

	if noLowercase && noNumeric && noSymbol && noUppercase {
		return "", errors.New("the all types of 'no-xx' flags cannot be used at once.")
	}

	typeCount := 0
	wordCount := map[string]int{}
	if !noLowercase {
		typeCount++
		wordCount[keys["l"]] = count
	}
	if !noNumeric {
		typeCount++
		wordCount[keys["n"]] = count
	}
	if !noSymbol {
		typeCount++
		wordCount[keys["s"]] = count
	}
	if !noUppercase {
		typeCount++
		wordCount[keys["u"]] = count
	}

	if typeCount > count {
		return "", errors.New("not enough word count ('-c' option) for selected options")
	}

	rand.Seed(time.Now().UnixNano())

	// get full length for selected options
	totalCount := 0
	keyMap := []string{}
	for key, value := range wordCount {
		totalCount += value
		keyMap = append(keyMap, key)
	}

	// reduce the word count to the specified
	for i := totalCount; i > count; i-- {
		min := 0
		max := len(keyMap)
		target := rand.Intn(max-min) + min

		if wordCount[keyMap[target]] == 1 {
			i++
			continue
		}

		wordCount[keyMap[target]]--
	}

	var buf bytes.Buffer
	for key, value := range wordCount {
		for i := 0; i < value; i++ {
			min := 0
			max := len(letters[key])
			target := rand.Intn(max-min) + min

			buf.WriteString(string(letters[key][target]))
		}
	}

	var pass bytes.Buffer
	randomOrder := rand.Perm(len(buf.String()))
	for _, value := range randomOrder {
		pass.WriteString(string(buf.String()[value]))
	}

	return pass.String(), nil
}
