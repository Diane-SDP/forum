package models

import (
	"database/sql"
	"fmt"
)

func CountComment(idpost int, db *sql.DB) string {
	query, err := db.Prepare("SELECT COUNT(*) FROM comment WHERE idPost = ?")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer query.Close()
	var output string
	err = query.QueryRow(idpost).Scan(&output)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return output
}
