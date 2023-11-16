package main

import (
	"customer/internal/config"
	"customer/internal/db/psql"
	"customer/internal/users"
	"customer/pkg/customer/pkg/prots"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	if err := run(":8001"); err != nil {
		log.Fatal(err)
	}
}

func run(serv string) error {
	//INIT CONF DB
	configDB, err := config.NewConfig("./config/database.yaml")
	if err != nil {
		return err
	}
	//INIT DB CONN
	s := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		configDB.Username, configDB.Password, configDB.Database, configDB.Host, configDB.Port,
	)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return err
	}
	repo := psql.NewDB(db)

	//GRPC USERS
	list, err := net.Listen("tcp", serv)
	if err != nil {
		return err
	}
	serverRegistration := grpc.NewServer()

	newUs := users.NewUserS(repo)

	prots.RegisterUserServiceServer(serverRegistration, newUs)
	if err := serverRegistration.Serve(list); err != nil {
		return err
	}
	return nil
}
