package jsonthings

import (
	"encoding/json"
    "path/filepath"
	"errors"
	"log"
	"os"
    _ "embed"
)

type Data struct {
    Name    string
    Tag     string
}

func getFilePath() string {
    fileName := "data.json"
	configDir, err := os.UserConfigDir() // Gets ~/.config on Linux/macOS, %APPDATA% on Windows
	if err != nil {
		panic("Could not get user config directory")
	}

	// Ensure ValTracker directory exists
	appDir := filepath.Join(configDir, "ValTracker")
	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		panic("Could not create config directory")
	}

	return filepath.Join(appDir, fileName)
}

func GetFileData() Data {
    filePath := getFilePath()
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        os.Create(filePath)
        os.WriteFile(filePath, []byte("{}"), 0666)
    }
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
    var payload Data
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error when Unmarshal: ", err)
    }
    return payload
}

func WriteFileData(name string, tag string) {
    filePath := getFilePath()
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        os.WriteFile(filePath, []byte("{}"), 0666)
    }
    data := Data{
        Name: name,
        Tag: tag,
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
