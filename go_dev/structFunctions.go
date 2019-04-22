package go_dev

import (
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
)


func PopulateUserPage(username string, db *sql.DB) *UserPage {
	var page *UserPage

	page.info = *GetUserInfo(username, db)

	page.projects = GetProjects(username, db)

	page.tasks = GetUserTasks(username, db)

	page.pins = GetUserPins(username, db)

	page.feed = GetUserFeed(username, db)

	return page
}

func PopulateProjectPage(id int,  db *sql.DB) Project {
	var thing Project

	return thing
}
