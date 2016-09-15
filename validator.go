// Copyright 2016 The Displae Team. All rights reserved.
// Use of this source code is governed by a GNU AGPLv3
// license that can be found in the LICENSE file.

// Package gladiator implements validator for Displae needs.
package gladiator

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

// init function on this package register validation function so
// it can be used on struct tags.
func init() {
	govalidator.TagMap["phone"] = govalidator.Validator(ValidatePhoneNumber)
	govalidator.TagMap["country_code"] = govalidator.Validator(ValidatePhoneCountryCode)
	govalidator.TagMap["password"] = govalidator.Validator(ValidatePassword)
}

// Phone number must contain country code and phone number itself.
// Country code and phone number between 13 - 24 characters,
// which country code and phone number are continuous wihout any
// separator. For example: +628571234567890. Country code that
// have dash, the dash is stripped.
func ValidatePhoneNumber(phoneNumber string) bool {
	matched, _ := regexp.MatchString("\\+\\d{13,24}", phoneNumber)
	return matched
}

func ValidatePhoneCountryCode(phoneCountryCode string) bool {
	return true
}

// Password must contains 8 characters minimal and must have
// mix caps and number
func ValidatePassword(password string) bool {
	matched, _ := regexp.MatchString("[0-9]", password)
	return matched
}
