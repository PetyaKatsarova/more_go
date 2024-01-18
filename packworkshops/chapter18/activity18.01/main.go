package main

import (
	"crypto/sha512"
	"database/sql"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	Id       string
	Password string
}

var testData = []*UserInfo{
	{
		Id:       "1",
		Password: "1234",
	},
	{
		Id:       "2",
		Password: "5678",
	},
}

func initializeDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS USER_INFO (USER_ID TEXT, PASSWORD TEXT)`)
	if err != nil { return err }
	stmt, err := db.Prepare(`INSERT INTO USER_INFO (USER_ID, PASSWORD) VALUES(?, ?)`)
	if err != nil { return err }
	for _, user := range testData {
		_, err := stmt.Exec(user.Id, user.Password)
		if err != nil {
			return err
		}
	}
	return nil
}

func tearDownDB(db * sql.DB) error {
	_, err := db.Exec("DROP TABLE USER_INFO")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getConnection() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "test.DB")
	if err != nil { return nil, fmt.Errorf("could not open db connection %v\n", err) }
	return conn, nil
}

func UpdatePassword(db *sql.DB, Id string, password string) error {
	query 			:= `UPDATE USER_INFO SET PASSWORD=? WHERE USER_ID=?`
	hashedPassword	:= sha512.Sum512([]byte(password))
	result, err     := db.Exec(query, string(hashedPassword[:]), Id)
	if err != nil { return err }

	rows, err 		:= result.RowsAffected()
	if err != nil { return err }

	fmt.Printf("storing encrypted password: \n%x\n", string(hashedPassword[:]))
	if rows == 0 {
		return fmt.Errorf("no row affected")
	}
	if rows > 1 {
		return fmt.Errorf("more than one row affected: %d", rows)
	}
	return nil
}

func GetPassword(db *sql.DB, userID string) (resp []byte, err error) {
	query := `SELECT PASSWORD FROM USER_DETAILS WHERE USER_ID = ?`
	row   := db.QueryRow(query, userID)
	switch err = row.Scan(&resp); err {
	case sql.ErrNoRows:
		return resp, fmt.Errorf("no rows returned")
	default:
		return resp, err
	}
}
/*
cipher[:] is a slicing operation applied to an array or a slice. It creates a new slice that references the entire underlying array,
 effectively making a copy of the data. 
string(cipher[:]) converts the cipher byte array (hash) into a string. However, note that this conversion may not produce a 
meaningful string representation of the hash since cryptographic hash values are typically binary data.
*/

func main() {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1) 
		/* os.Exit(1)
		It is a low-level mechanism for abruptly ending the program's execution and is typically used when you want to forcefully
		terminate the entire program, including all goroutines and resources, without any cleanup or additional actions.
		It's commonly used in command-line utilities or scripts to indicate a failure or error condition.
		*/
	}
	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tearDownDB(db) // execute just b4 the func returns: used for cleanup, resource management, etc
	err = UpdatePassword(db, "1", "NewPassword")
	if err != nil {
		fmt.Println("error updating password: ", err)
	}
	fmt.Println("retrieving hashed password from db")
	password, err := GetPassword(db, "1")
	if err != nil {
		fmt.Println("error retrieving password: ", err)
	}
	fmt.Println("checking password match")
	newPwdHash := sha512.Sum512([]byte("NewPassword"))
	if string(newPwdHash[:]) == string(password[:]) {
		fmt.Println("successful pwd match")
	}
}