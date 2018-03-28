package database

import (
	"github.com/fanzhangio/superkomputer/pkg/rest/models"
)

const (
	// UsersTable is table of user in db
	UsersTable = "users"

	CreateUsersTable = `
create table users (
  email text,
  firstname text,
  lastname text,
  username text primary key,
  id text,
  password text
);
`
)

// ListUsers lists all users in users table
func ListUsers() (result []*models.User, err error) {
	rows, err := db.Query("SELECT * FROM " + UsersTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result = make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.Username, &user.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
