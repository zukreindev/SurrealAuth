package util

import "strings"

func ValidateUsername(username string) bool {
	// username must be at least 4 characters long
	if len(username) < 4 {
		Log("Validations", "Username too short")
		return false
	}

	// username must be at most 20 characters long
	if len(username) > 20 {
		Log("Validations", "Username too long")
		return false
	}

	// username must not contain any special characters
	for _, char := range username {
		if char < 'a' || char > 'z' {
			Log("Validations", "Username contains special characters")
			return false
		}
	}

	return true
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		Log("Validations", "Password too short")
		return false
	}

	if len(password) > 64 {
		Log("Validations", "Password too long")
		return false
	}

	return true
}

func ValidateEmail(email string) bool {
	// email must be at least 4 characters long
	if len(email) < 4 {
		Log("Validations", "Email too short")
		return false
	}
	// email must be at most 20 characters long
	if len(email) > 20 {
		Log("Validations", "Email too long")
		return false
	}

	// email must contain an @ symbol
	if !strings.Contains(email, "@") {
		Log("Validations", "Email does not contain @ symbol")
		return false
	}

	// email regex 
	
	
	return true
}
