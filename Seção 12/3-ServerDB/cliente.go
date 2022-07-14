package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sId := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sId)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	case r.Method == "POST":
		addUsuario(w, r)
	case r.Method == "DELETE":
		delUsuario(w, r, id)
	case r.Method == "PUT":
		updUsuario(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db := createConnection()
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func addUsuario(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	var u Usuario
	json.NewDecoder(r.Body).Decode(&u)

	db.Query("insert into usuarios(nome) values(?)", u.Nome)

	json, _ := json.Marshal("Usuário inserido com sucesso")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func delUsuario(w http.ResponseWriter, r *http.Request, id int) {
	db := createConnection()
	defer db.Close()

	db.Query("delete from usuarios where id = ?", id)

	json, _ := json.Marshal("Usuário deletado com sucesso")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func updUsuario(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	var u Usuario
	json.NewDecoder(r.Body).Decode(&u)

	db.Query("update usuarios set nome = ? where id = ?", u.Nome, u.ID)

	json, _ := json.Marshal("Usuário atualizado com sucesso")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func createConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:84123@tcp/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
