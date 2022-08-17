package belajar_golang_database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

/* SQL command
CREATE TABLE `customer_advance` (
	`id` varchar(100) NOT NULL,
	`name` varchar(100) NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

alter table customer_advance
	add column email varchar(100),
	add column balance integer default 0,
	add column rating DOUBLE default 0.0,
	add column created_at TIMESTAMP default CURRENT_TIMESTAMP,
	add column birth_date DATE,
	add column married BOOLEAN default false;
*/

func ColumnTypeGetConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

func TestColumnType(t *testing.T) {
	db := ColumnTypeGetConnection()
	defer db.Close()

	ctx := context.Background()
	var sqlQuery string
	sqlQuery = `delete from customer_advance where id in ('noval', 'casca')`
	if _, err := db.ExecContext(ctx, sqlQuery); err != nil {
		panic(err)
	}
	sqlQuery = `insert into customer_advance (id, name, email, balance, rating, birth_date, married)
		values(
			'noval',
			'noval',
			'noval@gmail.com',
			150000,
			5.1,
			'1993-11-20',
			true
		),
		(
			'casca',
			'casca',
			null,
			160000,
			5.0,
			null,
			true
		)
	`
	if _, err := db.ExecContext(ctx, sqlQuery); err != nil {
		panic(err)
	}

	sqlQuery = "select id, name, email, balance, rating, birth_date, married, created_at from customer_advance"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var married bool
		var createdAt time.Time
		if err := rows.Scan(
			&id,
			&name,
			&email,
			&balance,
			&rating,
			&birthDate,
			&married,
			&createdAt,
		); err != nil {
			fmt.Println("Ada error: ", err.Error())
			continue
		}

		fmt.Println("==========")
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("Created at: ", createdAt)
		fmt.Println("")
	}

}

func ColumnTypeParamConnection(driver string, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}

func TestColumnTypeWithStruct(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one?parseTime=true"
	db := ColumnTypeParamConnection(driver, source)
	defer db.Close()

	ctx := context.Background()
	query := `select id, name, email, balance, rating, birth_date, married, created_at from customer_advance order by id asc`
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	type Customer struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     *string    `json:"email"`
		Balance   *int64     `json:"balance"`
		Rating    *float64   `json:"rating"`
		BirthDate *time.Time `json:"birth_date"`
		Married   *bool      `json:"married"`
		CreatedAt time.Time  `json:"created_at"`
	}
	var customers []Customer
	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance sql.NullInt64
		var rating sql.NullFloat64
		var birthDate sql.NullTime
		var married sql.NullBool
		var createdAt sql.NullTime

		if err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt); err != nil {
			fmt.Println("Ada error: ", err.Error())
			continue
		}
		var customer Customer
		customer.ID = id
		customer.Name = name
		if email.Valid {
			customer.Email = &email.String
		}
		if balance.Valid {
			customer.Balance = &balance.Int64
		}
		if rating.Valid {
			customer.Rating = &rating.Float64
		}
		if birthDate.Valid {
			customer.BirthDate = &birthDate.Time
		}
		if married.Valid {
			customer.Married = &married.Bool
		}
		if createdAt.Valid {
			customer.CreatedAt = createdAt.Time
		}
		customers = append(customers, customer)
	}

	jsonCustomers, _ := json.Marshal(customers)
	fmt.Println("JSON: ", string(jsonCustomers))
}
