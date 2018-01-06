package model

type db interface {
	GetDictionariItems() ([]*DictionaryItem, error)
}
