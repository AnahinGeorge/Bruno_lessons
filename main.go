package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//list := lesson_3.TaskList{Tasks: make(map[uint]lesson_3.Task)}
	//
	//task1 := lesson_3.NewTask("Уборка", 3)
	//task2 := lesson_3.NewTask("Мытьё полов", 5)
	//
	//list.Add(*task1)
	//list.Add(*task2)
	//
	//list.PrintList()

	//http.HandleFunc("/hello", helloHandler)
	//fmt.Println("Сервер запущен на http://localhost:8080")
	//http.ListenAndServe(":8080", nil)

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {

	nub := "Hello World"
	//wr, _ := json.Marshal(nub)
	c.IndentedJSON(http.StatusOK, nub)
}

//func helloHandler(w http.ResponseWriter, r *http.Retques) {
//
//	nub := "Hello World"
//
//	wr, _ := json.Marshal(nub)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(wr)
//
//}
