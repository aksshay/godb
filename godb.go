package main

import (
  "fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"os"
  "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

// type dsn struct {
// 	driver string,
// 	user string,
// 	password string,
// 	protocol string,
// 	address string,
// 	port int,
// 	database string,
// 	parameters string
// }

type Credentials struct {
	Database string `yaml:"name"`
	Password string `yaml:"mysql"`
}

func constructDsn(database string, password string) (dsn string) {
	return fmt.Sprintf("%v:%v@tcp(%v)/%v",
	  "root",
	  password,
	  "db",
		database,
	)
}

func main() {
  //db, err := sql.Open("mysql", "user:password@/dbname")
	var credsfile string
	if os.Getenv("creds_file") != "" {
		credsfile = os.Getenv("creds_file")
	}	else {
		credsfile = "/creds"
	}

	content, err := ioutil.ReadFile(credsfile)
	if err != nil {
		log.Fatal(err)
	}

	var creds Credentials
	err = yaml.Unmarshal(content, &creds)
	if err != nil {
		log.Fatal(err)
	}

	dsn := constructDsn(creds.Database, creds.Password)

	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test dsn connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

  rows, err := db.Query("SELECT user, host FROM mysql.user;")
	defer rows.Close()

	fmt.Printf(string(rows))


}
