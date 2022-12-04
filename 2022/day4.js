f = "day4/input";
const fs = require("fs");
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");

// Part 1
let start = performance.now();

const containedPairs = (input) => {
  return input.filter((line) => {
    line = line.split(",");
    I1 = line[0].split("-");
    I2 = line[1].split("-");
    a1 = Number(I1[0]);
    b1 = Number(I1[1]);
    a2 = Number(I2[0]);
    b2 = Number(I2[1]);
    if ((a1 <= a2 && b2 <= b1) || (a2 <= a1 && b1 <= b2)) {
      return true;
    }
  }).length;
};

Part1 = containedPairs(input);

// Part 2
const containedPairs2 = (input) => {
  return input.filter((line) => {
    line = line.split(",");
    I1 = line[0].split("-");
    I2 = line[1].split("-");
    a1 = Number(I1[0]);
    b1 = Number(I1[1]);
    a2 = Number(I2[0]);
    b2 = Number(I2[1]);
    if (
      (a1 <= a2 && a2 <= b1) ||
      (a1 <= b2 && b2 <= b1) ||
      (a2 <= a1 && a1 <= b2) ||
      (a2 <= b1 && b1 <= b2)
    ) {
      return true;
    }
  }).length;
};
Part2 = containedPairs2(input);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1", Part1);
console.log("Part 2", Part2);
