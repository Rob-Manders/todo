package models

import (
	"errors"
	"log"
	"todo/store"
)

func CreateItem(item store.Item) (store.Item, error) {
	if store.ListInitialised == false {
		return store.Item{}, errors.New("List not yet initialised.")
	}

	store.ListInstance.Set(item.Title, item)
	newItem, error := store.ListInstance.Get(item.Title)

	if error != nil {
		return store.Item{}, errors.New("Created item could not be verified.")
	}

	log.Printf("Item %s created.", newItem.Title)

	return newItem, nil
}
