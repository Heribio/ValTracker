package jsonthings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath(fileName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Could not get user config directory")
	}
	appDir := filepath.Join(configDir, "ValTracker")
	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		panic("Could not create config directory")
	}
    fullPath := filepath.Join(appDir, fileName)
    if _, err := os.Stat(fullPath); err == nil {
            return fullPath, nil
    }

    os.WriteFile(fullPath, []byte("{}"), 0666)

    return "", err
}

func ReadData(filename string, value any) error {
    path, err := getFilePath(filename)
    if err != nil {
        return err
    }
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, value)
}

func WriteData(filename string, value any) error {
    path, err := getFilePath(filename)
    if err != nil {
        fmt.Println("Directory does not exist")
    }
    data, err := json.MarshalIndent(value, "", "  ")
    if err != nil {
        return err
    }
     return os.WriteFile(path, data, 0644)
}

func PromptToken() {
    fmt.Println("Input Henrikdev valorant api token: ")
    var token string
    fmt.Scan(&token)
    tokenData := TokenData{
        ValApiToken: token,
    }
    WriteData("token.json", tokenData)
}
