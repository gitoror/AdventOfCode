const fs = require("fs");
const f = "day10/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");
let start = performance.now();

// Part 1

const signalStrenght = (input) => {
  ss = 0;
  cycle = 0;
  X = 1;
  input.forEach((line) => {
    line = line.split(" ");
    if (line.length == 1) {
      cycle++;
      if ([20, 60, 100, 140, 180, 220].includes(cycle)) {
        ss += cycle * X;
      }
    } else if (line.length == 2) {
      cycle++;
      if ([20, 60, 100, 140, 180, 220].includes(cycle)) {
        ss += cycle * X;
      }
      cycle++;
      if ([20, 60, 100, 140, 180, 220].includes(cycle)) {
        ss += cycle * X;
      }
      X += Number(line[1]);
    }
  });
  return ss;
};

part1 = signalStrenght(input);

// Part 2

function drawScreen(input) {
  screen = "";
  cycle = 1;
  X = 1;
  input.forEach((line) => {
    line = line.split(" ");
    if (line.length == 1) {
      if ([X - 1, X, X + 1].includes((cycle - 1) % 40)) {
        screen += "#";
      } else {
        screen += ".";
      }
      if (cycle % 40 == 0) {
        screen += "\n";
      }
      cycle++;
    } else if (line.length == 2) {
      if ([X - 1, X, X + 1].includes((cycle - 1) % 40)) {
        screen += "#";
      } else {
        screen += ".";
      }
      if (cycle % 40 == 0) {
        screen += "\n";
      }
      cycle++;
      if ([X - 1, X, X + 1].includes((cycle - 1) % 40)) {
        screen += "#";
      } else {
        screen += ".";
      }
      if (cycle % 40 == 0) {
        screen += "\n";
      }
      cycle++;
      X += Number(line[1]);
    }
  });
  return screen;
}

part2 = drawScreen(input);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);

console.log("Part 2 :", "\n" + part2);
