package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const SaveFolder = "saves/"

func SaveGameWithName(player *Character) error {
	if _, err := os.Stat(SaveFolder); os.IsNotExist(err) {
		err := os.Mkdir(SaveFolder, 0755)
		if err != nil {
			return fmt.Errorf("failed to create save folder: %v", err)
		}
	}

	fmt.Println("Enter a name for your save file: ")
	var saveName string
	fmt.Scanln(&saveName)

	filename := SaveFolder + saveName + ".json"
	data, err := json.MarshalIndent(player, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Game saved successfully!", saveName)
	return nil
}

func ListSaveFiles() ([]string, error) {
	var files []string
	entries, err := ioutil.ReadDir(SaveFolder)
	if err != nil {
		return files, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func LoadGameWithChoice() (*Character, error) {
	saves, err := ListSaveFiles()
	if err != nil || len(saves) == 0 {
		return nil, fmt.Errorf("no saves found")
	}
	fmt.Println("Available save files: ")
	for i, file := range saves {
		fmt.Printf("%d. %s", i+1, file)
	}
	fmt.Print("Choose a save file by inputting the number: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(saves) {
		return nil, fmt.Errorf("invalid choice")
	}

	selectedFile := SaveFolder +
		saves[choice-1]

	data, err :=
		ioutil.ReadFile(selectedFile)
	if err != nil {
		return nil, err
	}
	var player Character
	err = json.Unmarshal(data, &player)
	if err != nil {
		return nil, err
	}
	fmt.Println("Game loaded successfully!")
	return &player, nil
}

func AutoSaveGame(player *Character) error {
	if _, err := os.Stat(SaveFolder); os.IsNotExist(err) {
		err := os.Mkdir(SaveFolder, 0755)
		if err != nil {
			return fmt.Errorf("failed to create save folder: %v", err)
		}
	}
	filename := fmt.Sprintf("%sAutoSave_Level_%d_%s.json", SaveFolder, player.Level, player.Name)

	data, err := json.MarshalIndent(player, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Game auto-saved as %v\n", filename)
	return nil
}
