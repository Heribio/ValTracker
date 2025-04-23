package jsonthings

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func WriteFavoriteData(p WriteFavoriteParams) {
	data := GetFavoriteData()
    data.Favorites = append(data.Favorites, p.Player)
    if len(data.Favorites) >= 5 {
        data.Favorites= data.Favorites[1:]
    }

    filePath, err := getFilePath("favorite.json")
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        os.WriteFile(filePath, []byte("{}"), 0666)
    }

    byteValue, err := json.Marshal(data)
    if err != nil {
        log.Fatal("Error Marshal() data: ", err)
    }

    err = os.WriteFile(filePath, byteValue, 0644)
    if err != nil {
        log.Fatal("Error writing data to data.json: ", err)
    }
}

func GetFavoriteData() FavoriteData {
    filePath, err := getFilePath("favorite.json")
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(filePath, []byte("{}"), 0666)
	}
    
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
    var payload FavoriteData
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error when Unmarshal: ", err)
    }
    return payload
}
