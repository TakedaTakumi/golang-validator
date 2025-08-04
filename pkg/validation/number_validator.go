package validation

// NumberRangeValidator validates numbers within specified range
type NumberRangeValidator struct {
	Min float64
	Max float64
}

// NewNumberRangeValidator creates a new NumberRangeValidator
func NewNumberRangeValidator(min, max float64) *NumberRangeValidator {
	return &NumberRangeValidator{
		Min: min,
		Max: max,
	}
}

// Validate checks if the input number is within the specified range
func (v NumberRangeValidator) Validate(value interface{}) ValidationResult {
	var num float64
	var ok bool
	
	switch val := value.(type) {
	case int:
		num = float64(val)
		ok = true
	case float64:
		num = val
		ok = true
	case float32:
		num = float64(val)
		ok = true
	default:
		ok = false
	}
	
	if !ok {
		return ValidationResult{
			IsValid: false,
			Message: "input must be a number",
		}
	}
	
	if num < v.Min {
		return ValidationResult{
			IsValid: false,
			Message: "number is too small",
		}
	}
	
	if num > v.Max {
		return ValidationResult{
			IsValid: false,
			Message: "number is too large",
		}
	}
	
	return ValidationResult{
		IsValid: true,
		Message: "valid",
	}
}