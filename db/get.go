package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/phihdn/nc_student/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAllStudent return all students in DB
func GetAllStudent() (*[]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var students []models.Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func SearchStudentSimple(req *models.StudentSearchRequest) (*[]models.Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var filter bson.M
	bs, err := json.Marshal(req)
	err = json.Unmarshal(bs, &filter)
	if err != nil {
		log.Printf("marshal error: %v", err)
	}

	fmt.Println("Filter", filter)

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []models.Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

// GetAllStudentGroupByLastName return all students in DB and group by last name
func GetAllStudentGroupByLastName() (*[]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$last_name"},
			{"students", bson.D{
				{"$push", "$$ROOT"},
			}},
		}},
	}
	cur, err := Client.Database(DbName).Collection(ColName).Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var students []bson.M
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func GetStudentById(id int) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}

	var student models.Student
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, filter).Decode(&student)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	return &student, nil
}