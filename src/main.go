package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

type Vector []bool
type Board []Vector

func updateState(board Board) Board {
	newBoard := createBoard(len(board[0]), len(board))
	for i := range board {
		for j := range board[i] {
			newBoard[i][j] = updateCell(board, i, j)
		}
	}
	return newBoard
}

func createBoard(height, width int) Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make(Vector, width)
	}
	return board
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
				if board[(x+i+len(board))%len(board)][(y+j+len(board[0]))%len(board[0])] {
					liveNeighbors++
				}
			}
		}
	}
	return liveNeighbors
}

func updateBoardHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to update board")
	log.Println(r.Body)

	// Ensure that the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// Ensure that the content type is correct
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	// Allow cross-origin resource sharing (CORS)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Decode the JSON data sent from the client
	decoder := json.NewDecoder(r.Body)
	var board Board
	err := decoder.Decode(&board)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the board
	board = updateState(board)

	// Encode the updated board as JSON and send it back to the client
	encoder := json.NewEncoder(w)
	err = encoder.Encode(board)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Successfully updated board")
}

func main() {
    
    // Register the updateBoardHandler function as the handler for the "/update-board" endpoint
    http.HandleFunc("/update-board", updateBoardHandler)
    
    // Allow cross-origin resource sharing (CORS) for all endpoints
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    origins := handlers.AllowedOrigins([]string{"*"})
    
    // Start the HTTP server on port 8081 with CORS enabled
    log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(nil)))
}