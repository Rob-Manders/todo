package store

import "errors"

type Item struct {
	Title 	string	`json:"title"`
	Done 		bool		`json:"done"`
}

type List struct {
	Items 	map[string]Item
}

var ListInstance *List
var ListInitialised bool = false

func CreateList() *List {
	ListInstance = &List{Items: make(map[string]Item)}
	ListInitialised = true

	return ListInstance
}

func (l List) Get(key string) (Item, error) {
	if _, exists := l.Items[key]; !exists {
		return Item{}, errors.New("No item with key " + key + " in list.")
	}

	return l.Items[key], nil
}

func (l List) Set(key string, item Item) {
	l.Items[key] = item
}

func (l List) Delete(key string) {
	delete(l.Items, key)
}
