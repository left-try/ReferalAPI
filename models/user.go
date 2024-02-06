package models

import (
	"awesomeProject/database"
	"awesomeProject/utils"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	encryptedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword
	exec, err := prepare.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	user.Id = id
	return err
}

func (user *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := database.DB.QueryRow(query, user.Email)
	var receivedPassword string
	err := row.Scan(&user.Id, &receivedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckHashedPassword(user.Password, receivedPassword)
	if !passwordIsValid {
		return errors.New("user is not valid")
	}
	return nil
}
