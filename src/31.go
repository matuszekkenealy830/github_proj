package main

import (
    "fmt"
    "strings"
)

func main() {
    // Replace the following string with your own project name and version number
    projectName := "myProject"  // Replace this with the actual project name
    versionNumber := "0.1"      // Replace this with the actual project version

    // Generate a unique project ID using UUIDs
    projectID := fmt.Sprintf("%s-%d", projectName, versionNumber)

    // Get the current working directory and path
    dirPath, err := os.Getwd()
    if err != nil {
        fmt.Println("Error retrieving current working directory:", err)
        return
    }

    // Construct the full path to your project's directory
    projectDir := strings.Join([]string{dirPath}, "/") + "/" + projectID

    // Open a file named "README.md" in the project directory and read its content
    readFile, err := os.Open(projectDir + "/README.md")
    if err != nil {
        fmt.Println("Error opening README.md:", err)
        return
    }

    // Use strings.NewReader to read the entire file's content into a string buffer
    reader := bufio.NewReadCloser(readFile, 1024)
    content, _ := ioutil.ReadAll(reader)

    // Format and print the project description using Go's `fmt.Sprintf` method for readability
    fmt.Printf("## %s\n\n", projectName)
    fmt.Println(strings.TrimSpace(strings.Replace(content, "\n", "", -1)))

    // Close the reader to free up resources
    reader.Close()
}
