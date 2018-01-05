package db

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"github.com/ezhdanovskiy/gocsvparser/model"
)

type Config struct {
	FilePath string
}

func InitDb(cfg Config) (*csvDb, error) {
	log.Printf("InitDb('%s')", cfg.FilePath)
	p := &csvDb{FilePath: cfg.FilePath}
	p.readDictionary()
	return p, nil
}

type csvDb struct {
	FilePath string
	dict     []*model.DictionaryItem
}

func (d *csvDb) SelectPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	people = append(people,
		&model.Person{Id: 1, First: "fffff1", Last: "lllll1"},
		&model.Person{Id: 2, First: "fffff2", Last: "lllll2"})
	return people, nil
}

func (d *csvDb) GetDictionariItems() ([]*model.DictionaryItem, error) {
	return d.dict, nil
}

func (d *csvDb) readDictionary() {
	csvFile, _ := os.Open(d.FilePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		d.dict = append(d.dict, &model.DictionaryItem{
			SourceCategory:   line[0],
			SourceContractor: line[1],
			Category:         line[2],
			Contractor:       line[3],
			Comment:          line[4]})
	}
	dictJson, _ := json.Marshal(d.dict)
	fmt.Println(string(dictJson))
}

func decodeString(s string) string {
	sr := strings.NewReader(s)
	tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		log.Println(err)
		return s
	}
	return string(buf)
}
