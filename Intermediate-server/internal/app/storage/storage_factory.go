package storage

import "fmt"

func GetStorage(typeOfStorage string) (IStorage, error) {
	switch typeOfStorage {
	default:
		return nil, fmt.Errorf("Storage type is not defined")
	case StorageTasks:
		return getSingleTasksInstance(), nil
	case StorageTelegramApiConfigs:
		return getSingleTelegramApiConfigsInstance(), nil
	}
}
