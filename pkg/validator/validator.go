package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func HandleRequestError(err error) (string, string, string) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		e := errs[0]
		return strings.ToLower(e.Field()), e.Tag(), e.Param()
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return unmarshalTypeError.Field, "type_error", unmarshalTypeError.Type.String()
	}

	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		return "json", "syntax_error", fmt.Sprintf("%d", syntaxError.Offset)
	}

	return "unknown", "invalid", ""
}