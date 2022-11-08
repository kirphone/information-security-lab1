package main

import "math/rand"

type PolybiusSquare struct {
	letterToIndex map[rune]int
	alphabet      [][]rune
	width         int
}

func generateBigPolybiusSquare() *PolybiusSquare {
	width := 100
	size := width * width
	alphabetRaw := make([]rune, size)
	for i := 32; i < size; i++ {
		alphabetRaw[i] = rune(i)
	}

	rand.Shuffle(len(alphabetRaw), func(i, j int) {
		alphabetRaw[i], alphabetRaw[j] = alphabetRaw[j], alphabetRaw[i]
	})
	letterToIndex := make(map[rune]int)
	for i, letter := range alphabetRaw {
		letterToIndex[letter] = i
	}

	alphabet := make([][]rune, width)
	for i := range alphabet {
		alphabet[i] = make([]rune, width)
		for j := range alphabet[i] {
			alphabet[i][j] = alphabetRaw[i*width+j]
		}
	}
	return &PolybiusSquare{letterToIndex, alphabet, width}
}

func (square *PolybiusSquare) processLetter(letter rune, encrypt bool) rune {
	var delta int
	if encrypt {
		delta = 1
	} else {
		delta = -1
	}
	index := square.letterToIndex[letter]
	newRawIndex := remainder(index/square.width+delta, square.width)
	for square.alphabet[newRawIndex][index%square.width] == 0 {
		newRawIndex = remainder(newRawIndex+delta, square.width)
	}
	return square.alphabet[newRawIndex][index%square.width]
}

func remainder(dividend int, divisor int) int {
	if dividend < 0 {
		return divisor - 1
	} else {
		return dividend % divisor
	}
}
