package api

import (
	"github.com/etharra/simplebank/util"
	"github.com/go-playground/validator/v10"
)

/**
 * validCurrency is a custom validator function for checking if a given currency is supported.
 * It implements the validator.Func interface.
 *
 * Parameters:
 * - fl: validator.FieldLevel interface containing information about the field being validated
 *
 * Returns:
 * - bool: true if the currency is supported, false otherwise
 */
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
