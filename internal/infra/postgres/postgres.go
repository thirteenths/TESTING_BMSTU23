package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/thirteenths/test-bmstu23/internal/domain"
)

type Postgres struct {
	db *sql.DB
}

func NewMockPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func NewPostgres(databaseURL string) (*Postgres, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

const getAllEventQuery = "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS"

func (p *Postgres) GetAllEvent() ([]domain.Event, error) {
	var events []domain.Event

	rows, err := p.db.Query(getAllEventQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event domain.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

const getEventQuery = "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS WHERE ID=$1"

func (p *Postgres) GetEvent(id int) (domain.Event, error) {
	var event domain.Event

	err := p.db.QueryRow(getEventQuery, id).Scan(&event.ID, &event.Name, &event.Description, &event.Date)
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

const getUserEventQuery = "SELECT ID, NAME, DESCRIPTION, DATE FROM EVENTS WHERE ID=$1"

func (p *Postgres) GetUserEvent(idUser int) ([]domain.Event, error) {
	return []domain.Event{}, nil
}

const getUserQuery = "SELECT ID, EMAIL, HASH FROM USERS WHERE EMAIL=$1"

func (p *Postgres) GetUser(id int) (domain.User, error) {
	var user domain.User

	err := p.db.QueryRow(getUserQuery, id).Scan(&user.ID, &user.Email, &user.Hash)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

const updatePasswordQuery = "UPDATE USERS SET HASH=$1 WHERE EMAIL=$2"

func (p *Postgres) UpdatePassword(user domain.User) error {
	_, err := p.db.Exec(updatePasswordQuery, user.Hash, user.Email)
	return err
}

const getPasswordQuery = "SELECT ID, EMAIL, HASH FROM USERS WHERE EMAIL=$1"

func (p *Postgres) GetPassword(email string) (domain.User, error) {
	var user domain.User

	err := p.db.QueryRow(getPasswordQuery, email).Scan(&user.ID, &user.Email, &user.Hash)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

const createEventQuery = "INSERT INTO EVENTS (NAME, DESCRIPTION, DATE) VALUES ($1, $2, $3) RETURNING ID"

func (p *Postgres) CreateEvent(event domain.Event) (int, error) {
	var id int
	err := p.db.QueryRow(createEventQuery, event.Name, event.Description, event.Date).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

const deleteEventQuery = "DELETE FROM EVENTS WHERE ID=$1"

func (p *Postgres) DeleteEvent(id int) error {
	_, err := p.db.Exec(deleteEventQuery, id)
	return err
}

const updateEveryQuery = "UPDATE EVENTS SET NAME=$1, DESCRIPTION=$2, DATE=$3 WHERE ID=$4"

func (p *Postgres) UpdateEvent(event domain.Event, id int) error {
	_, err := p.db.Exec(updateEveryQuery, event.Name, event.Description, event.Date, id)
	return err
}
