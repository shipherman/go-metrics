package filestore

import (
    "os"
    "encoding/json"
    "github.com/shipherman/go-metrics/internal/storage"

)

func cleanFile(filename string) error {
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer f.Close()


    err = f.Truncate(0)
    if err != nil {
        return err
            }
    return nil
}

func WriteDataToFile(filename string, store storage.MemStorage) error {
    //clear file
    err := cleanFile(filename)
    if err != nil {
        return err
    }

    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer f.Close()


    data, err := json.MarshalIndent(store, "", "  ")
    if err != nil {
        return err
    }

    _, err = f.Write(data)
    if err != nil {
        return err
    }
    return nil
}
