package mon01

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/vinceyoumans/pan02/structs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	jsonfileName = "output.json"
)

func Mon01(mm structs.Mon) string {

	// create the JSON DIR if it does not exist
	step001(mm.DirTorecordJSON)

	log.Println("==========  in Mon01 Library")

	m := make(map[string]int64)
	var fileSize int64

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("----------event:", event)

				transEvent := strings.Contains(event.Name, "~")
				if transEvent {
					log.Println("event file is still in transition state")
					continue
				}

				if event.Has(fsnotify.Create) {
					log.Println("01----create------------------------------")
					log.Println("01 - file Name:", event.Name)
					fileSize = getEventDetails(event.Name)
					m[event.Name] = fileSize
					log.Println(m)
				}

				if event.Has(fsnotify.Write) {
					log.Println("02----Write")
					log.Println("02 - file Name:", event.Name)
					fileSize = getEventDetails(event.Name)
					m[event.Name] = fileSize
					log.Println(m)
				}

				if event.Has(fsnotify.Remove) {
					log.Println("03----Remove")
					log.Println("03 - file Name:", event.Name)
					delete(m, event.Name)
					log.Println(m)
				}

				if event.Has(fsnotify.Rename) {
					// this one is tricky because I don't know the other name unless I look at chmod
					//  so can't remove from the map
					log.Println("04 ----Rename")
					log.Println("-4 - file Name:", event.Name)
					m[event.Name] = fileSize
					log.Println(m)
				}

				if event.Has(fsnotify.Chmod) {
					log.Println("05 ----Chmod")
					log.Println("05 file Name:", event.Name)

					//  FIX this...
					//ename := strings.Replace(event.Name, "~", "", 1)
					//delete(m, ename)
					//log.Println(m)

					log.Println("xxxxxx===================")
					log.Println(m)

				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}

			// NOTE!!!!   I this this could be better.
			// passing m by reference could create a race condition
			// also, making this a go routine will probably corrupt
			// the json file.  On the other hand, opening, editing and closing a file
			// with every event ( as being done here... is also a not efficient.
			// a smarter idea would be to send the map update to a messageQ where another
			// Microservice could process the change to a log.
			step002(mm.DirTorecordJSON, m)
		}
	}()

	err = watcher.Add(mm.DirToMonitor)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})

	//===============
	return fmt.Sprintf("your string is bingo")
}

// ============================================
func step002(fileJSON_Dir string, m map[string]int64) {

	jsonData, err := convertMapToJSON(m)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	filePath := filepath.Join(fileJSON_Dir, jsonfileName)
	err = saveToFile(filePath, jsonData)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}
}

func convertMapToJSON(inputMap map[string]int64) ([]byte, error) {
	jsonData, err := json.MarshalIndent(inputMap, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Function to save data to a file
func saveToFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// ===========================================================
// ===========================================================
func getEventDetails(name string) int64 {

	log.Println("========    getEventDetails")

	fileInfo, err := os.Stat(name)
	if err != nil {
		log.Println("ee  =====   temp file issues")
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	log.Println("---------Size of the file:", fileSize)

	return fileSize
}

// ===========================================================
// ===========================================================
func step001(torecordJSON string) {
	// Call the function to create the subdirectory if it does not exist
	err := createDirectoryIfNotExists(torecordJSON)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// Combine directory path with file name
	filePath := filepath.Join(torecordJSON, jsonfileName)

	// Call the function to create the file if it does not exist
	err = createFileIfNotExists(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("Subdirectory and file created successfully:", filePath)

}

// Create a directory if it does not exist
func createDirectoryIfNotExists(dirPath string) error {

	log.Println("--- dirPath: ", dirPath)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create a file if it does not exist
func createFileIfNotExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}
