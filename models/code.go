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

func GetCodeByEmail(email string) (string, error) {
	query := "SELECT code FROM codes WHERE user_id IN (SELECT id FROM users WHERE email =?)"
	row := database.DB.QueryRow(query, email)
	var code string
	err := row.Scan(&code)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return code, nil
}

func GetCodeIdByCode(code string) (int64, error) {
	query := "SELECT id FROM codes WHERE code =?"
	row := database.DB.QueryRow(query, code)
	var id int64
	err := row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return id, nil
}
