const fs = require("fs");
const f = "day14/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");

// Part 1
let start = performance.now();

const drawPath = (input) => {
  let screen = createScreen(input);
  input.forEach((coords) => {
    coords = coords.split(" -> ");
    //console.log(coords);
    for (k = 0; k < coords.length; k++) {
      coords[k] = coords[k].split(",").map(Number);
    }

    for (k = 1; k < coords.length; k++) {
      //console.log(coords[k - 1], coords[k]);
      //printScreen(screen);
      if (coords[k][0] == coords[k - 1][0]) {
        //draw vert
        a = coords[k - 1][1];
        b = coords[k][1];
        if (coords[k - 1][1] > coords[k][1]) {
          a = coords[k][1];
          b = coords[k - 1][1];
        }
        for (pix = a; pix <= b; pix++) {
          screen[pix][coords[k][0]] = "#";
          //console.log("1", pix, coords[k][1]);
          //printScreen(screen);
        }
      }
      if (coords[k][1] == coords[k - 1][1]) {
        //draw hor
        a = coords[k - 1][0];
        b = coords[k][0];
        if (coords[k - 1][0] > coords[k][0]) {
          a = coords[k][0];
          b = coords[k - 1][0];
        }
        for (pix = a; pix <= b; pix++) {
          screen[coords[k][1]][pix] = "#";
          //console.log("2", pix, coords[k][0]);
          //printScreen(screen);
        }
      }
    }
  });
  // CAREFUL TO MODIFY FOR REAL INPUT
  screen[0][500] = "+";
  return screen;
};

const sizeScreen = (input) => {
  let n = 0;
  let p = 0;
  input.forEach((line) => {
    line = line.split(" -> ");
    line.forEach((coord) => {
      coord = coord.split(",");
      if (coord[0] > p) {
        p = Number(coord[0]);
      }
      if (coord[1] > n) {
        n = Number(coord[1]);
      }
    });
  });
  return [n + 1, p + 1];
};

const createScreen = (input) => {
  let screen = [];
  let size = sizeScreen(input);
  let n = size[0];
  let p = size[1];
  L = [];
  for (j = 0; j < p; j++) {
    L.push(".");
  }
  for (i = 0; i < n; i++) {
    screen.push([...L]);
  }
  return screen;
};

const printScreen = (screen) => {
  let draw = "";
  let n = screen.length;
  let p = screen[0].length;
  screen.forEach((line) => {
    line.forEach((pix, ind) => {
      draw += pix;
      if (ind == p - 1) {
        draw += "\n";
      }
    });
  });
  console.log(draw);
};

const sandFall = (input) => {
  let landscape = drawPath(input);
  let flow = false;
  let c = 0;
  while (!flow) {
    flow = sandUnitFall(landscape);
    c++;
  }
  return c - 1;
};
const sandUnitFall = (landscape) => {
  let blocked = false;
  let flow = false;
  let pos = [500, 0]; // x, y
  while (!blocked && !flow) {
    // check blocked
    if (landscape[pos[1] + 1][pos[0]] == ".") {
      pos = [pos[0], pos[1] + 1];
    } else if (landscape[pos[1] + 1][pos[0] - 1] == ".") {
      pos = [pos[0] - 1, pos[1] + 1];
    } else if (landscape[pos[1] + 1][pos[0] + 1] == ".") {
      pos = [pos[0] + 1, pos[1] + 1];
    } else {
      blocked = true;
    }
    if (
      pos[1] + 1 > landscape.length - 1 ||
      pos[0] - 1 < 0 ||
      pos[0] + 1 > landscape[0].length - 1
    ) {
      // check flow
      flow = true;
    }
  }
  landscape[pos[1]][pos[0]] = "o";
  return flow;
};

//printScreen(drawPath(input));
//console.log(sizeScreen(input));
//console.log(printScreen(createScreen(input)));

part1 = sandFall(input);

// Part 2

const drawPath2 = (input) => {
  let screen = createScreen2(input);
  input.forEach((coords) => {
    coords = coords.split(" -> ");
    //console.log(coords);
    for (k = 0; k < coords.length; k++) {
      coords[k] = coords[k].split(",").map(Number);
      coords[k][0] = coords[k][0] + Math.trunc(screen[0].length / 4);
      //console.log(coords[k]);
    }
    for (k = 1; k < coords.length; k++) {
      //console.log(coords[k - 1], coords[k]);
      //printScreen(screen);
      if (coords[k][0] == coords[k - 1][0]) {
        //draw vert
        a = coords[k - 1][1];
        b = coords[k][1];
        if (coords[k - 1][1] > coords[k][1]) {
          a = coords[k][1];
          b = coords[k - 1][1];
        }
        for (pix = a; pix <= b; pix++) {
          screen[pix][coords[k][0]] = "#";
          //console.log("1", pix, coords[k][1]);
          //printScreen(screen);
        }
      }
      if (coords[k][1] == coords[k - 1][1]) {
        //draw hor
        a = coords[k - 1][0];
        b = coords[k][0];
        if (coords[k - 1][0] > coords[k][0]) {
          a = coords[k][0];
          b = coords[k - 1][0];
        }
        for (pix = a; pix <= b; pix++) {
          screen[coords[k][1]][pix] = "#";
          //console.log("2", pix, coords[k][0]);
          //printScreen(screen);
        }
      }
    }
  });
  // CAREFUL TO MODIFY FOR REAL INPUT
  screen[0][500 + Math.trunc(screen[0].length / 4)] = "+";
  return screen;
};

const sizeScreen2 = (input) => {
  let n = 0;
  let p = 0;
  input.forEach((line) => {
    line = line.split(" -> ");
    line.forEach((coord) => {
      coord = coord.split(",");
      if (coord[0] > p) {
        p = Number(coord[0]);
      }
      if (coord[1] > n) {
        n = Number(coord[1]);
      }
    });
  });
  return [n + 1, 4 * (p + 1)];
};

const createScreen2 = (input) => {
  let screen = [];
  let size = sizeScreen2(input);
  let n = size[0];
  let p = size[1];
  L = [];
  for (j = 0; j < p; j++) {
    L.push(".");
  }
  for (i = 0; i < n; i++) {
    screen.push([...L]);
  }
  // Part 2
  screen.push([...L]);
  let M = [];
  for (j = 0; j < p; j++) {
    M.push("#");
  }
  screen.push([...M]);
  return screen;
};

const sandFall2 = (input) => {
  let landscape = drawPath2(input);
  let flow = false;
  let pyramid = false;
  let c = 0;
  while (!flow && !pyramid) {
    let b = sandUnitFall2(landscape);
    flow = b[0];
    pyramid = b[1];
    c++;
  }
  return c;
};
const sandUnitFall2 = (landscape) => {
  let blocked = false;
  let flow = false;
  let pyramid = false;
  let pos = [500 + Math.trunc(landscape[0].length / 4), 0]; // x, y
  while (!blocked && !flow) {
    // check blocked
    if (landscape[pos[1] + 1][pos[0]] == ".") {
      pos = [pos[0], pos[1] + 1];
    } else if (landscape[pos[1] + 1][pos[0] - 1] == ".") {
      pos = [pos[0] - 1, pos[1] + 1];
    } else if (landscape[pos[1] + 1][pos[0] + 1] == ".") {
      pos = [pos[0] + 1, pos[1] + 1];
    } else {
      blocked = true;
    }
    if (
      pos[1] + 1 > landscape.length - 1 ||
      pos[0] - 1 < 0 ||
      pos[0] + 1 > landscape[0].length - 1
    ) {
      // check flow
      flow = true;
    }
    if (pos[1] == 0) {
      pyramid = true;
    }
  }
  landscape[pos[1]][pos[0]] = "o";
  //printScreen(landscape);
  return [flow, pyramid];
};

part2 = sandFall2(input);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);
console.log("Part 2 :", part2);
