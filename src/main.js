const ROWS = 50;
const COLS = 50;

let grid = [];

for (let i = 0; i < ROWS; i++) {
  let row = [];
  for (let j = 0; j < COLS; j++) {
    row.push({ x: j, y: i, alive: false });
  }
  grid.push(row);
}

// create a function thats called update state, it only takes a grid
// and returns a new grid
// it takes a grid and returns a new grid

function( grid ) {
  // loop through the grid
  // for each cell, check the neighbors
  // if the cell is alive and has 2 or 3 neighbors, it stays alive
  // if the cell is dead and has 3 neighbors, it comes alive
  // otherwise, it dies
}