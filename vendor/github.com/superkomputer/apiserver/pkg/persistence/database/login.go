package database

import (
	"github.com/SuperKomputer/apiserver/pkg/rest/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(creds models.UserCredentials) error {
	// Salt and hash the password using the bcrypt algo
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	// Insert into db
	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		return err
	}
	return nil
}

func Login(creds models.UserCredentials) error {
	result := db.QueryRow("select password from users where username=$1", creds.Username)
	storedCreds := &models.UserCredentials{}
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		return err
	}
	return nil
}
