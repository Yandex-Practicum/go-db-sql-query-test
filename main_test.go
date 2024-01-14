package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

type DB struct {
	db *sql.DB
}

var mainDB DB

func TestMain(m *testing.M) {
	fmt.Println("Setting up database...")
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	mainDB.db = db
	fmt.Println("Let's test some tests...")
	code := m.Run()
	mainDB.db.Close()
	os.Exit(code)
}

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Test_SelectClient_WhenOk(t *testing.T) {
	clientID := 1

	client, err := selectClient(mainDB.db, clientID)
	if err != nil {
		t.Fatal(err)
	}
	if client.ID == 0 {
		t.Error("empty client ID")
	}
	if client.ID != clientID {
		t.Errorf("client id wrong, want: %d, got: %d", clientID, client.ID)
	}
	if client.FIO == "" {
		t.Error("empty fio")
	}
	if client.Login == "" {
		t.Error("empty login")
	}

	if client.Birthday == "" {
		t.Error("empty birthday")
	}

	if client.Email == "" {
		t.Error("empty email")
	}
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	clientID := -1

	client, err := selectClient(mainDB.db, clientID)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
	if client.ID != 0 {
		t.Errorf("client id wrong, want: 0, got: %d", client.ID)
	}
	if client.FIO != "" {
		t.Errorf("client fio wrong, want: , got: %s", client.FIO)
	}
	if client.Login != "" {
		t.Errorf("client login wrong, want: , got: %s", client.Login)
	}

	if client.Birthday != "" {
		t.Errorf("client birthday wrong, want: , got: %s", client.Birthday)
	}

	if client.Email != "" {
		t.Errorf("client email wrong, want: , got: %s", client.Email)
	}
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	id, err := insertClient(mainDB.db, cl)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("empty id")
	}
	client, err := selectClient(mainDB.db, id)
	if err != nil {
		t.Fatal(err)
	}

	if cl.FIO != client.FIO {
		t.Error("not equal clients")
	}
	if cl.Login != client.Login {
		t.Error("not equal clients")
	}
	if cl.Email != client.Email {
		t.Error("not equal clients")
	}
	if cl.Birthday != client.Birthday {
		t.Error("not equal clients")
	}
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(mainDB.db, cl)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("empty id")
	}
	_, err = selectClient(mainDB.db, id)
	if err != nil {
		t.Fatal(err)
	}
	err = deleteClient(mainDB.db, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = selectClient(mainDB.db, id)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
}
