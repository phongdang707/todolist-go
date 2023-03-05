package todolisttransport

import (
	"ToDoList/modules/todolistbiz"
	"ToDoList/modules/todolistmodel"
	"ToDoList/modules/todoliststorage"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data todolistmodel.TodoListModel

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		db, err := sql.Open("postgres", "postgres://postgres:123123@localhost:5432/postgres?sslmode=disable")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		store := todoliststorage.NewSQLStore(db)
		biz := todolistbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateTaskWithBiz(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)

	}
}
