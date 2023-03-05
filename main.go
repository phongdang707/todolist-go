package main

import (
	"ToDoList/modules/todolisttransport"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type DataOut struct {
	Name        string
	Description string
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Error:", err)
	}

	row, err := db.Query(`select $1, $2 from notes`, "name", "description")

	data := []DataOut{}

	for row.Next() {
		var name string
		var description string

		if err := row.Scan(&name, &description); err != nil {
			panic(err)
		}

		rowData := DataOut{Name: name, Description: description}
		data = append(data, rowData)
	}

	fmt.Println(data)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": data,
		})
	})

	r.POST("/", todolisttransport.CreateTask())

	r.Run(":8080")
}
