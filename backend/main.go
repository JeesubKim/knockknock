package main

import (
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMySQLStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := NewAPIServer(":8080", store)
	api.Serve()

}

// https://www.youtube.com/watch?v=2JNUmzuBNV0
// https://www.youtube.com/watch?v=H7tbjKFSg58
