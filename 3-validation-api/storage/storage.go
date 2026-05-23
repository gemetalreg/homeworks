package storage

import (
	"email/api/dto"
	"encoding/json"
	"fmt"
	"os"
)

func Save(data dto.VerifyDto) {
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

	fmt.Println("email и hash успешно сохранены в data.json!")
}

func Load() dto.VerifyDto {
	var loaded dto.VerifyDto
	fileData, _ := os.ReadFile("data.json")
	json.Unmarshal(fileData, &loaded)
	return loaded
}

func Clear() {
	err := os.Truncate("data.json", 0)
	if err != nil {
		panic(err)
	}
}
