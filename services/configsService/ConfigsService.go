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
	port          string
	databasePath  string
	baseFilesBath string
	certificateFilePath string
	keyFilePath string
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
		serviceInstance.port = serviceInstance.data["port"]
		serviceInstance.databasePath = serviceInstance.data["databasePath"]
		serviceInstance.baseFilesBath = serviceInstance.data["baseFilesPath"]
		serviceInstance.certificateFilePath = serviceInstance.data["certificateFilePath"]
		serviceInstance.keyFilePath = serviceInstance.data["keyFilePath"]
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

func (service ConfigsService) GetPort() string {
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

func (service ConfigsService) GetCertificateFilePath() string {
	return service.certificateFilePath
}

func (service ConfigsService) GetKeyFilePath() string {
	return service.keyFilePath
}
