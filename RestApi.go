package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type google struct{
	ID   string `json: "id"`
	Item string `json: "item"`
	Completed bool `json: "completed"`
}
var googles = []google{
	{ID: "1",Item:"clean Room",Completed: false},
	{ID: "2",Item:"Read Book",Completed: false},
	{ID: "3",Item:"Record video",Completed: false},
}
func getgoogles(context *gin.Context){
		context.IndentedJSON(http.StatusOK,googles)
}
func addgoogles(context *gin.Context){
	var newgoogle google
	if err:=context.ShouldBindJSON(&newgoogle); err!=nil{
		return
	}
	googles = append(googles,newgoogle )
	context.IndentedJSON(http.StatusCreated,newgoogle)

}
func getById(id string)(*google,error){
	for i,t:= range googles{
		if t.ID==id{
			return &googles[i],nil
		}
	}
	return nil, errors.New("googles not found")
}
func getgoogle(context *gin.Context){
	id := context.Param("id")
	google,err:= getById(id)
	if err!=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found google"})
		return
	}
	context.IndentedJSON(http.StatusOK,google)
}
func togglegoogle(context *gin.Context){
	id := context.Param("id")
	google,err:= getById(id)
	if err!=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found google"})
		return
	}
	google.Completed=!google.Completed
	context.IndentedJSON(http.StatusOK,google)
}
// func deleteById(context *gin.Context){
// 	id := context.Param("id")
	
	
	
// }
func main(){
	router:= gin.Default()
	router.GET("/googles",getgoogles)
	router.GET("/googles/:id",getgoogle)
	router.PATCH("/googles/:id",togglegoogle)
	router.POST("/googles",addgoogles)
	fmt.Print("strating server")
	router.Run("localhost:9090")
}