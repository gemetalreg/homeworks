package storage

import (
	"email/api/configs"
	"encoding/json"
	"fmt"
	"os"
)

func Save(data []configs.Config) {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	// 3. Записываем полученный JSON в локальный файл (перезапись, если был)
	err = os.WriteFile("data.json", fileData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}

	fmt.Println("Список успешно сохранен в data.json!")
}

func Load() []configs.Config {
	var loadedList []configs.Config
	fileData, _ := os.ReadFile("data.json")
	json.Unmarshal(fileData, &loadedList)
	return loadedList

}
