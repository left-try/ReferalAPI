package models

import (
	"awesomeProject/database"
	"awesomeProject/utils"
	"database/sql"
	"errors"
)

type Code struct {
	Id     int64
	Code   string `binding:"required"`
	UserId int64  `binding:"required"`
}

func (code *Code) Create() error {
	code.Code = utils.GenerateCode(code.UserId)
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

func (code *Code) GetByEmail(email string) (*Code, error) {
	query := "SELECT id, code FROM codes WHERE user_id IN (SELECT id FROM users WHERE email =?)"
	row := database.DB.QueryRow(query, email)
	var codeCode string
	err := row.Scan(&code.Id, &codeCode)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	code.Code = codeCode
	return code, nil
}
