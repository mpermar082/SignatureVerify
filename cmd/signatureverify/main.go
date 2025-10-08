// cmd/signatureverify/main.go
package main

import (
    "flag"
    "log"
    "os"

    "signatureverify/internal/signatureverify"
)

func main() {
    // Define flags with clear descriptions
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Path to the input file")
    output := flag.String("output", "", "Path to the output file")
    flag.Parse()

    // Create a new app instance with verbose logging
    app := signatureverify.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}