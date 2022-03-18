package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wonpanu/personal-blog/service/pkg/handler"
	"github.com/wonpanu/personal-blog/service/pkg/repo"
	"github.com/wonpanu/personal-blog/service/pkg/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectMongoDB() (*mongo.Client, error) {
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPortRaw := os.Getenv("MONGO_PORT")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")
	mongoPort, _ := strconv.Atoi(mongoPortRaw)

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongoUser, mongoPass, mongoHost, mongoPort)
	fmt.Println("mongoConnectionString:", mongoURI)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected 🎉")
	return client, err
}

func main() {
	fmt.Println("🎉 Welcome to Personal Blog service 🙂")

	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	client, err := connectMongoDB()
	defer client.Disconnect(context.TODO())
	if err != nil {
		log.Println("Failed to disconnect MongoDB!")
	}

	blogRepo := repo.NewBlogRepo(client)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(handler.Response{
			Status: "ok",
			Data:   "🎉 Welcome to Personal Blog service 🙂",
		})
	})
	app.Post("/personal-blog", blogHandler.Create)
	app.Get("/personal-blog", blogHandler.GetAll)
	app.Get("/personal-blog/:id", blogHandler.GetByBlogID)
	app.Post("/personal-blog/:id", blogHandler.UpdateByBlogID)
	app.Delete("/personal-blog/:id", blogHandler.DeleteByBlogID)

	portRaw := os.Getenv("PORT")
	port, err := strconv.Atoi(portRaw)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
