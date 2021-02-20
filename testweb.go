package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
)

// https://gowebexamples.com/hello-world/
// http://localhost/test
// go install code.google.com/p/odbc
// https://sqlchoice.azurewebsites.net/en-us/sql-server/developer-get-started/go/rhel/step/2.html
// https://www.easysoft.com/support/kb/kb01045.html

func testweb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}

func testweb2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":80", nil)
}

// AYWVPSQL11\SQLP224
// sqlmsop02.sce.eix.com,1989
// https://laptrinhx.com/golang-package-denisenkom-go-mssqldb-1257563911/

func testsql() {

<<<<<<< HEAD
	condb, errdb := sql.Open("odbc", "DSN=dw;")
=======
	condb, errdb := sql.Open("odbc", "DSN=dw;Trusted_Connection=True;")
>>>>>>> 66c708fda119c48828a70547d5f0c7d5d5c9257e
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	var (
		sn   string
		isvc int
	)

<<<<<<< HEAD
	rows, err := condb.Query("select top 10 name,age from dbo.test")
=======
	rows, err := condb.Query("select top 10 namge,age from dbo.test")
>>>>>>> 66c708fda119c48828a70547d5f0c7d5d5c9257e

	if err != nil {

		fmt.Println("getting error")
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&sn, &isvc)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(sn, isvc)
	}

	defer condb.Close()

}
