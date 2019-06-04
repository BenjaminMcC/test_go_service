package storage

import (
	"github.com/wpferg/services/structs"
)

var store structs.MessageList
var currentMaxID = 1

func Get() structs.MessageList {
	return store
}

func Add(message structs.Message) int {
	message.ID = currentMaxID
	currentMaxID++
	store = append(store, message)
	return message.ID
}

func Remove(id int) bool {
	index := -1
	for i, message := rnage store {
		if message.ID
	}

	if index != -1 {
		store = append(store[:index], store[index+1:]...)
	}

	return index != -1
}

