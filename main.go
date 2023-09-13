package main

import (
	"fmt"
	"regexp"
)

func main() {
	creditCardList := []string{
		"4123456789123456",
		"5123-4567-8912-3456",
		"61234-567-8912-3456",
		"4123356789123456",
		"5133-3367-8912-3456",
		"5123 - 3567 - 8912 - 3456",
	}

	resultCheck := validateCreditCardNumbers(creditCardList)
	for _, valid := range resultCheck {
		if valid {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}

func validateCreditCardNumbers(creditCardList []string) []bool {
	validResults := make([]bool, len(creditCardList))

	for i, creditCardString := range creditCardList {
		validResults[i] = isValidCreditCardNumber(creditCardString)
	}

	return validResults
}

func isValidCreditCardNumber(creditCardString string) bool {
	// Check for valid separators using regular expression
	validSeparator := regexp.MustCompile(`^([0-9]{4}-?){3}[0-9]{4}$`)
	if !validSeparator.MatchString(creditCardString) {
		return false
	}

	// Remove hyphens for further validation
	creditCardString = regexp.MustCompile(`-`).ReplaceAllString(creditCardString, "")

	if len(creditCardString) != 16 {
		return false
	}

	if hasNoRepeatOfFourConsecutiveDigits(creditCardString) &&
		containsOnlyNumbers(creditCardString) &&
		beginsWithCorrectNumber(creditCardString) {
		return true
	}

	return false
}

func hasNoRepeatOfFourConsecutiveDigits(creditCardString string) bool {
	for i := 0; i < len(creditCardString)-3; i++ {
		if creditCardString[i] == creditCardString[i+1] &&
			creditCardString[i] == creditCardString[i+2] &&
			creditCardString[i] == creditCardString[i+3] {
			return false
		}
	}
	return true
}

func containsOnlyNumbers(creditCardString string) bool {
	validNumber := regexp.MustCompile(`^[0-9]+$`)
	return validNumber.MatchString(creditCardString)
}

func beginsWithCorrectNumber(creditCardString string) bool {
	firstChar := creditCardString[0]
	return firstChar == '4' || firstChar == '5' || firstChar == '6'
}
