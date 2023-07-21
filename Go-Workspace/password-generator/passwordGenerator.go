package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type Password string

type PasswordRequirements struct {
	MinLength        int
	RequireUpperCase bool
	RequireLowerCase bool
	RequireNumbers   bool
	RequireSpecial   bool
	SpecialChars     string
}

type ValidationResult struct {
	MinLength        bool
	RequireUpperCase bool
	RequireLowerCase bool
	RequireNumbers   bool
	RequireSpecial   bool
}

func (p Password) IsValid(requirements PasswordRequirements) (bool, ValidationResult) {
	passingNumber := 0
	var result ValidationResult

	// Check if password meets each requirement
	if len(p) >= requirements.MinLength {
		passingNumber++ // Increment passingNumber if minimum length requirement is met
		result.MinLength = true
	}

	// Check if password contains uppercase letters
	if requirements.RequireUpperCase && containsUpperCase(p) {
		passingNumber++ // Increment passingNumber if uppercase requirement is met
		result.RequireUpperCase = true
	}

	// Check if password contains lowercase letters
	if requirements.RequireLowerCase && containsLowerCase(p) {
		passingNumber++ // Increment passingNumber if lowercase requirement is met
		result.RequireLowerCase = true
	}

	// Check if password contains numbers
	if requirements.RequireNumbers && containsNumber(p) {
		passingNumber++ // Increment passingNumber if numbers requirement is met
		result.RequireNumbers = true
	}

	// Check if password contains special characters
	if requirements.RequireSpecial && containsSpecialChars(p, requirements.SpecialChars) {
		passingNumber++ // Increment passingNumber if special characters requirement is met
		result.RequireSpecial = true
	}

	return passingNumber == 5, result // Compare passingNumber with desired value
}

func containsUpperCase(password Password) bool {
	for _, char := range password {
		if strings.ContainsAny(string(char), "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return true
		}
	}
	return false
}

func containsLowerCase(password Password) bool {
	for _, char := range password {
		if strings.ContainsAny(string(char), "abcdefghijklmnopqrstuvwxyz") {
			return true
		}
	}
	return false
}

func containsNumber(password Password) bool {
	for _, char := range password {
		if strings.ContainsAny(string(char), "0123456789") {
			return true
		}
	}
	return false
}

func containsSpecialChars(password Password, specialChars string) bool {
	for _, char := range password {
		if strings.ContainsAny(string(char), specialChars) {
			return true
		}
	}
	return false
}

func generateRandomCharacter(pool string) (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
	if err != nil {
		return "", err
	}
	return string(pool[index.Int64()]), nil
}

func main() {
	requirements := PasswordRequirements{
		MinLength:        8,
		RequireUpperCase: true,
		RequireLowerCase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
		SpecialChars:     "!@#$%^&*",
	}

	characterPool := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890!@#$%^&*"

	fmt.Println("Enter desired password length (minimum 8): ")
	var desiredPasswordLength int
	for {
		_, err := fmt.Scanln(&desiredPasswordLength)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer value.")
			continue
		}

		if desiredPasswordLength < 8 {
			fmt.Println("Length must be 8 or greater. Please enter a valid length.")
		} else {
			break
		}
	}

	password := Password("")
	for i := 0; i < desiredPasswordLength; i++ {
		char, err := generateRandomCharacter(characterPool)
		if err != nil {
			fmt.Println("Error generating random character:", err)
			return
		}
		password += Password(char)
	}

	fmt.Println("Generated Password:", password)
	validated, res := password.IsValid(requirements)
	fmt.Println("Validated:", validated)
	fmt.Println("Validation Result:", res)
}
