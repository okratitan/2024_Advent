package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Check for consecutive letters spelling XMAS or SAMX in a row at the given index
func checkHorizontal(row []rune, index int) int {
	// If there aren't enough characters left in the row to spell XMAS or SAMX - we can return early
	if index+3 > len(row)-1 {
		return 0
	}
	// If the character does not start with X or S, then this can not be the word XMAS or SAMX - we can return early
	if row[index] != 'X' && row[index] != 'S' {
		return 0
	}
	// Combine the index character and the next 3 to form a 4 character string and see if it equals XMAS or SAMX
	checkString := string([]rune{row[index], row[index+1], row[index+2], row[index+3]})
	if checkString != "XMAS" && checkString != "SAMX" {
		return 0
	}
	return 1
}

// Check for consecutive letters spelling XMAS OR SAMX in a column at the given index
func checkVertical(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell XMAS or SAMX - we can return early
	if rowIndex+3 > len(wordSearch)-1 {
		return 0
	}
	// If the character does not start with X or S, then this can not be the word XMAS or SAMX - we can return early
	if wordSearch[rowIndex][columnIndex] != 'X' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// Combine the index character and the next 3 to form a 4 character string and see if it equals XMAS or SAMX
	checkString := string([]rune{wordSearch[rowIndex][columnIndex], wordSearch[rowIndex+1][columnIndex],
		wordSearch[rowIndex+2][columnIndex], wordSearch[rowIndex+3][columnIndex]})
	if checkString != "XMAS" && checkString != "SAMX" {
		return 0
	}
	return 1
}

// Check for consecutive letters spelling XMAS or SAMX diagonally down and right in rows and columns at the given index
func checkDiagonalForward(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell XMAS or SAMX - we can return early
	if rowIndex+3 > len(wordSearch)-1 {
		return 0
	}
	// If there aren't enough characters left in the row to spell XMAS or SAMX - we can return early
	if columnIndex+3 > len(wordSearch[rowIndex])-1 {
		return 0
	}
	// If the character does not start with X or S, then this can not be the word XMAS or SAMX - we can return early
	if wordSearch[rowIndex][columnIndex] != 'X' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// Combine the index character and the next 3 to form a 4 character string and see if it equals XMAS or SAMX
	checkString := string([]rune{wordSearch[rowIndex][columnIndex], wordSearch[rowIndex+1][columnIndex+1],
		wordSearch[rowIndex+2][columnIndex+2], wordSearch[rowIndex+3][columnIndex+3]})
	if checkString != "XMAS" && checkString != "SAMX" {
		return 0
	}
	return 1
}

// Check for consecutive letters spelling XMAS or SAMX diagonally down and left in rows and columns at the given index
func checkDiagonalBackward(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell XMAS or SAMX - we can return early
	if rowIndex+3 > len(wordSearch)-1 {
		return 0
	}
	// If there aren't enough characters left in the row to spell XMAS or SAMX - we can return early
	if columnIndex-3 < 0 {
		return 0
	}
	// If the character does not start with X or S, then this can not be the word XMAS or SAMX - we can return early
	if wordSearch[rowIndex][columnIndex] != 'X' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// Combine the index character and the next 3 to form a 4 character string and see if it equals XMAS or SAMX
	checkString := string([]rune{wordSearch[rowIndex][columnIndex], wordSearch[rowIndex+1][columnIndex-1],
		wordSearch[rowIndex+2][columnIndex-2], wordSearch[rowIndex+3][columnIndex-3]})
	if checkString != "XMAS" && checkString != "SAMX" {
		return 0
	}
	return 1
}

// Check for consecutive letters spelling MAS or SAM diagonally down and right in rows and columns at the given index
func checkDiagonalForwardXLeg(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell MAS or SAM - we can return early
	if rowIndex+2 > len(wordSearch)-1 {
		return 0
	}
	// If there aren't enough characters left in the row to spell MAS or SAM - we can return early
	if columnIndex+2 > len(wordSearch[rowIndex])-1 {
		return 0
	}
	// If the character does not start with M or S, then this can not be the word MAS or SAM - we can return early
	if wordSearch[rowIndex][columnIndex] != 'M' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// Combine the index character and the next 2 to form a 3 character string and see if it equals MAS or SAM
	checkString := string([]rune{wordSearch[rowIndex][columnIndex], wordSearch[rowIndex+1][columnIndex+1],
		wordSearch[rowIndex+2][columnIndex+2]})
	if checkString != "MAS" && checkString != "SAM" {
		return 0
	}
	return 1
}

// Check for consecutive letters spelling MAS or SAM diagonally down and left in rows and columns at the given index
func checkDiagonalBackwardXLeg(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell MAS or SAM - we can return early
	if rowIndex+2 > len(wordSearch)-1 {
		return 0
	}
	// If there aren't enough characters left in the row to spell MAS or SAM - we can return early
	if columnIndex-2 < 0 {
		return 0
	}
	// If the character does not start with M or S, then this can not be the word MAS or SAM - we can return early
	if wordSearch[rowIndex][columnIndex] != 'M' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// Combine the index character and the next 2 to form a 3 character string and see if it equals MAS or SAM
	checkString := string([]rune{wordSearch[rowIndex][columnIndex], wordSearch[rowIndex+1][columnIndex-1],
		wordSearch[rowIndex+2][columnIndex-2]})
	if checkString != "MAS" && checkString != "SAM" {
		return 0
	}
	return 1
}

// Check for instances where MAS and/or SAM create an X by crossing diagonally on the A
func checkForXmases(wordSearch [][]rune, rowIndex int, columnIndex int) int {
	// If there aren't enough characters left in the column to spell MAS or SAM - we can return early
	if rowIndex+2 > len(wordSearch)-1 {
		return 0
	}
	// If there aren't enough characters left in the row to spell MAS or SAM twice diagonally - we can return early
	if columnIndex+2 > len(wordSearch[rowIndex])-1 {
		return 0
	}
	// If the character does not start with M or S, then this can not be the word MAS or SAM - we can return early
	if wordSearch[rowIndex][columnIndex] != 'M' && wordSearch[rowIndex][columnIndex] != 'S' {
		return 0
	}
	// If the character does not start with M or S, then this can not be the word MAS or SAM - we can return early
	if wordSearch[rowIndex][columnIndex+2] != 'M' && wordSearch[rowIndex][columnIndex+2] != 'S' {
		return 0
	}
	// If we do not find two diagonals crossing in a 3 column width from the current index, return early
	if checkDiagonalForwardXLeg(wordSearch, rowIndex, columnIndex) == 0 ||
		checkDiagonalBackwardXLeg(wordSearch, rowIndex, columnIndex+2) == 0 {
		return 0
	}
	return 1
}

// Solve the word search for finding all instances of XMAS and SAMX, then all crossing instances of MAS and SAM
func solveWordSearch(wordSearch [][]rune) {
	totalWords := 0
	totalXes := 0
	for row, _ := range wordSearch {
		for column, _ := range wordSearch[row] {
			totalWords += checkHorizontal(wordSearch[row], column)
			totalWords += checkVertical(wordSearch, row, column)
			totalWords += checkDiagonalForward(wordSearch, row, column)
			totalWords += checkDiagonalBackward(wordSearch, row, column)
			totalXes += checkForXmases(wordSearch, row, column)
		}
	}
	fmt.Println(totalWords)
	fmt.Println(totalXes)
}

func main() {
	var wordSearch [][]rune

	// Open the input data
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read the input data in line by line scanning each line into the word search slice as a slice of runes
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordSearch = append(wordSearch, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Solve the word search
	solveWordSearch(wordSearch)
}

/*--- Day 4: Ceres Search ---
"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:


..X...
.SAMX.
.A..A.
XMAS.S
.X....
The actual word search will be full of letters instead. For example:

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:

....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
Take a look at the little Elf's word search. How many times does XMAS appear?

--- Part Two ---
The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

M.S
.A.
M.S
Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?*/
