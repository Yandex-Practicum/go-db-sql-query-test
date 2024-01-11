package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	clientID := 1

	client, err := selectClient(db, clientID)
	if err != nil {
		t.Fatal(err)
	}
	if client.ID != clientID {
		t.Error("Unequal IDs")
	}
	if client.FIO == "" {
		t.Fatal("empty fio")
	}
	if client.Login == "" {
		t.Fatal("empty login")
	}

	if client.Birthday == "" {
		t.Fatal("empty birthday")
	}

	if client.Email == "" {
		t.Fatal("empty email")
	}
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := -1

	client, err := selectClient(db, clientID)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
	if client.ID != 0 {
		t.Fatal("not empty ID")
	}
	if client.FIO != "" {
		t.Fatal("not empty fio")
	}
	if client.Login != "" {
		t.Fatal("not empty login")
	}

	if client.Birthday != "" {
		t.Fatal("not empty birthday")
	}

	if client.Email != "" {
		t.Fatal("not empty email")
	}
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("empty id")
	}
	client, err := selectClient(db, id)
	if err != nil {
		t.Fatal(err)
	}

	if cl.FIO != client.FIO {
		t.Fatal("not equal clients")
	}
	if cl.Login != client.Login {
		t.Fatal("not equal clients")
	}
	if cl.Email != client.Email {
		t.Fatal("not equal clients")
	}
	if cl.Birthday != client.Birthday {
		t.Fatal("not equal clients")
	}
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("empty id")
	}
	_, err = selectClient(db, id)
	if err != nil {
		t.Fatal(err)
	}
	err = deleteClient(db, id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = selectClient(db, id)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
}
