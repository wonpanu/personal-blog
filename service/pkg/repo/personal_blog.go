package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/wonpanu/personal-blog/service/pkg/entity"
	"github.com/wonpanu/personal-blog/service/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IBlogRepo interface {
	Create(blog entity.Blog) (entity.Blog, error)
	GetAll() ([]entity.Blog, error)
	GetByBlogID(ID string) (entity.Blog, error)
	UpdateByBlogID(ID string, blog entity.Blog) (entity.Blog, error)
	DeleteByBlogID(ID string) error
}

const (
	database   = "personal-blog"
	collection = "wonpanu-personal-blog"
)

type MongoRepo struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (r MongoRepo) Create(blog entity.Blog) (entity.Blog, error) {
	seed := time.Now().Unix()
	blog.ID = util.Hash(fmt.Sprint(seed))
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := r.DB.Collection(collection).InsertOne(ctx, blog)
	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (r MongoRepo) GetAll() ([]entity.Blog, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	cur, err := r.DB.Collection(collection).Find(ctx, bson.M{})
	var blogs []entity.Blog
	if err != nil {
		return blogs, err
	}
	cur.All(ctx, &blogs)
	return blogs, nil
}

func (r MongoRepo) GetByBlogID(ID string) (entity.Blog, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	result := r.DB.Collection(collection).FindOne(ctx, filter)
	var blog entity.Blog
	result.Decode(&blog)
	return blog, nil
}

func (r MongoRepo) UpdateByBlogID(ID string, blog entity.Blog) (entity.Blog, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	blog.ID = ID
	update := bson.M{
		"$set": blog,
	}
	_, err := r.DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (r MongoRepo) DeleteByBlogID(ID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	_, err := r.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func NewBlogRepo(mongo *mongo.Client) MongoRepo {
	return MongoRepo{
		DB: mongo.Database(database),
	}
}
