package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "sort"
)

func main() {
    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        fmt.Println("Please provide a folder path or file paths.")
        os.Exit(1)
    }

    var files []string
    outputFilename := ""

    fileInfo, err := os.Stat(args[0])
    if err != nil {
        fmt.Println("Error accessing path:", err)
        os.Exit(1)
    }

    if fileInfo.IsDir() {
        // It's a directory; search for .cnc files
        dirname := args[0]
        outputFilename = filepath.Base(dirname) + ".cnc"

        files, err = filepath.Glob(filepath.Join(dirname, "*.cnc"))
        if err != nil {
            fmt.Println("Error reading directory:", err)
            os.Exit(1)
        }
        if len(files) == 0 {
            fmt.Println("No .cnc files found in directory.")
            os.Exit(1)
        }
        sort.Strings(files)
    } else {
        // It's assumed to be multiple files directly inputted
        files = args
		sort.Strings(files) // Sort files by name
        outputPath := filepath.Dir(files[0])
        outputFilename = filepath.Base(outputPath) + ".cnc"
    }

    // Concatenate files
    outputPath := outputFilename
    outputFile, err := os.Create(outputPath)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        os.Exit(1)
    }
    defer outputFile.Close()

    for _, file := range files {
        err := appendFileContents(file, outputFile)
        if err != nil {
            fmt.Printf("Error writing contents of %s: %s\n", file, err)
            os.Exit(1)
        }
    }

    fmt.Println("Concatenation complete. Output file created:", outputPath)
}

func appendFileContents(filePath string, outputFile *os.File) error {
    inputFile, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    _, err = io.Copy(outputFile, inputFile)
    if err != nil {
        return err
    }
    return nil
}