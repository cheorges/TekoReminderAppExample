package repository

import (
	"fmt"
	"reminders/pkg/model"
)

var nextId = 1

type InMemoryStorageImpl struct {
	reminders *[]model.Reminder
}

func NewInMemoryStorageImpl() *InMemoryStorageImpl {
	nextId = 1
	reminders := make([]model.Reminder, 0)
	return &InMemoryStorageImpl{reminders: &reminders}
}

func (storage InMemoryStorageImpl) Get(id int) (*model.Reminder, error) {
	for _, reminder := range *storage.reminders {
		if id == reminder.Id {
			return &reminder, nil
		}
	}
	return nil, fmt.Errorf("id not exist")
}

func (storage InMemoryStorageImpl) GetAll() []model.Reminder {
	return *storage.reminders
}

func (storage InMemoryStorageImpl) Delete(id int) bool {
	for i, reminder := range *storage.reminders {
		if id == reminder.Id {
			*storage.reminders = append((*storage.reminders)[:i], (*storage.reminders)[i+1:]...)
			return true
		}
	}
	return false
}

func (storage InMemoryStorageImpl) Create(task string) (int, error) {
	reminder := model.New(task)
	reminder.Id = nextId
	nextId++
	*storage.reminders = append(*storage.reminders, reminder)
	return reminder.Id, nil
}

func (storage InMemoryStorageImpl) Update(id int) (*model.Reminder, error) {
	for i, reminder := range *storage.reminders {
		if id == reminder.Id {
			(&(*storage.reminders)[i]).Toggle()
			return &(*storage.reminders)[i], nil
		}
	}
	return nil, fmt.Errorf("id not exist")
}
