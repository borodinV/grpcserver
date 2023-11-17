package main

import (
	"context"
	"flag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpc/server/api"
	"grpc/server/migrations"
	"grpc/server/proto"
	"grpc/server/repo"
	"log"
	"net"
)

const (
	configFlagName        = "cfg"
	configFlagDescription = "path to configuration file in work dir"
)

var configuration = flag.String(configFlagName, "configs", configFlagDescription)

func initConfig() error {
	viper.AddConfigPath(*configuration)
	viper.SetConfigName("config2")
	return viper.ReadInConfig()
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("----> unary interceptor", info.FullMethod)
	return handler(ctx, req)
}

func main() {

	flag.Parse()

	if err := initConfig(); err != nil {
		log.Fatalf("Config initialization error: %v", err)
	}

	listener, err := net.Listen("tcp", viper.GetString("host")+":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	err = migrations.InitMigrations("migrations", "up")
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	db, err := repo.InitPostgresDB()
	if err != nil {
		log.Fatalf("DB initialization error: %v", err)

	}
	defer db.Close()

	server := api.NewServer(repo.NewRepository(db))

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor))
	proto.RegisterLibraryServer(grpcServer, server)

	log.Printf("Starting grpc listener on port: " + viper.GetString("port"))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
