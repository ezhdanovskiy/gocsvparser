package model

import "log"

type Model struct {
	db
}

func New(db db) *Model {
	log.Printf("Create Model")
	return &Model{
		db: db,
	}
}

func (m *Model) Dictionary() ([]*DictionaryItem, error) {
	return m.GetDictionariItems()
}
