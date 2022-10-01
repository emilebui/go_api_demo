package load_config

import (
	"fmt"
	"go_api/helpers"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	// SQL Config
	SqlURI string `yaml:"SQLALCHEMY_DATABASE_URI"`

	//GRPC
	GrpcPort string `yaml:"GRPC_PORT"`

	//GRPC-GATEWAY
	GatewayAddr string `yaml:"GATEWAY_ADDR"`

	// Skip files config
	SkipList []string `yaml:"SKIP_FILES"`
}

func LoadConfigInit() Config {
	filename := helpers.GetDirPath("config.yaml")

	return LoadConfigLogic(filename)
}

func LoadConfigLogic(path string) Config {
	config := Config{}

	yamlFile, err := os.ReadFile(path)
	helpers.ErrorHandler(err)

	err = yaml.Unmarshal(yamlFile, &config)
	helpers.ErrorHandler(err)

	fmt.Println("Finished Loading Config")
	return config
}

// Conf Singleton Config Object
var Conf = LoadConfigInit()
