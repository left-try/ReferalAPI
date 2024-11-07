package models

import (
	"awesomeProject/database"
	"awesomeProject/utils"
	"database/sql"
)

type Code struct {
	Id     int64
	Code   string `binding:"required"`
	UserId int64  `binding:"required"`
}

func (code *Code) Create(userId int64) error {
	code.Code = utils.GenerateCode(userId)
	code.UserId = userId
	query := "INSERT INTO codes(code) VALUES (?,?)"
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

	exec, err := prepare.Exec(code.Code)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	code.Id = id
	return nil
}

func (code *Code) Delete() error {
	query := "DELETE FROM codes WHERE id =?"
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

	_, err = prepare.Exec(code.Id)
	return err
}
