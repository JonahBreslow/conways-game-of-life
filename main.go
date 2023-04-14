package main

// import the fmt and math/rand packages
import (
	"fmt"
	"math/rand"
)

type Vector []bool
type Board []Vector 

// create an arbitrarily large 2d board for the game
func createBoard (height, width int) Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make(Vector, width)
	}
	return board
}

func initializeBoard(board Board, probability float32) Board {
	for i := range board {
		for j := range board[i] {
			if rand.Float32() < probability {
				board[i][j] = true
			}
		}
	}
	return board
}

func updateState(board Board) Board { 
	newBoard := createBoard(len(board[0]), len(board))
	for i := range board {
		for j := range board[i] {
			newBoard[i][j] = updateCell(board, i, j)
		}
	}
	return newBoard
}

func updateCell(board Board, x, y int) bool {
	if isBoardExtinct(board) {
		fmt.Println("Game is Over")	
	}
	liveNeighbors := getLiveNeighbors(board, x, y)
	if board[x][y] {
		if liveNeighbors == 2 || liveNeighbors == 3 {
			return true
		}
		return false
	}
	if liveNeighbors == 3 {
		return true
	}
	return false
}

func isBoardExtinct(board Board) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				return false
			}
		}
	}
	return true
}

func getLiveNeighbors(board Board, x, y int) int {
	liveNeighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				// The modulus operator is used to wrap the board as a torus
				if board[(x+i+len(board))%len(board)][(y+j+len(board[0]))%len(board[0])] {
					liveNeighbors++
				}
			}
		}
	}
	return liveNeighbors
}

func printBoard(board Board) {
	// print the boards as 1s and 0s. make sure to add an extra line after each board
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
    board := createBoard(10, 10) 
	board = initializeBoard(board, 0.1)
	if isBoardExtinct(board) {
		fmt.Println("Game is Over")
		return
	}
	printBoard(board)
	for i := 0; i < 100; i++ {
		board = updateState(board)
		printBoard(board)
		if isBoardExtinct(board) {
			fmt.Println("Game is Over")
			return
		}
	}	
}