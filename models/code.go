package models

import (
	"database/sql"
	"errors"
	"referralAPI/database"
	"referralAPI/utils"
)

type Code struct {
	Id     int64
	Code   string `binding:"required"`
	UserId int64
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

func GetUserIdByCode(code string) (int64, error) {
	query := "SELECT userId FROM codes WHERE code =?"
	row := database.DB.QueryRow(query, code)
	var userId int64
	err := row.Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetReferrals(userId int64) ([]int64, error) {
	query := "SELECT user_id FROM users WHERE referrer_id =?"
	rows, err := database.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var referrals []int64
	for rows.Next() {
		var referrerId int64
		err := rows.Scan(&referrerId)
		if err != nil {
			return nil, err
		}
		referrals = append(referrals, referrerId)
	}
	return referrals, nil
}
