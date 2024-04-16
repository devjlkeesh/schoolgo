package services

import (
	"errors"
	"log"
	"school23/internal/database"
	"school23/internal/models"
	"school23/internal/utils"
)

func GetUsers() []models.User {
	return database.GetAllUsers()
}

func UserCreate(username string, password string, role string) error {
	exists, err := database.UserExistsByUsername(username)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	if exists {
		return errors.New(utils.UsernameAlreadTaken)
	}
	user := models.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	err = database.InsertUser(user)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func UserUpdate(id int, username string, role string, status bool) error {
	user, err := database.FindUserByUsername(username)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	if user.Id != id {
		return errors.New(utils.UsernameAlreadTaken)
	}
	err = database.UpdateUser(id, username, role, status)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func GetUser(id int) (models.User, error) {
	user, err := database.FindUserById(id)
	if err != nil {
		log.Println(err)
		return models.User{}, errors.New(utils.UserNotFound)
	}
	return user, nil
}

func DeleteUserById(id int) error {
	err := database.DeleteUserById(id)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}
