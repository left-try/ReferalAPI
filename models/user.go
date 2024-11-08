package models

import (
	"database/sql"
	"errors"
	"referralAPI/database"
	"referralAPI/utils"
)

type User struct {
	Id         int64
	Email      string `binding:"required"`
	Password   string `binding:"required"`
	ReferrerId int64
}

func (user *User) Save() error {
	var query string
	if user.ReferrerId == -1 {
		query = "INSERT INTO users(email, password) VALUES (?, ?)"
	} else {
		query = "INSERT INTO users(email, password, referrer_id) VALUES (?, ?, ?)"
	}
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			return
		}
	}(prepare)
	encryptedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword
	var exec sql.Result
	if user.ReferrerId == -1 {
		exec, err = prepare.Exec(user.Email, user.Password)
	} else {
		exec, err = prepare.Exec(user.Email, user.Password, user.ReferrerId)
	}
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
