package common

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestStringLengthValidator(t *testing.T) {
	t.Run("最小文字数境界で正しく検証される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("ab")
		assert.False(t, result.IsValid, "expected validation to fail for string shorter than minimum length")
		
		result = validator.Validate("abc")
		assert.True(t, result.IsValid, "expected validation to pass for string at minimum length")
	})
	
	t.Run("最大文字数境界で正しく検証される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("12345678901")
		assert.False(t, result.IsValid, "expected validation to fail for string longer than maximum length")
		
		result = validator.Validate("1234567890")
		assert.True(t, result.IsValid, "expected validation to pass for string at maximum length")
	})
	
	t.Run("範囲内の文字数は有効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("hello")
		assert.True(t, result.IsValid, "expected validation to pass for string within valid range")
	})
	
	t.Run("文字列以外の入力は無効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate(123)
		assert.False(t, result.IsValid, "expected validation to fail for non-string input")
	})
}

func TestStringLengthValidatorProperties(t *testing.T) {
	t.Run("プロパティ_範囲内の文字数は必ず有効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		rapid.Check(t, func(t *rapid.T) {
			length := rapid.IntRange(5, 15).Draw(t, "length")
			// 文字のスライスを使用して正確な長さの文字列を生成
			chars := rapid.SliceOfN(rapid.RuneFrom([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")), length, length).Draw(t, "chars")
			testString := string(chars)
			result := validator.Validate(testString)
			assert.True(t, result.IsValid, "expected validation to pass for string of length %d: %s", len(testString), testString)
		})
	})
	
	t.Run("プロパティ_最小文字数未満は必ず無効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		rapid.Check(t, func(t *rapid.T) {
			length := rapid.IntRange(0, 4).Draw(t, "length")
			// 文字のスライスを使用して正確な長さの文字列を生成
			chars := rapid.SliceOfN(rapid.RuneFrom([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")), length, length).Draw(t, "chars")
			testString := string(chars)
			result := validator.Validate(testString)
			assert.False(t, result.IsValid, "expected validation to fail for string of length %d: %s", len(testString), testString)
		})
	})
	
	t.Run("プロパティ_最大文字数超過は必ず無効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		rapid.Check(t, func(t *rapid.T) {
			length := rapid.IntRange(16, 50).Draw(t, "length")
			// 文字のスライスを使用して正確な長さの文字列を生成
			chars := rapid.SliceOfN(rapid.RuneFrom([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")), length, length).Draw(t, "chars")
			testString := string(chars)
			result := validator.Validate(testString)
			assert.False(t, result.IsValid, "expected validation to fail for string of length %d: %s", len(testString), testString)
		})
	})
	
	t.Run("プロパティ_同じ入力に対する検証結果は常に同一である", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		rapid.Check(t, func(t *rapid.T) {
			input := rapid.String().Draw(t, "input")
			result1 := validator.Validate(input)
			result2 := validator.Validate(input)
			
			assert.Equal(t, result1.IsValid, result2.IsValid, "validation results differ for same input: %s", input)
		})
	})
}