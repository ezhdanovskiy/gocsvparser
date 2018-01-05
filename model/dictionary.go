package model

type DictionaryItem struct {
	SourceCategory   string `json:"sourceCategory"`
	SourceContractor string `json:"sourceContractor"`
	Category         string `json:"category"`
	Contractor       string `json:"contractor"`
	Comment          string `json:"comment"`
}
