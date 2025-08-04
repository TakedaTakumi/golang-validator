package main

import (
	"fmt"
	"golang-validator/pkg/validation"
)

func main() {
	fmt.Println("Go Validation Package Demo")
	fmt.Println("==========================")
	
	// String Length Validation Demo
	fmt.Println("\n1. String Length Validation:")
	stringValidator := validation.NewStringLengthValidator(3, 10)
	
	testStrings := []string{"ab", "hello", "this is too long"}
	for _, str := range testStrings {
		result := stringValidator.Validate(str)
		fmt.Printf("  Input: '%s' -> Valid: %t, Message: %s\n", str, result.IsValid, result.Message)
	}
	
	// Number Range Validation Demo
	fmt.Println("\n2. Number Range Validation:")
	numberValidator := validation.NewNumberRangeValidator(10, 100)
	
	testNumbers := []interface{}{5, 50, 150, 99.5}
	for _, num := range testNumbers {
		result := numberValidator.Validate(num)
		fmt.Printf("  Input: %v -> Valid: %t, Message: %s\n", num, result.IsValid, result.Message)
	}
	
	// Email Validation Demo
	fmt.Println("\n3. Email Validation:")
	emailValidator := validation.NewEmailValidator()
	
	testEmails := []string{
		"user@example.com",
		"invalid-email",
		"test@domain.co.jp",
		"@invalid.com",
	}
	
	for _, email := range testEmails {
		result := emailValidator.Validate(email)
		fmt.Printf("  Input: '%s' -> Valid: %t, Message: %s\n", email, result.IsValid, result.Message)
	}
	
	fmt.Println("\nDemo completed!")
}