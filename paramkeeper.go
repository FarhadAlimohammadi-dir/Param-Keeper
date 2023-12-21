package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run script.go <inputfile> <outputfile>")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        os.Exit(1)
    }
    defer file.Close()

    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        os.Exit(1)
    }
    defer outFile.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        url := scanner.Text()
        if strings.Contains(url, "?") {
            _, err := outFile.WriteString(url + "\n")
            if err != nil {
                fmt.Println("Error writing to output file:", err)
                os.Exit(1)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }
}

