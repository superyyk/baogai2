package test1

import (
	"context"
	"log"

	"github.com/superyyk/yishougai/db"
	"github.com/superyyk/yishougai/tool"
	"github.com/superyyk/yishougai/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = db.MsgCollection

type TestData struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
	Age  string `bson:"age"`
}

func MongoAddOne(c *gin.Context) {
	id := utils.RandInt(10)
	name := c.Query("name")
	age := c.Query("age")
	t := &TestData{
		Id:   tool.Int_string(id),
		Name: name + tool.Int_string(id),
		Age:  age,
	}
	objId, err := collection.InsertOne(context.TODO(), &t)
	if err != nil {
		tool.Fail(c, "插入数据失败", err)
		return
	} else {
		tool.Success(c, "success", objId)
	}

}

// 2、新增一条数据
func AddOne(t *TestData) {
	objId, err := collection.InsertOne(context.TODO(), &t)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("录入数据成功，objId:", objId)
}

// 3、删除一条数据
func Del(m bson.M) {
	deleteResult, err := collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
}

// 4、编辑一条数据
func EditOne(c *gin.Context) {
	//var t *TestData
	//var m bson.M //
	filter := bson.D{{
		"id", c.Query("id"),
	}}

	//update := bson.M{"$set": t}
	update := bson.D{
		{
			"$inc", bson.D{
				{
					"name", c.Query("name"),
				},
				{
					"age", c.Query("age"),
				},
			},
		},
	}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		tool.Fail(c, "修改失败", err)
		return
	}
	tool.Success(c, "修改成功", updateResult)
}

// 5、更新数据 - 存在更新，不存在就新增
func Update(t *TestData, m bson.M) {
	update := bson.M{"$set": t}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

// 6、模糊查询
func Sectle(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var t TestData
		if err = cur.Decode(&t); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{xx}: ", t)
	}
	_ = cur.Close(context.Background())
}

// 7、准确搜索一条数据
func GetOne(m bson.M) {
	var one TestData
	err := collection.FindOne(context.Background(), m).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.FindOne: ", one)
}

// 8、获取多条数据
func GetList(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*TestData
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	for _, one := range all {
		log.Println("Id:", one.Id, " - name:", one.Name, " - age:", one.Age)
	}
}

// 9、统计collection的数据总数
func Count() {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("collection.CountDocuments:", count)
}
