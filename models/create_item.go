package models

import (
	"errors"
	"log"
	"todo/store"
	"github.com/gosimple/slug"
)

func CreateItem(item store.Item) (store.Item, error) {
	if store.ListInitialised == false {
		return store.Item{}, errors.New("List not yet initialised.")
	}

	key := slug.Make(item.Title)

	log.Printf("Creating '%s'", item.Title)

	store.ListInstance.Set(key, item)
	newItem, error := store.ListInstance.Get(item.Title)

	if error != nil {
		return store.Item{}, errors.New("Created item could not be verified.")
	}


	return newItem, nil
}
