const fs = require("fs");
const fEx = "day15/ex.txt";
const inputEx = fs.readFileSync(`./inputs/${fEx}`, "utf-8").trim().split("\n");

const fEx2 = "day15/ex2.txt";
const inputEx2 = fs
  .readFileSync(`./inputs/${fEx2}`, "utf-8")
  .trim()
  .split("\n");

const f = "day15/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");

// Part 1
let start = performance.now();

function mergeIntervals(intervals) {
  if (intervals.length <= 1) return intervals;
  var stack = [];
  var top = [];
  intervals = intervals.sort(function (a, b) {
    return a[0] - b[0];
  });
  stack.push(intervals[0]);
  for (var i = 1; i < intervals.length; i++) {
    top = stack[stack.length - 1];
    if (top[1] < intervals[i][0]) {
      stack.push(intervals[i]);
    } else if (top[1] < intervals[i][1]) {
      top[1] = intervals[i][1];
      stack.pop();
      stack.push(top);
    }
  }
  return stack;
}

class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
  static dst(a, b) {
    let dx = Math.abs(a.x - b.x);
    let dy = Math.abs(a.y - b.y);
    return dx + dy;
  }
  equals(p) {
    return this.x == p.x && this.y == p.y;
  }
}

const findNbOccupied = (input, y0) => {
  let I = [];
  let map = new Map();
  input.forEach((line) => {
    line = line.replaceAll("x=", "");
    line = line.replaceAll("y=", "");
    line = line.split(": closest beacon is at ");
    sensor = line[0].split("Sensor at ")[1].split(", ").map(Number);
    beacon = line[1].split(", ").map(Number);
    let S = new Point(sensor[0], sensor[1]);
    let B = new Point(beacon[0], beacon[1]);
    addS = true;
    addB = true;
    for (let key of map.keys()) {
      if (S.equals(key)) {
        addS = false;
      }
      if (B.equals(key)) {
        addB = false;
      }
    }
    if (addS == true) {
      map.set(S, "S");
    }
    if (addB == true) {
      map.set(B, "B");
    }
    let d = Point.dst(S, B);
    if (S.y - d <= y0 && y0 <= S.y) {
      let i = y0 + d - S.y;
      I.push([S.x - i, S.x + i]);
    }
    if (S.y + d >= y0 && y0 > S.y) {
      let i = S.y + d - y0;
      I.push([S.x - i, S.x + i]);
    }
  });
  I = mergeIntervals(I);
  let res = 0;
  map.forEach((v, p) => {
    I.forEach((interval) => {
      if (p.y == y0 && interval[0] <= p.x && p.x <= interval[1]) {
        res -= 1;
      }
    });
  });
  I.forEach((arr) => {
    res += arr.at(-1) - arr.at(0) + 1;
  });
  return res;
};

part1 = findNbOccupied(input, 2000000);
part1Ex = findNbOccupied(inputEx, 10);
part1Ex2 = findNbOccupied(inputEx2, 3);

// Part 2
const fillLines = (Lines, S, y1, d) => {
  for (let j = S.y - d; j <= S.y - 1; j++) {
    let i = j - S.y + d;
    if (0 <= j && j <= y1) {
      if (S.x + i > y1) {
        if (S.x - i > y1) {
        } else if (0 <= S.x - i && S.x - i <= y1) {
          Lines[j].push([S.x - i, y1]);
        } else if (S.x - i < 0) {
          Lines[j].push([0, y1]);
        }
      } else if (0 <= S.x + i && S.x + i <= y1) {
        if (0 <= S.x - i && S.x - i <= y1) {
          Lines[j].push([S.x - i, S.x + i]);
        } else if (S.x - i < 0) {
          Lines[j].push([0, S.x + i]);
        }
      }
    }
    if (0 <= 2 * S.y - j && 2 * S.y - j <= y1) {
      if (S.x + i > y1) {
        if (S.x - i > y1) {
        } else if (0 <= S.x - i && S.x - i <= y1) {
          Lines[2 * S.y - j].push([S.x - i, y1]);
        } else if (S.x - i < 0) {
          Lines[2 * S.y - j].push([0, y1]);
        }
      } else if (0 <= S.x + i && S.x + i <= y1) {
        if (0 <= S.x - i && S.x - i <= y1) {
          Lines[2 * S.y - j].push([S.x - i, S.x + i]);
        } else if (S.x - i < 0) {
          Lines[2 * S.y - j].push([0, S.x + i]);
        }
      }
    }
    if (0 <= S.y && S.y <= y1) {
      if (S.x + d > y1) {
        if (S.x - d > y1) {
        } else if (0 <= S.x - d && S.x - d <= y1) {
          Lines[S.y].push([S.x - d, y1]);
        } else if (S.x - d < 0) {
          Lines[S.y].push([0, y1]);
        }
      } else if (0 <= S.x + d && S.x + d <= y1) {
        if (0 <= S.x - d && S.x - d <= y1) {
          Lines[S.y].push([S.x - d, S.x + d]);
        } else if (S.x - d < 0) {
          Lines[S.y].push([0, S.x + d]);
        }
      }
    }
  }
};

const findSignal = (input, y1) => {
  let set = new Set();
  let Lines = [];
  let res = [];
  for (k = 0; k <= y1; k++) {
    Lines.push([]);
    res.push(0);
  }
  input.forEach((line) => {
    line = line.replaceAll("x=", "");
    line = line.replaceAll("y=", "");
    line = line.split(": closest beacon is at ");
    sensor = line[0].split("Sensor at ")[1].split(", ").map(Number);
    beacon = line[1].split(", ").map(Number);
    addS = true;
    addB = true;
    let S = new Point(sensor[0], sensor[1]);
    let B = new Point(beacon[0], beacon[1]);
    for (let key of set.keys()) {
      if (S.equals(key)) {
        addS = false;
      }
      if (B.equals(key)) {
        addB = false;
      }
    }
    if (addS == true) {
      set.add(S);
    }
    if (addB == true) {
      set.add(B);
    }
    let d = Point.dst(S, B);
    fillLines(Lines, S, y1, d);
  });

  for (k = 0; k <= y1; k++) {
    Lines[k] = mergeIntervals(Lines[k]);
    Lines[k].forEach((arr) => {
      res[k] += arr.at(-1) - arr.at(0) + 1;
    });
    if (res[k] == y1) {
      let ySignal = k;
      for (let r = 0; r < Lines[k].length - 1; r++) {
        if (Lines[k][r][1] + 1 != Lines[k][r + 1][0]) {
          xSignal = Lines[k][r][1] + 1;
        }
      }
      if (Lines[k].at(-1)[1] != y1) {
        xSignal = y1;
      }
      return xSignal * 4000000 + ySignal;
    }
  }
};

part2Ex2 = findSignal(inputEx2, 5);
part2Ex = findSignal(inputEx, 20);
//part2 = findSignal(input, 4000000);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 Ex 2 :", part1Ex2);
console.log("Part 1 Ex :", part1Ex);
console.log("Part 1 :", part1);

console.log("Part 2 Ex 2 :", part2Ex2);
console.log("Part 2 Ex :", part2Ex);
//console.log("Part 2 :", part2);

let start2 = performance.now();
a = 0;
for (k = 0; k <= 4000000; k++) {
  a++;
}
let end2 = performance.now();
console.log("Execution time :", (end2 - start2).toFixed(2), "ms");
