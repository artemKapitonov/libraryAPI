package repository

import (
	"fmt"

	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

func (r *Repository) NewUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) User(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)

	if err := r.db.Get(&user, query, username, password); err != nil {
		return user, err
	}

	return user, nil
}
