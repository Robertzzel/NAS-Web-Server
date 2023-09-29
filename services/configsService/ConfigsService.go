package configsService

import (
	"encoding/json"
	"os"
	"strconv"
)

const (
	configFile = "./configs.json"
)

var (
	serviceInstance *ConfigsService = nil
)

type ConfigsService struct {
	data          map[string]string
	host          string
	port          int64
	databasePath  string
	baseFilesBath string
	memoryPerUser int64
}

func NewConfigsService() (*ConfigsService, error) {
	if serviceInstance == nil {
		serviceInstance = &ConfigsService{data: make(map[string]string)}
		fileData, err := os.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(fileData, &serviceInstance.data); err != nil {
			return nil, err
		}

		serviceInstance.host = serviceInstance.data["host"]
		serviceInstance.port, err = strconv.ParseInt(serviceInstance.data["port"], 10, 64)
		if err != nil {
			return nil, err
		}
		serviceInstance.databasePath = serviceInstance.data["databasePath"]
		serviceInstance.baseFilesBath = serviceInstance.data["baseFilesPath"]
		serviceInstance.memoryPerUser, err = strconv.ParseInt(serviceInstance.data["memoryPerUser"], 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return serviceInstance, nil
}

func (service ConfigsService) GetHost() string {
	return service.host
}

func (service ConfigsService) GetPort() int64 {
	return service.port
}

func (service ConfigsService) GetDatabasePath() string {
	return service.databasePath
}

func (service ConfigsService) GetBaseFilesPath() string {
	return service.baseFilesBath
}

func (service ConfigsService) GetMemoryPerUser() int64 {
	return service.memoryPerUser
}
