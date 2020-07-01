package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

type User struct {
	Name          string `form:"userName"`
	PassWord      string `form:"passWord"`
	PassWordAgain string `form:"passWordAgain" binding:"eqfield=PassWord"`
	Age           int    `form:"age"`
}

func newUser(name string, passWord string) User {
	return User{
		Name:     name,
		PassWord: passWord,
	}
}

func main() {

	url := "mongodb://localhost:27017"
	dbName := "test"
	colName := "newUsers"
	var err error
	var userCollectionPtr *mongo.Collection

	userCollectionPtr, err = getCollectionPtr(url, dbName, colName)

	engine := gin.Default()

	engine.LoadHTMLGlob("./login.html")

	engine.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"promptInformation": "please input your personal message here",
		})
	})

	engine.POST("/upload", func(context *gin.Context) {
		var user User
		var outputMessage string
		err = context.ShouldBind(&user)
		//loginOrRegister:=context.Query("loginOrRegister")
		loginOrRegister := context.PostForm("loginOrRegister")
		if err != nil {
			outputMessage = err.Error()
		} else {
			if loginOrRegister == "login" {
				if !userExists(userCollectionPtr,user.Name) {
					outputMessage="user does not exist"
				} else if isUserNameMatchesPassword(userCollectionPtr, user.Name, user.PassWord) {
					outputMessage = "login successfully"
				} else {
					outputMessage = "user name does not match the password"
				}
			} else {
				err = registerNewUser(userCollectionPtr, user.Name, user.PassWord)
				if err != nil {
					outputMessage = user.Name + " has registered successfully"
				} else {
					outputMessage = err.Error()
				}
			}
		}
		context.HTML(http.StatusOK, "login.html", gin.H{
			"promptInformation": outputMessage,
		})
	})

	engine.Run()
}

func getCollectionPtr(url string, dbName string, colName string) (*mongo.Collection, error) {
	opts := options.Client().ApplyURI(url)

	// 连接数据库
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Println("service can not use")
		log.Fatal(err)
	}

	// 获取数据库和集合
	return client.Database(dbName).Collection(colName), err
}

func registerInitialUsers(collectionPtr *mongo.Collection, initialRegisteredUsers []interface{}) error {
	insertManyResult, err := collectionPtr.InsertMany(context.Background(), initialRegisteredUsers[0:])
	if err != nil {
		println("insert error:", insertManyResult.InsertedIDs)
	}
	return err
}

func isUserNameMatchesPassword(collectionPtr *mongo.Collection, name string, passWord string) bool {
	var usersFound []*User
	usersFound = usersPtrWithGivenName(collectionPtr, name)
	fmt.Println("usersFound", len(usersFound))
	for i := 0; i < len(usersFound); i++ {
		if usersFound[i].PassWord == passWord {
			return true
		}
	}
	return false
}

func usersPtrWithGivenName(collectionPtr *mongo.Collection, name string) []*User {
	var usersFound []*User
	usersCursor, err := collectionPtr.Find(context.Background(), bson.M{"name": name})
	if err != nil {
		fmt.Println("collection find error")
	}

	err = usersCursor.All(context.Background(), &usersFound)

	return usersFound
}

func userExists(collectionPtr *mongo.Collection, name string) bool {
	users :=usersPtrWithGivenName(collectionPtr,name)
	if len(users)==0{
		return false
	}
	return  true
}

func registerNewUser(collectionPtr *mongo.Collection, name string, passWord string) error {
	var previousUsersWithThisName []*User
	previousUsersWithThisName = usersPtrWithGivenName(collectionPtr, name)
	if len(previousUsersWithThisName) != 0 {
		return errors.New("this users name has exists, please change your name")
	}
	insertResult, err := collectionPtr.InsertOne(context.Background(), newUser(name, passWord))
	if err != nil {
		fmt.Println("insert error", insertResult.InsertedID)
		return err
	}
	return nil
}
