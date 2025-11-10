package validator

import (
	"fmt"

	"github.com/juster/antlr-playgrounds/internal/parser"
)

type ValidationResult struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func ValidateFormula(formulaStr string) ValidationResult {
	_, err := parser.Parser.ParseString("", formulaStr)
	if err != nil {
		return ValidationResult{
			Valid: false,
			Error: fmt.Sprintf("syntax error: %v", err),
		}
	}

	return ValidationResult{
		Valid: true,
	}
}
