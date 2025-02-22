package jsonthings

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
    Name    string
    Tag     string
}

type TokenData struct {
    ValApiToken string
}

func getFilePath(fileName string) string {
	configDir, err := os.UserConfigDir()
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


func GetTokenData() TokenData {
    filePath := getFilePath("token.json")
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        os.Create(filePath)
        os.WriteFile(filePath, []byte("{}"), 0666)
    }
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
    var payload TokenData
    err = json.Unmarshal(content, &payload)
    if err != nil {
        os.WriteFile(filePath, []byte("{}"), 0666)
        log.Fatal("Error when Unmarshal: ", err)
    }
    return payload
}

func GetFileData(filename string) Data {
    filePath := getFilePath(filename)
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

func WriteTokenData(token string) {
    filePath := getFilePath("token.json")
    if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
        os.WriteFile(filePath, []byte("{}"), 0666)
    }

    data := TokenData{
        ValApiToken: token,
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

func WriteFileData(name string, tag string) {
    filePath := getFilePath("data.json")
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

func PromptToken() {
    fmt.Println("Input Henrikdev valorant api token: ")
    var token string
    fmt.Scan(&token)
    WriteTokenData(token)
}
