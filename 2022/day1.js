f = "day1/input.txt";
const fs = require("fs");
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n\n");
let start = performance.now();

// Part 1
const maxW = (input) => {
  weights = input.map((L) => {
    L = L.split("\n").map(Number);
    return L.reduce((acc, v) => acc + v);
  });
  return weights.sort().at(-1);
};

// Part 2
const top3 = (input) => {
  weights = input.map((W) => {
    W = W.split("\n").map(Number);
    return W.reduce((acc, v) => acc + v);
  });
  return weights
    .sort()
    .slice(-3)
    .reduce((acc, v) => acc + v);
};
P1 = maxW(input);
P2 = top3(input);
let end = performance.now();
console.log((end - start).toFixed(2));

console.log(P1);
console.log(P2);
