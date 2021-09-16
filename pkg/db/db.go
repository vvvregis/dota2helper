package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"

	"github.com/vvvregis/dota2-helper/pkg/config"
)

func connect(config config.Config) *sql.DB {

	connectionString := config.Database.User + ":" + config.Database.Password + "@/" + config.Database.DbName
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db

}

func Insert(config config.Config, id string) {
	db := connect(config)

	defer db.Close()
	query := "INSERT INTO `enemies` (`steam_id`, `date`) VALUES (?,?);"

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	currentDate := time.Now().Format("2006-01-02")

	_, error := statement.Exec(id, currentDate)

	if error != nil {
		log.Fatalln(error)
	}
}
