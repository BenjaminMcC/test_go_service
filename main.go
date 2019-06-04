package main

import (
	"log"
	"strconv"

	"github.com/wpferg/services/structs"
)

const PORT = 8080

var messageId = 0

func createMessage(message string, sender string) structs.Message {
	messageId++
	return structs.Message{
		ID:      messageId,
		Sender:  sender,
		Message: message,
	}
}

func main() {
	log.Println("Creating dummy messages")

	storage.Add(createMessage("Testing", "1234"))
	storage.Add(createMessage("Testing Again", "5679"))
	storage.Add(createMessage("Testing A Third Time", "9012"))

	http.HandleFunc("/", httpHandlers.HandleRequest)

	var err = http.ListenandServer(":"+strconv.Itoa(PORT), nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
