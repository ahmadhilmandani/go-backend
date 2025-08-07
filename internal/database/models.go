package database

import "database/sql"

type Models struct {
	User      UserModel
	Event     EventModel
	Attendees AttendeesModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User:      UserModel{DB: db},
		Event:     EventModel{DB: db},
		Attendees: AttendeesModel{DB: db},
	}
}
