package cron

import "referralAPI/database"

func DeleteInTime() {
	query := "DELETE FROM codes WHERE created_at < DATE('now', '-1 day')"
	_, err := database.DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
