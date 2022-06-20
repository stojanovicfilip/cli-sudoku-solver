package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the sudoku solver!")
	time.Sleep(time.Second * 1)
	fmt.Printf("Type in the numbers of your sudoku board starting from the top row.\n")
	fmt.Printf("After each row click enter to move to the next row.\n")
	fmt.Printf("For unknown fields type in number 0. Use whitespace between fields.\n")

	scanner := bufio.NewScanner(os.Stdin)
	board := [][]int{}

	for {
		if len(board) == 9 {
			solvedBoard := solveSudoku(board)
			fmt.Println("The solution is:")
			for i := 0; i < len(solvedBoard); i++ {
				fmt.Println(solvedBoard[i])
			}
			break
		}
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		srow := strings.Split(text, " ")
		irow := []int{}

		for i := 0; i < len(srow); i++ {
			intr, _ := strconv.Atoi(srow[i])
			irow = append(irow, int(intr))
		}
		board = append(board, irow)

		if len(text) != 0 {
		} else {
			// exit if user entered an empty string
			break
		}
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Please enter rows")
	}
}

func solveSudoku(board [][]int) [][]int {
	solvePartialSudoku(0, 0, board)
	return board
}

func solvePartialSudoku(row int, col int, board [][]int) bool {
	var currentRow = row
	var currentCol = col

	if currentCol == len(board[currentRow]) {
		currentRow += 1
		currentCol = 0
		if currentRow == len(board) {
			return true
		}
	}

	if board[currentRow][currentCol] == 0 {
		return tryDigitsAtPosition(currentRow, currentCol, board)
	}

	return solvePartialSudoku(currentRow, currentCol+1, board)
}

func tryDigitsAtPosition(row int, col int, board [][]int) bool {
	for digit := 1; digit < 10; digit++ {
		if isValidAtPosition(digit, row, col, board) {
			board[row][col] = digit
			if solvePartialSudoku(row, col+1, board) {
				return true
			}
		}
	}

	board[row][col] = 0
	return false
}

func isValidAtPosition(value int, row int, col int, board [][]int) bool {
	rowIsValid := !rowContains(board, row, value)
	columnIsValid := !columnContains(board, col, value)

	if !rowIsValid || !columnIsValid {
		return false
	}

	// Check subgrid constraints
	subgridRowStart := (row / 3) * 3
	subgridColStart := (col / 3) * 3

	for rowIdx := 0; rowIdx < 3; rowIdx++ {
		for colIdx := 0; colIdx < 3; colIdx++ {
			rowToCheck := subgridRowStart + rowIdx
			colToCheck := subgridColStart + colIdx
			existingValue := board[rowToCheck][colToCheck]

			if existingValue == value {
				return false
			}
		}
	}

	return true
}

func rowContains(board [][]int, row int, value int) bool {
	for _, element := range board[row] {
		if value == element {
			return true
		}
	}
	return false
}

func columnContains(board [][]int, col int, value int) bool {
	for _, row := range board {
		if row[col] == value {
			return true
		}
	}
	return false
}
