import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Board from './components/Board';
import './App.css';

function updateBoard(board) {
  console.log('updateBoard called with board:', board);
  axios.post('http://localhost:8081/update-board', board)
    .then(response => {
      // handle the updated board returned by the backend
      console.log(response.data);
    })
    .catch(error => {
      // handle errors
      console.log(error);
    });
}

function App() {
  const [board, setBoard] = useState([]);
  const [rows, setRows] = useState(100);
  const [cols, setCols] = useState(100);
  const [isRunning, setIsRunning] = useState(false);
  const [hasStarted, setHasStarted] = useState(false);
  const [intervalId, setIntervalId] = useState(null);

  // Initialize board on mount
  useEffect(() => {
    const initialBoard = initializeBoard(false, rows, cols);
    setBoard(initialBoard);
  }, [rows, cols]);

  // Function to initialize board
  const initializeBoard = (value, rows, cols) => {
    const board = [];
    for (let i = 0; i < rows; i++) {
      board.push([]);
      for (let j = 0; j < cols; j++) {
        board[i].push(value);
      }
    }
    return board;
  };

  const [isBoardModified, setIsBoardModified] = useState(false);
  // Function to toggle cell state
  const toggleCellState = (i, j) => {
    const newBoard = [...board];
    newBoard[i][j] = !newBoard[i][j];
    setBoard(newBoard);
    setIsBoardModified(true);
  };

  // Function to start animation
  const startAnimation = () => {
    console.log('Start animation called');
    if (isBoardModified) {
      setIsRunning(true);
      setHasStarted(true);
      const id = setInterval(() => {
        updateBoard(board);
      }, 1000);
      setIntervalId(id);
    } else {
      alert('Please modify the board before starting the game.');
    }
  };

  // Function to stop animation
  const stopAnimation = () => {
    setIsRunning(false);
    clearInterval(intervalId);
  };

  return (
    <div className="App">
      <h1>Conway's Game of Life</h1>
      <div>
        <button onClick={() => setBoard(initializeBoard(0, rows, cols))}>Reset</button>
        {!isRunning ? (
          <button onClick={startAnimation}>Start</button>
        ) : (
          <button onClick={stopAnimation}>Stop</button>
        )}
      </div>
      <Board board={board} toggleCellState={toggleCellState} setHasStarted={setHasStarted} />
    </div>
  );
}

export default App;
