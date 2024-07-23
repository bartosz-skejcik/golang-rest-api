package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"regexp"
	"restapi/internal/models"
)

func ReadJSONFile[T any](path string, result *T) error {
	// Open JSON file
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, result)
}

func WriteJSONFile[T any](filename string, data T) error {
	// Open or create the file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON encoder and encode the data into the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: format the JSON with indentation
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func ValidateUser(user models.User, users models.Users) error {
	if UserExists(users, user.Email) {
		return errors.New("User already exists")
	}

	// Check if Username is not empty and meets length requirements
	if len(user.Username) == 0 {
		return errors.New("username cannot be empty")
	}
	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	// Check if Email is not empty and matches a simple email pattern
	if len(user.Email) == 0 {
		return errors.New("email cannot be empty")
	}
	if !IsValidEmail(user.Email) {
		return errors.New("email format is invalid")
	}

	// Check if FullName is not empty
	if len(user.FullName) == 0 {
		return errors.New("full name cannot be empty")
	}

	// If all checks pass, return nil (no error)
	return nil
}

func IsValidEmail(email string) bool {
	// Simple regex pattern for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func UserExists(users []models.User, userEmail string) bool {
	for _, user := range users {
		if user.Email == userEmail {
			return true
		}
	}
	return false
}
