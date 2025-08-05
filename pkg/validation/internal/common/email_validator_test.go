package common

import (
	"strings"
	"testing"
	
	"pgregory.net/rapid"
)

func TestEmailValidator(t *testing.T) {
	t.Run("正しい形式のメールアドレスは有効と判定される", func(t *testing.T) {
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
	
	t.Run("不正な形式のメールアドレスは無効と判定される", func(t *testing.T) {
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
	
	t.Run("文字列以外の入力は無効と判定される", func(t *testing.T) {
		validator := NewEmailValidator()
		
		result := validator.Validate(123)
		if result.IsValid {
			t.Errorf("expected validation to fail for non-string input")
		}
	})
}

func TestEmailValidatorProperties(t *testing.T) {
	t.Run("プロパティ_生成された有効なメールアドレスは必ず有効と判定される", func(t *testing.T) {
		validator := NewEmailValidator()
		
		rapid.Check(t, func(t *rapid.T) {
			email := genValidEmail().Draw(t, "email")
			result := validator.Validate(email)
			if !result.IsValid {
				t.Fatalf("expected validation to pass for valid email: %s", email)
			}
		})
	})
	
	t.Run("プロパティ_アットマークを含まない文字列は必ず無効と判定される", func(t *testing.T) {
		validator := NewEmailValidator()
		
		rapid.Check(t, func(t *rapid.T) {
			// アットマークを含まない文字列を生成
			email := rapid.String().
				Filter(func(s string) bool { return !strings.Contains(s, "@") }).
				Draw(t, "emailWithoutAt")
			
			result := validator.Validate(email)
			if result.IsValid {
				t.Fatalf("expected validation to fail for string without @: %s", email)
			}
		})
	})
	
	t.Run("プロパティ_同じ入力に対する検証結果は常に同一である", func(t *testing.T) {
		validator := NewEmailValidator()
		
		rapid.Check(t, func(t *rapid.T) {
			input := rapid.String().Draw(t, "input")
			
			result1 := validator.Validate(input)
			result2 := validator.Validate(input)
			
			if result1.IsValid != result2.IsValid {
				t.Fatalf("validation results differ for same input: %s", input)
			}
		})
	})
	
	t.Run("プロパティ_空文字列は必ず無効と判定される", func(t *testing.T) {
		validator := NewEmailValidator()
		
		result := validator.Validate("")
		if result.IsValid {
			t.Errorf("expected validation to fail for empty string")
		}
	})
}

// genValidEmail generates valid email addresses for property-based testing
func genValidEmail() *rapid.Generator[string] {
	return rapid.Custom(func(t *rapid.T) string {
		// Valid local part (before @)
		localPart := rapid.OneOf(
			rapid.StringMatching(`^[a-zA-Z0-9]+$`),
			rapid.StringMatching(`^[a-zA-Z0-9]+[._][a-zA-Z0-9]+$`),
			rapid.StringMatching(`^[a-zA-Z0-9]+\+[a-zA-Z0-9]+$`),
		).Draw(t, "localPart")
		
		// Valid domain
		domain := rapid.OneOf(
			rapid.Just("example.com"),
			rapid.Just("test.org"),
			rapid.Just("sample.net"),
			rapid.Just("demo.co.jp"),
			rapid.StringMatching(`^[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`),
			rapid.StringMatching(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`),
		).Draw(t, "domain")
		
		return localPart + "@" + domain
	})
}