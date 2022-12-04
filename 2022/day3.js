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
    S1 = new Set();
    S2 = new Set();
    for (var l of line.slice(0, line.length / 2)) {
      S1.add(l);
    }
    for (var l of line.slice(line.length / 2)) {
      S2.add(l);
    }
    S = new Set([...S1].filter((x) => S2.has(x)));
    return acc + prio[S.values().next().value];
  }, 0);
};

// Part 2
const priorityGroups = (input) => {
  count = 0;
  for (let i = 0; i < input.length; i += 3) {
    S1 = new Set();
    S2 = new Set();
    S3 = new Set();
    for (var l of input[i]) {
      S1.add(l);
    }
    for (var l of input[i + 1]) {
      S2.add(l);
    }
    for (var l of input[i + 2]) {
      S3.add(l);
    }
    Si = new Set([...S1].filter((x) => S2.has(x)));
    S = new Set([...Si].filter((x) => S3.has(x)));
    count += prio[S.values().next().value];
  }
  return count;
};

Part1 = prioritySharedItems(input);
Part2 = priorityGroups(input);
let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", Part1); //7691
console.log("Part 2 :", Part2); //2508

/*

Execution time : 75.68 ms
Part 1 : 7691
Part 2 : 2508
*/
