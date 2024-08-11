package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key string, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = map[string]string{}
	}
	if _, exits := v.FieldErrors[key]; !exits {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key string, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedInt(value int, permittedValues ...int) bool {
	for _, permitted := range permittedValues {
		if value == permitted {
			return true
		}
	}
	return false
}
