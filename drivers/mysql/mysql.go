package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type connect struct {
	Database *sql.DB
}

var App connect

func InitMysql() *sql.DB {
	godotenv.Load()
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")
	host := os.Getenv("DATABASE_HOST")
	migration := os.Getenv("MIGRATION")

	log.Println(username, password, database, host)

	dsn := fmt.Sprintf("%v:%v@tcp(%s)/%v?parseTime=true", username, password, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}

	if migration == "true" {
		Migration(db)
	}

	return db
}

func Migration(db *sql.DB) {
	_, err := db.Exec(`DROP TABLE IF EXISTS cake;`)
	if err != nil {
		log.Println("error drop table", err)
	}

	_, err = db.Exec(`CREATE TABLE cake (
		id int NOT NULL AUTO_INCREMENT,
		title text NULL,
		description text NULL,
		rating numeric NULL,
		image varchar(255) NULL,
		created_at timestamp NULL DEFAULT now(),
		updated_at timestamp NULL DEFAULT now(),
		deleted_at timestamp NULL,
		PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Println("error create table", err)
	}
}
