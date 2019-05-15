// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func init() {
	if validator, isValidator := binding.Validator.Engine().(*validator.Validate); isValidator {
		validator.RegisterValidation("locale", validateLocale)
	}
}

// Custom validation rule for Locale.
// (can be replaced by native oneof from v9)
func validateLocale(
	validationSettings *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string,
) bool {
	if localeFieldValue, isLocaleField := field.Interface().(Locale); isLocaleField {
		if _, isLocaleSupported := localeSupported[localeFieldValue]; isLocaleSupported {
			return true
		}
	}

	return false
}
