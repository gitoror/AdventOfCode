const fs = require("fs");
const f = "day9/input.txt";
const input = fs
  .readFileSync(`./inputs/${f}`, "utf-8")
  .trim()
  .split("\n")
  .map((line) => {
    line = line.split(" ");
    line[1] = Number(line[1]);
    return line;
  });
let start = performance.now();

// Part 1

const countVisitedPos = (input) => {
  visited = new Set();
  H = { x: 0, y: 0 };
  T = { x: 0, y: 0 };
  visited.add(`${0},${0}`);
  input.forEach((line) => {
    dir = line[0];
    dl = line[1];
    for (i = 0; i < dl; i++) {
      if (dir == "U") {
        H.y++;
        if (!isAdjacent(H, T)) {
          T.y++;
          if (T.x < H.x) {
            T.x++;
          } else if (T.x > H.x) {
            T.x--;
          }
        }
      }
      if (dir == "D") {
        H.y--;
        if (!isAdjacent(H, T)) {
          T.y--;
          if (T.x < H.x) {
            T.x++;
          } else if (T.x > H.x) {
            T.x--;
          }
        }
      }
      if (dir == "R") {
        H.x++;
        if (!isAdjacent(H, T)) {
          T.x++;
          if (T.y < H.y) {
            T.y++;
          } else if (T.y > H.y) {
            T.y--;
          }
        }
      }
      if (dir == "L") {
        H.x--;
        if (!isAdjacent(H, T)) {
          T.x--;
          if (T.y < H.y) {
            T.y++;
          } else if (T.y > H.y) {
            T.y--;
          }
        }
      }
      if (!visited.has(`${T.x},${T.y}`)) {
        visited.add(`${T.x},${T.y}`);
      }
    }
  });
  //console.log(visited);
  return visited.size;
};

function isAdjacent(H, T) {
  return (
    (T.x == H.x && T.y == H.y) ||
    (T.x == H.x && (T.y == H.y + 1 || T.y == H.y - 1)) ||
    (T.y == H.y && (T.x == H.x + 1 || T.x == H.x - 1)) ||
    (T.x == H.x && (T.y == H.y + 1 || T.y == H.y - 1)) ||
    (T.x == H.x - 1 && T.y == H.y - 1) ||
    (T.x == H.x - 1 && T.y == H.y + 1) ||
    (T.x == H.x + 1 && T.y == H.y + 1) ||
    (T.x == H.x + 1 && T.y == H.y - 1)
  );
}

part1 = countVisitedPos(input);

// Part 2

function countVisitedLongRope(input, ropeSize) {
  visited = new Set();
  rope = [];
  for (k = 0; k < ropeSize; k++) {
    rope.push({ x: 0, y: 0 });
  }
  visited.add(`${0},${0}`);
  input.forEach((line) => {
    dir = line[0];
    dl = line[1];
    head = 0;
    for (k = 0; k < dl; k++) {
      moveHead2(rope, head, dir, visited);
    }
  });
  return visited.size;
}

function moveHead2(rope, head, dir, visited) {
  if (dir == "U") {
    rope[head].y++;
  } else if (dir == "D") {
    rope[head].y--;
  } else if (dir == "R") {
    rope[head].x++;
  } else if (dir == "L") {
    rope[head].x--;
  } else if (dir == "UL") {
    rope[head].y++;
    rope[head].x--;
  } else if (dir == "UR") {
    rope[head].y++;
    rope[head].x++;
  } else if (dir == "DL") {
    rope[head].y--;
    rope[head].x--;
  } else if (dir == "DR") {
    rope[head].y--;
    rope[head].x++;
  }
  followHead2(rope, head, visited);
}

function followHead2(rope, head, visited) {
  if (head <= rope.length - 2) {
    let tail = head + 1;
    Hx = rope[head].x;
    Hy = rope[head].y;
    Tx = rope[tail].x;
    Ty = rope[tail].y;
    if (isAdjacent(rope[head], rope[tail])) {
    } else {
      if (Hx == Tx - 2) {
        if (Hy == Ty + 2 || Hy == Ty + 1) {
          moveHead2(rope, tail, "UL", visited);
        } else if (Hy == Ty) {
          moveHead2(rope, tail, "L", visited);
        } else if (Hy == Ty - 1 || Hy == Ty - 2) {
          moveHead2(rope, tail, "DL", visited);
        }
      } else if (Hx == Tx - 1) {
        if (Hy == Ty + 2) {
          moveHead2(rope, tail, "UL", visited);
        } else if (Hy == Ty - 2) {
          moveHead2(rope, tail, "DL", visited);
        }
      } else if (Hx == Tx) {
        if (Hy == Ty + 2) {
          moveHead2(rope, tail, "U", visited);
        } else if (Hy == Ty - 2) {
          moveHead2(rope, tail, "D", visited);
        }
      } else if (Hx == Tx + 1) {
        if (Hy == Ty + 2) {
          moveHead2(rope, tail, "UR", visited);
        } else if (Hy == Ty - 2) {
          moveHead2(rope, tail, "DR", visited);
        }
      } else if (Hx == Tx + 2) {
        if (Hy == Ty + 2 || Hy == Ty + 1) {
          moveHead2(rope, tail, "UR", visited);
        } else if (Hy == Ty) {
          moveHead2(rope, tail, "R", visited);
        } else if (Hy == Ty - 1 || Hy == Ty - 2) {
          moveHead2(rope, tail, "DR", visited);
        }
      }
    }
    if (
      tail == rope.length - 1 &&
      !visited.has(`${rope[tail].x},${rope[tail].y}`)
    ) {
      visited.add(`${rope[tail].x},${rope[tail].y}`);
    }
  } else {
  }
}

part2 = countVisitedLongRope(input, 10);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);

console.log("Part 2 :", part2);
