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
	query := "INSERT INTO codes(code.Code, user_id) VALUES (?,?)"
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

	exec, err := prepare.Exec(code.Code, userId)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	code.Id = id
	return nil
}
