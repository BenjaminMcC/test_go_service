package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BenjaminMcC/test_go_service/httpHandlers"

	"github.com/BenjaminMcC/test_go_service/storage"
	"github.com/BenjaminMcC/test_go_service/structs"
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
	storage.Add(createMessage("Testing Again", "5678"))
	storage.Add(createMessage("Testing A Third Time", "9012"))

	log.Println("Attempting to start HTTP Server.")

	httpHandlers.HandleRequests()

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
	log.Println("here 1")

	log.Println("here")

}
