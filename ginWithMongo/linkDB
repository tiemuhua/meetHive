package main1

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	//"net/http"
)

/*type User struct {
	Name          string `form:"userName"`
	PassWord      string `form:"passWord"`
	PassWordAgain string `form:"passWordAgain" binding:"eqfield=PassWord"`
	Age           int    `form:"age"`
}*/

func newUser(name string, passWord string) User {
	return User{
		Name:     name,
		PassWord: passWord,
	}
}

var (
	initialRegisteredUsers = []interface{}{
		&User{
			Name:     "tiemuhua",
			PassWord: "20000202tb",
			Age:      200,
		},
		&User{
			Name:     "tom",
			PassWord: "123456",
			Age:      10,
		},
		&User{
			Name:     "Jack",
			PassWord: "qwerty",
			Age:      18,
		},
	}
)

func main1() {
	inputName := "ha"
	inputPassWord := "glgjssy" //苟利国家生死以
	url := "mongodb://localhost:27017"
	dbName := "test"
	colName := "newUsers"
	var err error
	var userCollectionPtr *mongo.Collection

	userCollectionPtr, err = getCollectionPtr(url, dbName, colName)
	if err != nil {
		fmt.Println("get collection error")
	}

	err = userCollectionPtr.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 设置索引
	/*idx := mongo.IndexModel{
		Keys:    bsonx.Doc{{"password", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	var idxRet string
	idxRet, err = userCollectionPtr.Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("usersCollection.Indexes().CreateOne:", idxRet)*/

	var insertOneResultPtr *mongo.InsertOneResult
	insertOneResultPtr, err = userCollectionPtr.InsertOne(context.Background(), initialRegisteredUsers[0])

	if err != nil {
		println("insert one error", insertOneResultPtr.InsertedID)
	}

	err = registerInitialUsers(userCollectionPtr, initialRegisteredUsers)
	if err != nil {
		fmt.Println("register initial users error")
	}

	registerNewUser(userCollectionPtr, inputName, inputPassWord)

	userExist := isUserNameMatchesPassword(userCollectionPtr, inputName, inputPassWord)
	if userExist {
		fmt.Println("user exist")
	} else {
		fmt.Println("user do not exist")
	}
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
	usersFound=usersPtrWithGivenName(collectionPtr,name)
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

	usersCursor.All(context.Background(), &usersFound)

	return  usersFound
}

func registerNewUser(collectionPtr *mongo.Collection, name string, passWord string) error {
	var previousUsersWithThisName []*User
	previousUsersWithThisName=usersPtrWithGivenName(collectionPtr,name)
	if len(previousUsersWithThisName) !=0 {
		return errors.New("this users name has exists, please change your name")
	}
	insertResult, err := collectionPtr.InsertOne(context.Background(), newUser(name, passWord))
	if err != nil {
		fmt.Println("insert error", insertResult.InsertedID)
		return err
	}
	return nil
}
