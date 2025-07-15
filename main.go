package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatalf("program terminated with error: %w", err)
		}
		input := cleanInput(scanner.Text())
		fmt.Println(input)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
