package models

import (
	"awesomeProject/database"
	"time"
)

type Event struct {
	Id       int64
	Name     string `binding:"required"`
	Desc     string `binding:"required"`
	Location string `binding:"required"`
	Date     time.Time
	UserId   int64
}

func (event *Event) Save() error {
	query := "INSERT INTO events(name, desc, location, date, user_id) VALUES (?, ?, ?, ?, ?)"
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	exec, err := prepare.Exec(event.Name, event.Desc, event.Location, event.Date, event.UserId)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	event.Id = id
	return err
}

func (event *Event) Update() error {
	query := `
		UPDATE events
		SET	name = ?, desc = ?, location = ?, date = ?
		WHERE id = ?
	`
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(event.Name, event.Desc, event.Location, event.Date, event.Id)
	return err
}

func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(event.Id)
	return err
}

func (event *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(event.Id, userId)
	return err
}

func (event *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	prepare, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(event.Id, userId)
	return err
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Desc, &event.Location, &event.Date, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Desc, &event.Location, &event.Date, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}
