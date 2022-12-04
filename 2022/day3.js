f = "day3/input";
const fs = require("fs");
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");
let start = performance.now();

// Part 1
const aplh = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
prio = {};
for (var l of aplh) {
  prio[l] = Object.keys(prio).length + 1;
}

const prioritySharedItems = (input) => {
  return input.reduce((acc, line) => {
    comp1 = line.slice(0, line.length / 2).split("");
    comp2 = line.slice(line.length / 2).split("");
    comp1.forEach((letter1) => {
      comp2.forEach((letter2) => {
        if (letter1 == letter2) {
          letter = letter1;
        }
      });
    });
    return acc + prio[letter];
  }, 0);
};

// Part 2
const priorityGroups = (input) => {
  count = 0;
  for (let i = 0; i < input.length; i += 3) {
    L1 = input[i].split("");
    L2 = input[i + 1].split("");
    L3 = input[i + 2].split("");
    L1.forEach((letter1) => {
      L2.forEach((letter2) => {
        L3.forEach((letter3) => {
          if ((letter3 == letter2) & (letter2 == letter1)) {
            letter = letter1;
          }
        });
      });
    });
    count += prio[letter];
  }
  return count;
};

Part1 = prioritySharedItems(input);
Part2 = priorityGroups(input);
let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", Part1); //7691
console.log("Part 2 :", Part2); //2508
