package main

//import package and library go
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB

var err error

var tpl *template.Template

//conect db and set template
func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/animal")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//function checked error
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//main function first execute
func main() {

	defer db.Close()
	//get from index() function
	http.HandleFunc("/", index)

	//get from userForm() function
	http.HandleFunc("/userForm", userForm)

	//get from createUsers() function
	http.HandleFunc("/createUsers", createUsers)

	//get from editUsers() function
	http.HandleFunc("/editUsers", editUsers)

	//get from updateUsers() function
	http.HandleFunc("/updateUsers", updateUsers)

	//get from deleteUsers() function
	http.HandleFunc("/deleteUsers", deleteUsers)

	//run server in 127.0.0.1:8080
	log.Println("Server is up on 8080 port")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

//deklarasi users variabel properti
type user struct {
	ID    int64
	Name  string
	Class string
	Legs  string

}

//list user
func index(w http.ResponseWriter, req *http.Request) {
	rows, e := db.Query(
		`SELECT id,
		name,
		class,
		legs,
		FROM animal;`)

	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	animal := make([]animal, 0)
	for rows.Next() {
		anm := animal{}
		rows.Scan(&anm.ID, &anm.Name, &anm.Class, &anm.Legs)
		animal = append(animal, anm)
	}
	log.Println(animal)
	tpl.ExecuteTemplate(w, "index.html", animal)
}

//form create user
func userForm(w http.ResponseWriter, req *http.Request) {
	err = tpl.ExecuteTemplate(w, "userForm.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//action create users
func createUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		anm := anm{}
		anm.Name = req.FormValue("name")
		anm.FirstName = req.FormValue("class")
		usr.Legs = req.FormValue("legs")
		
//

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		

		_, e = db.Exec("INSERT INTO animal (name,class, legs) VALUES (?,?,?)", anm.Name,
			anm.Class,
			anm.Legs,
		)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Methof not supported", http.StatusMethodNotAllowed)
}

//form edit data
func editUsers(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	rows, err := db.Query(
		`SELECT id,
	 	name,
		class,
		legs
		FROM animal
		WHERE id = ` + id + `;`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	anm := animal{}
	for rows.Next() {
		rows.Scan(&anm.ID, &anm.Name, &anm.Class, &anm.Legs)
	}
	tpl.ExecuteTemplate(w, "editUser.html", anm)
}

//action edit users
func updateUsers(w http.ResponseWriter, req *http.Request) {
	_, er := db.Exec("UPDATE animal set name = ?, class = ?, legs = ? WHERE id = ?",
		req.FormValue("name"),
		req.FormValue("class"),
		req.FormValue("legs"),
		req.FormValue("id"),
	)

	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

//action deleted users
func deleteUsers(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")

	if id == "" {
		http.Error(w, "Please Send ID", http.StatusBadRequest)
		return
	}

	_, er := db.Exec("DELETE FROM animal WHERE id = ?", id)

	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
