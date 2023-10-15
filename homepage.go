package main

import (
	"net/http"
)

type User struct {
	UserID        int    `db:"user_id"`
	Username      string `db:"username"`
	Email         string `db:"email"`
	Password      string `db:"password"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	DateOfBirth   string `db:"date_of_birth"`
	Location      string `db:"location"`
	AccountStatus int    `db:"account_status"`
	CreatedAt     string `db:"created_at"`
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.Location,
			&user.AccountStatus,
			&user.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)

	}

	RenderHTMLTemplate(w, users, "pages/index.html", "components/header.html", "components/aside.html")
}
