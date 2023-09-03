package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reminders/pkg/model"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "reminder"
)

type PostgresqlStorageImpl struct {
	db *sql.DB
}

func NewPostgresStorageImpl() *PostgresqlStorageImpl {
	postgresqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlConnection)
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &PostgresqlStorageImpl{db: db}
}

func (storage *PostgresqlStorageImpl) Get(id int) (*model.Reminder, error) {
	var reminder model.Reminder
	row := storage.db.QueryRow("SELECT id, task, done FROM public.reminder WHERE id = $1", id)
	err := row.Scan(&reminder.Id, &reminder.Task, &reminder.Done)
	if err != nil {
		return nil, err
	}
	return &reminder, nil
}

func (storage *PostgresqlStorageImpl) GetAll() []model.Reminder {
	var reminders = make([]model.Reminder, 0)
	rows, err := storage.db.Query("SELECT id, task, done FROM public.reminder")
	if err != nil {
		return reminders
	}

	for rows.Next() {
		var r model.Reminder
		err = rows.Scan(&r.Id, &r.Task, &r.Done)
		if err != nil {
			log.Println("Could not read values from row")
			return reminders
		}
		reminders = append(reminders, r)
	}

	return reminders
}

func (storage *PostgresqlStorageImpl) Delete(id int) bool {
	_, err := storage.db.Exec("DELETE FROM public.reminder WHERE id = $1", id)
	return err == nil
}

func (storage *PostgresqlStorageImpl) Create(task string) (int, error) {
	result, err := storage.db.Exec("INSERT INTO public.reminder (task, done) VALUES ($1, $2)", task, false)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (storage *PostgresqlStorageImpl) Update(id int) (*model.Reminder, error) {
	result, err := storage.db.Exec("UPDATE public.reminder SET done = not done WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return storage.Get(int(insertId))
}
