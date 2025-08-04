package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"testing/quick"
)

func TestEmailValidator(t *testing.T) {
	t.Run("shouldValidateValidEmail", func(t *testing.T) {
		validator := NewEmailValidator()
		
		validEmails := []string{
			"test@example.com",
			"user.name@domain.co.jp",
			"firstname+lastname@example.org",
			"email123@test-domain.com",
		}
		
		for _, email := range validEmails {
			result := validator.Validate(email)
			if !result.IsValid {
				t.Errorf("expected validation to pass for valid email: %s", email)
			}
		}
	})
	
	t.Run("shouldRejectInvalidEmail", func(t *testing.T) {
		validator := NewEmailValidator()
		
		invalidEmails := []string{
			"invalid-email",
			"@example.com",
			"test@",
			"test..test@example.com",
			"test@example",
			"",
		}
		
		for _, email := range invalidEmails {
			result := validator.Validate(email)
			if result.IsValid {
				t.Errorf("expected validation to fail for invalid email: %s", email)
			}
		}
	})
	
	t.Run("shouldHandleNonStringInput", func(t *testing.T) {
		validator := NewEmailValidator()
		
		result := validator.Validate(123)
		if result.IsValid {
			t.Errorf("expected validation to fail for non-string input")
		}
	})
}

func TestEmailValidatorProperties(t *testing.T) {
	t.Run("propertyValidEmailsShouldPass", func(t *testing.T) {
		validator := NewEmailValidator()
		
		property := func() bool {
			email := generateValidEmail()
			result := validator.Validate(email)
			return result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyEmailsWithoutAtShouldFail", func(t *testing.T) {
		validator := NewEmailValidator()
		
		property := func() bool {
			email := generateStringWithoutAt()
			if strings.Contains(email, "@") {
				return true
			}
			
			result := validator.Validate(email)
			return !result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyValidationResultsAreConsistent", func(t *testing.T) {
		validator := NewEmailValidator()
		
		property := func(input string) bool {
			result1 := validator.Validate(input)
			result2 := validator.Validate(input)
			
			return result1.IsValid == result2.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyEmptyStringShouldFail", func(t *testing.T) {
		validator := NewEmailValidator()
		
		result := validator.Validate("")
		if result.IsValid {
			t.Errorf("expected validation to fail for empty string")
		}
	})
}

func generateValidEmail() string {
	domains := []string{"example.com", "test.org", "sample.net", "demo.co.jp"}
	localNames := []string{"user", "test", "admin", "contact", "info"}
	
	local := localNames[rand.Intn(len(localNames))]
	domain := domains[rand.Intn(len(domains))]
	
	return fmt.Sprintf("%s@%s", local, domain)
}

func generateStringWithoutAt() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := rand.Intn(20) + 1
	
	var result strings.Builder
	result.Grow(length)
	
	for i := 0; i < length; i++ {
		result.WriteByte(chars[rand.Intn(len(chars))])
	}
	
	return result.String()
}