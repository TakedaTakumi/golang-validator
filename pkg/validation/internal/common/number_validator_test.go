package common

import (
	"testing"
	
	"pgregory.net/rapid"
)

func TestNumberRangeValidator(t *testing.T) {
	t.Run("整数の最小値境界で正しく検証される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate(5)
		if result.IsValid {
			t.Errorf("expected validation to fail for integer below minimum")
		}
		
		result = validator.Validate(10)
		if !result.IsValid {
			t.Errorf("expected validation to pass for integer at minimum")
		}
	})
	
	t.Run("整数の最大値境界で正しく検証される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate(105)
		if result.IsValid {
			t.Errorf("expected validation to fail for integer above maximum")
		}
		
		result = validator.Validate(100)
		if !result.IsValid {
			t.Errorf("expected validation to pass for integer at maximum")
		}
	})
	
	t.Run("小数の最小値境界で正しく検証される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		result := validator.Validate(9.5)
		if result.IsValid {
			t.Errorf("expected validation to fail for float below minimum")
		}
		
		result = validator.Validate(10.5)
		if !result.IsValid {
			t.Errorf("expected validation to pass for float at minimum")
		}
	})
	
	t.Run("小数の最大値境界で正しく検証される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		result := validator.Validate(101.0)
		if result.IsValid {
			t.Errorf("expected validation to fail for float above maximum")
		}
		
		result = validator.Validate(100.5)
		if !result.IsValid {
			t.Errorf("expected validation to pass for float at maximum")
		}
	})
	
	t.Run("数値以外の入力は無効と判定される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate("not a number")
		if result.IsValid {
			t.Errorf("expected validation to fail for non-number input")
		}
	})
}

func TestNumberRangeValidatorProperties(t *testing.T) {
	t.Run("プロパティ_範囲内の整数は必ず有効と判定される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		rapid.Check(t, func(t *rapid.T) {
			num := rapid.IntRange(10, 100).Draw(t, "num")
			result := validator.Validate(num)
			if !result.IsValid {
				t.Fatalf("expected validation to pass for integer in range: %d", num)
			}
		})
	})
	
	t.Run("プロパティ_範囲内の小数は必ず有効と判定される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		rapid.Check(t, func(t *rapid.T) {
			num := rapid.Float64Range(10.5, 100.5).Draw(t, "num")
			result := validator.Validate(num)
			if !result.IsValid {
				t.Fatalf("expected validation to pass for float in range: %f", num)
			}
		})
	})
	
	t.Run("プロパティ_最小値未満の数値は必ず無効と判定される", func(t *testing.T) {
		validator := NewNumberRangeValidator(50, 100)
		
		rapid.Check(t, func(t *rapid.T) {
			num := rapid.IntMax(49).Draw(t, "num")
			result := validator.Validate(num)
			if result.IsValid {
				t.Fatalf("expected validation to fail for integer below minimum: %d", num)
			}
		})
	})
	
	t.Run("プロパティ_最大値超過の数値は必ず無効と判定される", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 50)
		
		rapid.Check(t, func(t *rapid.T) {
			num := rapid.IntMin(51).Draw(t, "num")
			result := validator.Validate(num)
			if result.IsValid {
				t.Fatalf("expected validation to fail for integer above maximum: %d", num)
			}
		})
	})
	
	t.Run("プロパティ_同じ入力に対する検証結果は常に同一である", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		rapid.Check(t, func(t *rapid.T) {
			num := rapid.Int().Draw(t, "num")
			result1 := validator.Validate(num)
			result2 := validator.Validate(num)
			
			if result1.IsValid != result2.IsValid {
				t.Fatalf("validation results differ for same input: %d", num)
			}
		})
	})
}