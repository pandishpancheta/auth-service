package main

import (
	"auth-service/pkg/config"
	"auth-service/pkg/db"
	"auth-service/pkg/pb"
	"auth-service/pkg/service"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	log.Println("DB_HOST: ", cfg.DB_HOST)
	log.Println("DB_PORT: ", cfg.DB_PORT)
	log.Println("DB_USER: ", cfg.DB_USER)
	log.Println("DB_PASS: ", cfg.DB_PASS)
	log.Println("DB_NAME: ", cfg.DB_NAME)

	pgdb := db.Init(cfg)
	db.InitTables(pgdb)

	lis, err := net.Listen("tcp", "localhost:"+cfg.TCP_PORT)
	if err != nil {
		panic(err)
	}

	as := service.NewAuthService(pgdb, cfg)
	us := service.NewUserService(pgdb, cfg)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, as)
	pb.RegisterUserServiceServer(grpcServer, us)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
