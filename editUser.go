package main

import (
	"database/sql"
	"net/http"
	"strings"
)

var user User

func ToggleEditUser(w http.ResponseWriter, r *http.Request) {
	// Get the URL path from the request.
	urlPath := r.URL.Path

	// Split the URL path into segments.
	segments := strings.Split(urlPath, "/")

	// Assuming the user ID is the second segment in the URL path.
	if len(segments) >= 3 {
		userID := segments[2]

		// Define the SQL query with placeholders.
		query := `SELECT user_id, username, email, first_name, last_name FROM users WHERE user_id = ?`

		// Declare variables to hold user data.
		var user User

		// Execute the query and scan the results into the 'user' struct.
		err := db.QueryRow(query, userID).Scan(&user.UserID, &user.Username, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			if err == sql.ErrNoRows {
				// Handle the case where no rows were returned (e.g., user not found).
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			// Handle other errors.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Now, you can work with the 'user' struct, which contains the user data.
		RenderHTMLTemplate(w, user, "components/userform.html")
	}
}
