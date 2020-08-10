package utility

import (
	"database/sql"
	"fmt"
	"log"
	//	"messageServer/db"
)

//Message struct is the json data of message to be saved
type Messages struct {
	Id          int    `json:"id,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	User        string `json:"user,omitempty"`
}

/*
Name: AddMessage
Input values: message structure, db connection string
Return values: -
Description: This function takes care of adding the message to DB in this instance
*/
func AddMessage(m Messages, db *sql.DB) error {
	log.Println("Adding Message...")
	// sqlStatement := `
	// INSERT INTO message (id, "user", message, description,now())
	// VALUES ($1, $2, $3, $4, $5)`
	sql := fmt.Sprintf(`INSERT INTO public.message("user", message, description, created_at) VALUES ('%s', '%s', '%s', now());`, m.User, m.Message, m.Description)
	log.Println("query formed is", sql)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

/*
Name: ListAllMessages
Input values: db connection string
Return values: message structure, error
Description: This function takes care of getting all the messages
*/
func ListAllMessages(db *sql.DB) (mess []Messages, err error) {
	log.Println("Listing Message...")
	rows, err := db.Query("SELECT * from public.message")
	if err != nil {
		// handle this error better than this
		return mess, err
	}
	defer rows.Close()
	for rows.Next() {
		var message Messages
		if err = rows.Scan(&message.Id, &message.User, &message.Message, &message.Description, &message.CreatedAt); err != nil {
			return mess, err
		}
		mess = append(mess, message)
		fmt.Println("messsages retrieved are ", len(mess), mess)
	}
	return mess, err

}

/*
Name: GetOneIDMessage
Input values: message id,db connection string
Return values: message structure, error
Description: This function takes care of getting the messages from DB based on messageId
*/
func GetOneIDMessage(id int, db *sql.DB) (mess Messages, err error) {
	log.Println("Getting one Message...", id)
	statement := fmt.Sprintf("SELECT id,user,message,description,created_at from public.message where id='%d'", id)
	log.Println("query formed is", statement)
	rows := db.QueryRow(statement)
	var message Messages

	if err = rows.Scan(&message.Id, &message.User, &message.Message, &message.Description, &message.CreatedAt); err != nil {
		log.Println("Error", err)
		return mess, err
	}
	//	mess = append(mess, message)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		log.Println(mess)
	default:
		panic(err)
	}

	return mess, err
}

/*
Name: GetOneuserMessage
Input values: username,db connection string
Return values: message structure, error
Description: This function takes care of getting the messages based on username from DB
*/
func GetOneuserMessage(username string, db *sql.DB) (mess Messages, err error) {
	log.Println("Getting one Message...", username)
	statement := fmt.Sprintf("SELECT * from public.message where user='%s'", username)
	log.Println("query formed is", statement)
	rows, err := db.Query(statement)
	if err != nil {
		log.Println("Error", err)
		return mess, err
	}
	var message Messages

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&message.Id, &message.User, &message.Message, &message.Description, &message.CreatedAt); err != nil {
			log.Println("Error", err)
			return mess, err
		}
		//	mess = append(mess, message)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			log.Println(mess)
		default:
			panic(err)
		}
	}
	return mess, err
}

/*
Name: IsPalindrome
Input values: message string
Return values: Boolean(true/false)
Description: This function takes care of checking a string is palindrome or not
*/

func IsPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
