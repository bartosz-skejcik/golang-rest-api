package functions

import (
	"errors"
	"log"
	"restapi/internal/models"
	"restapi/internal/utils"
)

func GetAllUsers(users *models.Users) error {
	var data models.DatabaseType
	err := utils.ReadJSONFile("data/database.json", &data)
	if err != nil {
		log.Printf("Failed to read JSON file: %v", err.Error())
		return err
	}

	*users = data.Users

	return nil
}

func GetUserById(user *models.User, userId int) error {
	var data models.DatabaseType
	readErr := utils.ReadJSONFile("data/database.json", &data)
	if readErr != nil {
		return errors.New("Failed to read from the database")
	}

	var userObject models.User

	for _, user := range data.Users {
		if user.Id == userId {
			userObject = user
		}
	}

	if userObject.Id == 0 {
		return errors.New("No user with the given id")
	} else {
		*user = userObject
	}

	return nil
}

func CreateUser(user *models.User) error {
	var users models.Users
	var data models.DatabaseType

	if err := GetAllUsers(&users); err != nil {
		return err
	}

	if err := utils.ValidateUser(*user, users); err != nil {
		return err
	}

	// Determine the new ID
	var newID int
	if len(users) > 0 {
		newID = users[len(users)-1].Id + 1
	} else {
		newID = 1 // Start with 1 if no users exist
	}

	// Assign the new ID to the user
	user.Id = newID

	users = append(users, *user)

	data.Users = users

	// Write the updated data to the JSON file
	if err := utils.WriteJSONFile("data/database.json", &data); err != nil {
		return err
	}

	return nil
}
