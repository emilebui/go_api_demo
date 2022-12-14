package load_config

import (
	"fmt"
	"go_api/helpers"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	// CONFIG TO LOG TO FILE OR STDOUT
	Log2File bool `yaml:"LOG_TO_FILE"`

	//GRPC
	GrpcPort string `yaml:"GRPC_PORT"`

	//GRPC-GATEWAY
	GatewayAddr string `yaml:"GATEWAY_ADDR"`

	// Skip files config
	SkipList []string `yaml:"SKIP_FILES"`
}

// LoadConfigInit Initialize config from `config.yaml`
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
