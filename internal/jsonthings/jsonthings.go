package jsonthings

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Data struct {
    Name    string
    Tag     string
}

func GetFileData() Data {
    if _, err := os.Stat("data.json"); errors.Is(err, os.ErrNotExist) {
        os.Create("data.json")
        os.WriteFile("data.json", []byte("{}"), 0666)
    }
    content, err := os.ReadFile("data.json")
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
    if _, err := os.Stat("data.json"); errors.Is(err, os.ErrNotExist) {
        os.WriteFile("data.json", []byte("{}"), 0666)
    }
    data := Data{
        Name: name,
        Tag: tag,
    }

    byteValue, err := json.Marshal(data)
    if err != nil {
        log.Fatal("Error Marshal() data: ", err)
    }

    err = os.WriteFile("data.json", byteValue, 0644)
    if err != nil {
        log.Fatal("Error writing data to data.json: ", err)
    }
}
