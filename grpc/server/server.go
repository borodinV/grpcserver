package main

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpc/server/proto"
	"grpc/server/repo"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	if err := initConfig(); err != nil {
		log.Fatalf("Config initialization error: %v", err)
	}

	db, err := repo.NewPostgresDB(repo.Config{
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("DB initialization error: %v", err)

	}

	server := NewServer(db)

	grpcServer := grpc.NewServer()
	proto.RegisterLibraryServer(grpcServer, server)

	log.Printf("Starting grpc listener on port: " + viper.GetString("port"))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
