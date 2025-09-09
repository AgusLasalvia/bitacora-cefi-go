package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func init() {
	databaseUrl := "bitacora.db"
	fmt.Println("Conectando a:", databaseUrl)

	var err error
	Database, err = sql.Open("sqlite3", databaseUrl)
	if err != nil {
		log.Fatal("Error creando conexión: ", err)
	}

	// Validar conexión real
	if err := Database.Ping(); err != nil {
		log.Fatal("Error conectando a SQLite: ", err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS equipment_data( 
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
		name TEXT NOT NULL,
		lab TEXT NOT NULL,
		equipment TEXT NOT NULL,
		startDateTime TEXT NOT NULL,
		endDateTime TEXT,
		received TEXT,
		returned TEXT,
		comments TEXT,
		timestamp TEXT
	);`

	_, err = Database.Exec(createTables)
	if err != nil {
		log.Fatal("Error creando tabla principal: ", err)
	}

	defer Database.Close()

	log.Println("✅ Base de datos conectada con éxito y tablas listas")
}
