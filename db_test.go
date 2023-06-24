package main

import (
	"testing"
)

func Test_select(t *testing.T) {
	db, err := newDB()
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT * from test_tables")
	if err != nil {
		panic(err)
	}
	var id string
	var name string
	if err := row.Scan(&id, &name); err != nil {
		panic(err)
	}
	if id == "id" && name == "name" {
		t.Log("success")
		return
	}
	t.Error("error", id, name)
}
