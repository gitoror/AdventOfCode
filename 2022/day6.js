const fs = require("fs");
const input = fs.readFileSync(`./inputs/day6/input.txt`, "utf-8").trim();
let start = performance.now();

// Part 1
const part1 = (input) => {
  i = 0;
  while (i < input.length) {
    if (
      input.at(i) == input.at(i + 1) ||
      input.at(i) == input.at(i + 2) ||
      input.at(i) == input.at(i + 3) ||
      input.at(i + 1) == input.at(i + 2) ||
      input.at(i + 1) == input.at(i + 3) ||
      input.at(i + 2) == input.at(i + 3)
    ) {
      i += 1;
    } else {
      return i + 4;
    }
  }
  return -1;
};

// Part 2

const part2 = (input) => {
  i = 0;
  while (i <= input.length - 14) {
    // find if letters all diff
    distChar = true;
    for (j = i; j <= i + 12; j++) {
      for (k = j + 1; k <= i + 13; k++) {
        if (input.at(j) == input.at(k)) {
          distChar = false;
        }
      }
    }
    if (distChar) {
      return i + 14;
    }
    i++;
  }
};

const compare = (i, j, input) => {
  return input.at(i) == input.at(j);
};

Part1 = part1(input);
Part2 = part2(input);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", Part1);
console.log("Part 2 :", Part2);
