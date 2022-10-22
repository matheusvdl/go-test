package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

type  todo struct {
	ID string `json:"id"`
	Item string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learning API", Completed: false},
	{ID: "2", Item: "Create API", Completed: false},
	{ID: "3", Item: "Test API", Completed: false},
};

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}
 
func addTodo(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo);

	context.IndentedJSON(http.StatusCreated, newTodo);
}

func getTodoById(id string)(*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id := context.Param("id");
	todo, err := getTodoById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo); 
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id");
	todo, err := getTodoById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed;
	context.IndentedJSON(http.StatusOK, todo);
}


func main(){
router := gin.Default();
router.GET("/todos", getTodos)
router.GET("/todos/:id", getTodo)
router.POST("/todos", addTodo)
router.PATCH("/todos/:id", toggleTodoStatus)
router.Run("localhost:9090");
}