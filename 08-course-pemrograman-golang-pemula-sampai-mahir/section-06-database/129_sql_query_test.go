package belajar_golang_database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func SQLQueryGetConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Second)

	return db
}

func TestSQLQuery(t *testing.T) {
	db := SQLQueryGetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := `select * from customer`
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id string
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		message := fmt.Sprintf("ID: %s, Name: %s", id, name)
		fmt.Println(message)
	}
	defer rows.Close()
}

func SQLQueryParamConnection(driver string, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}

func TestSQLQueryWithStruct(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := SQLQueryParamConnection(driver, source)
	defer db.Close()

	type Customer struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var customers []Customer
	ctx := context.Background()
	query := `select id, name from customer order by id asc`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
			continue
		}
		customers = append(customers, customer)
	}

	jsonCustomers, _ := json.Marshal(customers)
	fmt.Println("Data customer: ", string(jsonCustomers))
}
