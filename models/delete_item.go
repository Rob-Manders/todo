package models

import (
	"errors"
	"log"
	"todo/store"
)

func DeleteItem(key string) (error) {
	if store.ListInitialised == false {
		return errors.New("List not yet initialised.")
	}

	log.Printf("Deleting '%s'.", key)

	store.ListInstance.Delete(key)

	return nil
}