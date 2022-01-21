package main

import (
	"log"
	"net"
	"os"
	pb "github.com/AndyMile/articles/app/proto"
	grpcHandler "github.com/AndyMile/articles/app/provider/grpc"
	"github.com/AndyMile/articles/app/repository"
	"github.com/AndyMile/articles/app/service/article"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func DBconnection(user string, password string, host string, dbname string) *gorm.DB {
	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := DBconnection(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()

	articleRepo := repository.NewArticleRepo(db)

	s := grpc.NewServer()
	as := service.NewArticleService(articleRepo)
	h := grpcHandler.NewHandler(as)

	pb.RegisterArticleServer(s, h)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	if err := s.Serve(l); err != nil {
		panic(err)
	}
}