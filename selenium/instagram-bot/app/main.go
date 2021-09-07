package main

import (
	"app/business.go"
	"app/db/migration"
)

func main() {
	migration.AutoMigration()
	business.CollectUserInfo()
}
