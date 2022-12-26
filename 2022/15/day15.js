const fs = require("fs");
const fEx = "day15/ex.txt";
const inputEx = fs.readFileSync(`./inputs/${fEx}`, "utf-8").trim().split("\n");

const fE2 = "day15/ex2.txt";
const inputEx2 = fs.readFileSync(`./inputs/${fE2}`, "utf-8").trim().split("\n");

const f = "day15/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");

const Execute = (fn, parameters, msg = "Result :") => {
  let start = performance.now();
  let res = fn(...parameters);
  let end = performance.now();
  console.log(msg, res, "Time :", (end - start).toFixed(2), "ms");
};

// Part 1

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

const dstManhattan = (S, B) => {
  return Math.abs(S[0] - B[0]) + Math.abs(S[1] - B[1]);
};

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
      if (
        p.y == y0 &&
        interval[0] <= p.x &&
        p.x <= interval[1] &&
        map.get(p) == "B"
      ) {
        res -= 1;
      }
    });
  });
  I.forEach((arr) => {
    res += arr.at(-1) - arr.at(0) + 1;
  });
  return res;
};

// Part 2

const Sensors = (input) => {
  let sensors = [];
  input.forEach((line) => {
    line = line.replaceAll("x=", "");
    line = line.replaceAll("y=", "");
    line = line.split(": closest beacon is at ");
    sensor = line[0].split("Sensor at ")[1].split(", ").map(Number);
    beacon = line[1].split(", ").map(Number);
    d = dstManhattan(sensor, beacon);
    sensors.push({ x: sensor[0], y: sensor[1], d: d });
  });

  return sensors;
};

const findSignal2 = (input, MAX) => {
  let sensors = Sensors(input);
  for (let S of sensors) {
    let noBeaconsZone = [];
    for (let i = 0; i <= S.d + 1; i++) {
      noBeaconsZone.push(`${S.x - S.d - 1 + i} ${S.y - i}`);
      noBeaconsZone.push(`${S.x + S.d + 1 - i} ${S.y - i}`);
      noBeaconsZone.push(`${S.x - S.d - 1 + i} ${S.y + i}`);
      noBeaconsZone.push(`${S.x + S.d + 1 - i} ${S.y + i}`);
    }
    for (let point of noBeaconsZone) {
      point = point.split(" ").map(Number);
      if (
        point[0] >= 0 &&
        point[0] <= MAX &&
        point[1] >= 0 &&
        point[1] <= MAX
      ) {
        let count = 0;
        for (let S of sensors) {
          let s = [S.x, S.y];
          if (dstManhattan(s, point) > S.d) {
            count++;
          }
        }
        if (count == sensors.length) {
          return point[0] * 4000000 + point[1];
        }
      }
    }
  }
};

Execute(findNbOccupied, [inputEx, 10], "Part 1 ex :");
Execute(findNbOccupied, [input, 2000000], "Part 1 :");
Execute(findSignal, [inputEx, 20], "Part 2 ex :");
Execute(findSignal2, [input, 4000000], "Part 2 :");

// const findSignal = (input, MAX) => {
//   let sensors = Sensors(input);
//   sensors = sensors.sort((s1, s2) => {
//     if (s1.x - s2.x == 0) {
//       return s1.y - s2.y;
//     }
//     return s1.x - s2.x;
//   });
//   let Lines = [];
//   for (i = 0; i <= MAX; i++) {
//     Lines.push({ intervals: [], occupation: 0 });
//   }
//   for (let S of sensors) {
//     for (let i = Math.max(0, S.y - S.d); i <= Math.min(MAX, S.y + S.d); i++) {
//       let L = S.d - Math.abs(i - S.y);
//       newI = [Math.max(0, S.x - L), Math.min(MAX, S.x + L)];
//       if (Lines[i].intervals.length == 0) {
//         Lines[i].intervals.push(newI);
//         Lines[i].occupation += newI[1] - newI[0] + 1;
//       } else {
//         precI = Lines[i].intervals.at(-1);
//         let a0 = precI[0];
//         let b0 = precI[1];
//         let a1 = newI[0];
//         let b1 = newI[1];
//         if (a0 <= a1 && a1 <= b0 && b0 < b1) {
//           Lines[i].intervals.at(-1)[1] = b1;
//           Lines[i].occupation += b1 - b0;
//         } else if (a1 < a0 && b0 < b1) {
//           Lines[i].intervals.at(-1)[0] = a1;
//           Lines[i].intervals.at(-1)[1] = b1;
//           Lines[i].occupation += b1 - b0 + a0 - a1;
//         } else if (a1 < a0 && a0 <= b1 && b1 <= b0) {
//           Lines[i].intervals.at(-1)[0] = a1;
//           Lines[i].occupation += a0 - a1;
//         } else if (b0 < a1) {
//           Lines[i].intervals.push(newI);
//           Lines[i].occupation += b1 - a1 + 1;
//         } else if (b1 < a0) {
//           Lines[i].intervals.push(newI);
//           Lines[i].occupation += b1 - a1 + 1;
//         }
//       }
//     }
//   }
//   for (i = 0; i <= MAX; i++) {
//     if (Lines[i].occupation == MAX) {
//       ySignal = i;
//       for (let r = 0; r < Lines[i].intervals.length - 1; r++) {
//         if (Lines[i].intervals[r][1] + 1 != Lines[i].intervals[r + 1][0]) {
//           xSignal = Lines[i].intervals[r][1] + 1;
//         }
//       }
//       if (Lines[i].intervals.at(-1)[1] != MAX) {
//         xSignal = MAX;
//       }
//       return xSignal * 4000000 + ySignal;
//     }
//   }
// };
