package main

import (
	"encoding/json"
	"errors"
	"messageServer/utility"
	"net/http"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

//initiating the map
func init() {
	mapLocker.dbMap = make(map[int]utility.Messages)
}

//structure to accomadate map with locking mechanism
var mapLocker struct {
	dbMap map[int]utility.Messages //map acting as database inhouse
	lock  sync.Mutex
}

//counter to be used for message id, may be we can use rand as well
var IdCounter int

/*
Func: AddMessage
Input values: http response, http request
Return values: -
Description: This function takes care of handling the add message http request
             and sending the ack back to client
*/
func AddMessage(w http.ResponseWriter, r *http.Request) {
	var mess utility.Messages
	log.Println("conroller..")
	err := json.NewDecoder(r.Body).Decode(&mess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	IdCounter++
	log.Debug("message received is:", mess)
	//err = utility.AddMessage(mess, db)
	mess.Id = IdCounter
	AddMessageMap(mess, IdCounter)
	// if err != nil {
	// 	log.Println("error", err)
	// }
	w.Write([]byte("Created...\n"))
}

/*
Func: ListMessage
Input values: http response, http request
Return values: -
Description: This function takes care of handling the list message http request
             and sending the json response back to client along with status code
*/
func ListMessage(w http.ResponseWriter, r *http.Request) {
	log.Info("ListMessage...")
	// //messages, err := utility.ListAllMessages(db)
	// if err != nil {
	// 	log.Println("error", err)
	// }
	log.Println("messages are", mapLocker.dbMap)
	json.NewEncoder(w).Encode(mapLocker.dbMap)
}

/*
Func: GetOneMessage
Input values: http response, http request
Return values: -
Description: This function takes care of handling the Get message http request
             and both based on messagID and messageDescription and sending the json response back to client along with status code
*/
func GetOneMessage(w http.ResponseWriter, r *http.Request) {
	log.Info("GetOneMessage...")
	messageId := r.FormValue("id")
	username := r.FormValue("username")
	var messages utility.Messages
	var err error
	switch {
	case messageId != "":
		log.Debug("getting message using message id")
		id, _ := strconv.Atoi(messageId)
		//messages, err = utility.GetOneIDMessage(id, db)
		messages, err = GetOneMessageMap(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		return
	case username != "":
		log.Debug("getting message using username")

		// log.Println("getting message using username")
		// messages, err = utility.GetOneuserMessage(username, db)
		// if err != nil {
		// 	log.Println("error", err)
		// }
		json.NewEncoder(w).Encode("yet to Implement...")
		return
	}

	log.Debug("messages are", messages)
	json.NewEncoder(w).Encode(messages)
}

/*
Func: DeleteOneMessage
Input values: http response, http request
Return values: -
Description: This function takes care of handling the Delete message http request
             and both based on messagID and sending the json response back to client along with status code
*/
func DeleteOneMessage(w http.ResponseWriter, r *http.Request) {
	messageId := r.FormValue("id")
	id, _ := strconv.Atoi(messageId)
	log.Println("messages id", messageId)
	err := DeleteOneMessageMap(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(mapLocker.dbMap)
}

/*
Name: AddMessageMap
Input values: message structure, message id
Return values: -
Description: This function takes care of adding the message to DB/map in this instance
*/
func AddMessageMap(m utility.Messages, id int) {
	log.Debug("AddMessageMap...", m)
	mapLocker.lock.Lock() //making sure no parallel write operations on this map
	mapLocker.dbMap[id] = m
	mapLocker.lock.Unlock()
}

/*
Name: GetOneMessageMap
Input values: message id
Return values: message structure, error
Description: This function takes care of getting the messages based on messageId
*/
func GetOneMessageMap(messageId int) (mess utility.Messages, err error) {
	log.Debug("GetOneMessageMap...", messageId)

	if _, ok := mapLocker.dbMap[messageId]; !ok {
		err = errors.New("message not found")
		log.Error(err)
	}
	if !utility.IsPalindrome(mess.Message) {
		err = errors.New("message not palindrome")
		return
	}
	mess = mapLocker.dbMap[messageId]
	return
}

/*
Name: DeleteOneMessageMap
Input values: message id
Return values: error
Description: This function takes care of deleting the messages based on messageId
*/
func DeleteOneMessageMap(messageId int) (err error) {
	log.Debug("GetOneMessageMap...", messageId)

	if _, ok := mapLocker.dbMap[messageId]; !ok {
		err = errors.New("message not found")
		log.Error(err)
	}
	delete(mapLocker.dbMap, messageId)
	return
}
