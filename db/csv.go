package db

import "github.com/ezhdanovskiy/gocsvparser/model"

type Config struct {
	FilePath string
}

func InitDb(cfg Config) (*csvDb, error) {
	p := &csvDb{FilePath: cfg.FilePath}
	return p, nil
}

type csvDb struct {
	FilePath string
}

func (d *csvDb) SelectPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	people = append(people,
		&model.Person{Id: 1, First: "fffff1", Last: "lllll1"},
		&model.Person{Id: 2, First: "fffff2", Last: "lllll2"})
	return people, nil
}
