import React from 'react';

function Board(props) {
  const { board, toggleCellState } = props;

  return (
    <div className="board">
      {board.map((row, i) => (
        <div key={i} className="row">
          {row.map((cell, j) => (
            <div
              key={`${i}-${j}`}
              className={`cell ${cell ? 'alive' : 'dead'}`}
              onClick={() => toggleCellState(i, j)}
            ></div>
          ))}
        </div>
      ))}
    </div>
  );
}

export default Board;