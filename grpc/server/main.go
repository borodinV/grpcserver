package main

import (
	"flag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpc/server/api"
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
	defer db.Close()

	repository := repo.NewRepository(db)
	server := api.NewServer(repository)

	grpcServer := grpc.NewServer()
	proto.RegisterLibraryServer(grpcServer, server)

	log.Printf("Starting grpc listener on port: " + viper.GetString("port"))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
func initConfig() error {
	viper.AddConfigPath(*configuration)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
