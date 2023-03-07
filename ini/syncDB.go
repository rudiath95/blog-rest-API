package ini

import "github.com/rudiath95/blog-rest-API/models"

func SyncDatabases() {
	DB.AutoMigrate(
		models.User{},
		models.UserInfo{},

		models.Blog{},
		models.Comment{},
	)
}
