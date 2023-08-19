package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/winartodev/go-grpc/config"
	todoHandler "github.com/winartodev/go-grpc/handler"
	todoRepository "github.com/winartodev/go-grpc/repository/mysql"
	"github.com/winartodev/go-grpc/usecase"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var config config.Config
	config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FJakarta",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open(fmt.Sprint(config.Database.Driver), connectionString)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.TodoList.Host, config.TodoList.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	todoRepository := todoRepository.NewTodoRepository(db)

	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	todoHandler.NewTodoHandler(grpcServer, todoUsecase)

	fmt.Println("Server Run at :9000")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	grpcServer.GracefulStop()
	fmt.Printf("\nServer gracefully stopped.")
}
