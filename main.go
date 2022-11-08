package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var line string
	var filename string
	polybiusSquare := generateBigPolybiusSquare()

	for {
		fmt.Println("Введите команду")

		fmt.Scanf("%s\n", &line)

		if line == "encrypt" || line == "decrypt" {
			fmt.Println("Введите имя файла")
			fmt.Scanf("%s\n", &filename)

			chanFromFile := make(chan string)
			go readFromFile(filename, chanFromFile)
			f, err := os.Create("output.txt")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			encrypt := line == "encrypt"
			for fileLine := range chanFromFile {
				for _, letter := range fileLine {
					f.WriteString(fmt.Sprintf("%c", polybiusSquare.processLetter(letter, encrypt)))
				}
				f.WriteString("\n")
			}
			f.Sync()
			fmt.Println("Готово! Результат в файле output.txt")
		} else if line == "exit" {
			break
		} else {
			fmt.Println("Не понимаю эту команду")
		}
	}
}

func printSquare(polybiusSquare *PolybiusSquare) {
	for i := range polybiusSquare.alphabet {
		for j := range polybiusSquare.alphabet[i] {
			if polybiusSquare.alphabet[i][j] == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%c ", polybiusSquare.alphabet[i][j])
			}
		}
		println()
	}
}

func readFromFile(filename string, output chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output <- scanner.Text()
	}
	close(output)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
