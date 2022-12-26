const fs = require("fs");
const input = fs.readFileSync(`./inputs/day5/input.txt`, "utf-8").split("\n\n");
let start = performance.now();

// Part 1
const buildStackArray = (stacks) => {
  stackArray = [];
  nbStacks = Number(stacks.at(-1).at(-2));
  stacks.pop();
  for (k = 0; k < nbStacks; k++) {
    stackArray.push([]);
  }
  stacks.forEach((line) => {
    for (var j = 1; j < line.length; j += 4) {
      char = line.at(j);
      if (char != " ") {
        stackArray[Math.floor(j / 4)].push(line.at(j));
      }
    }
  });
  return stackArray;
};

const arrangeStacks = (input, move) => {
  stacks = input[0].split("\n");
  instructions = input[1].split("\n");
  stackArray = buildStackArray(stacks);
  instructions.forEach((inst) => {
    inst = inst.split(" ");
    qt = Number(inst[1]);
    from = Number(inst[3]) - 1;
    to = Number(inst[5]) - 1;
    move(qt, from, to, stackArray);
  });
  return stackArray.reduce((acc, stack) => {
    return acc + stack[0];
  }, "");
};

const move1 = (qt, from, to, stackArray) => {
  for (k = 0; k < qt; k++) {
    letter = stackArray[from].shift();
    stackArray[to].unshift(letter);
  }
};

Part1 = arrangeStacks(input, move1);

// Part 2
const move2 = (qt, from, to, stackArray) => {
  letter = stackArray[from].splice(0, qt);
  stackArray[to].unshift(...letter);
};

Part2 = arrangeStacks(input, move2);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", Part1);
console.log("Part 2 :", Part2);
