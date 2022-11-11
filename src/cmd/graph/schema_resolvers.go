package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/big"
	"regexp"
	"strings"

	"github.com/LuigiAzevedo/graphql_min_roman_numeral/cmd/graph/generated"
	"github.com/LuigiAzevedo/graphql_min_roman_numeral/cmd/graph/model"
)

// find the smallest prime roman numeral inside a word
func (r *mutationResolver) Search(ctx context.Context, input model.Word) (*model.Payload, error) {
	// map with roman numerals and respective integers
	var RomanNumerals = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// filter all chars that aren't roman numerals and split them with a blank space
	re := regexp.MustCompile("[^XIVLCDM]")
	text := re.ReplaceAllString(input.Text, " ")

	// save all roman numerals in a slice
	numerals := strings.Fields(text)

	// variables used to store the smallest roman numeral and integer
	var minRoman string
	var minValue int

	for _, numeral := range numerals {

		total, greatest := 0, 0
		for i := len(numeral) - 1; i >= 0; i-- {
			// convert a single roman numeral to integer
			num := RomanNumerals[rune(numeral[i])]

			// verify if the roman numeral will add or deduct from the total
			if num >= greatest {
				greatest = num
				total += num
				continue
			}
			total -= num
		}

		// check if the integer is a prime number and store inside the map
		// ProbablyPrime is 100% accurate for inputs less than 2^64
		if big.NewInt(int64(total)).ProbablyPrime(0) {
			// only store the smallest number
			if minValue == 0 || total < minValue {
				minValue = total
				minRoman = numeral
			}
		}
	}

	return &model.Payload{
		Number: &minRoman,
		Value:  &minValue,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
