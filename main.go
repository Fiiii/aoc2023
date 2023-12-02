package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get the folder name from the command line arguments
	folderName := os.Args[1]
	if folderName == "" {
		fmt.Println("please provide a folder name, ie. make newday-1")
		return
	}

	mainFileName := "main.go"
	testFileName := "main_test.go"
	inputFileName := "input.txt"

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
		return
	}

	scratcherFolderPath := fmt.Sprintf("%s/%s", dir, folderName)
	inputFilePath := scratcherFolderPath + "/" + inputFileName

	// Check if the folder already exists
	if _, err := os.Stat(scratcherFolderPath); err == nil {
		fmt.Printf("folder '%s' already exists.\n", folderName)
	} else if os.IsNotExist(err) {
		fmt.Printf("creating folder '%s'.\n", folderName)

		// Folder does not exist, create it
		if err := os.Mkdir(scratcherFolderPath, os.ModePerm); err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}

		// Create the main.go file and write the code
		filePath := scratcherFolderPath + "/" + mainFileName
		if err := os.WriteFile(filePath, []byte(mainCode(folderName, inputFilePath)), os.ModePerm); err != nil {
			fmt.Printf("error creating %s: %v\n", mainFileName, err)
			return
		}

		// Create the main_test.go file and write the code
		tc := getTestCode(folderName)
		filePath = scratcherFolderPath + "/" + testFileName
		if err := os.WriteFile(filePath, []byte(tc), os.ModePerm); err != nil {
			fmt.Printf("error creating %s: %v\n", testFileName, err)
			return
		}

		// Create the txt file for aoc testing data
		filePath = scratcherFolderPath + "/" + inputFileName
		if err := os.WriteFile(filePath, nil, os.ModePerm); err != nil {
			fmt.Printf("error creating %s: %v\n", inputFileName, err)
			return
		}

		// Open the file in Goland if needed. It needs to have executable permissions & command
		//openFileInGoland(filePath)
	} else {
		fmt.Println("Error checking folder:", err)
		return
	}
}

func getTestCode(day string) string {
	return fmt.Sprintf(`package main

import (
	"testing"
)

func Test_Main_Day_%s(t *testing.T) {
	t.Parallel()
}`, day)
}

func mainCode(day, inputFile string) string {
	code := fmt.Sprintf(`package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: %s")

	file, err := os.Open("%s")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	end := time.Now()
	fmt.Println("Execution time: ", end.Sub(start))
}
`, day, inputFile)

	return code
}

// TODO: if someone need to open the file in Goland, it needs to have executable permissions & command.
func openFileInGoland(path string) {
	shell := "zsh"
	command := fmt.Sprintf("goland %s", path)
	// Create a command object
	cmd := exec.Command(shell, "-c", command)

	fmt.Println("Running command:", command)
	// Run the command and capture its output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}
