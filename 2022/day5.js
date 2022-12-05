f = "day5/input.txt";
const fs = require("fs");
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").split("\n\n");
let start = performance.now();

// Part 1
const arrangeStacks = (input) => {
  stacks = input[0].split("\n");
  instructions = input[1].split("\n");
  stackArray = buildStackArray(stacks);
  instructions.forEach((inst) => {
    inst = inst.split(" ");
    qt = Number(inst[1]);
    from = Number(inst[3]) - 1;
    to = Number(inst[5]) - 1;
    for (k = 0; k < qt; k++) {
      el = stackArray[from].shift();
      stackArray[to].unshift(el);
    }
  });
  sol = "";
  stackArray.forEach((l) => {
    sol += l[0];
  });
  return sol;
};

const move = (inst, { stackArray }) => {};

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
Part1 = arrangeStacks(input);
// Part 2
const arrangeStacks2 = (input) => {
  stacks = input[0].split("\n");
  instructions = input[1].split("\n");
  stackArray = buildStackArray(stacks);
  instructions.forEach((inst) => {
    inst = inst.split(" ");
    qt = Number(inst[1]);
    from = Number(inst[3]) - 1;
    to = Number(inst[5]) - 1;
    el = stackArray[from].slice(0, qt);
    stackArray[from] = stackArray[from].slice(qt);
    stackArray[to].unshift(...el);
  });

  sol = "";
  stackArray.forEach((l) => {
    sol += l[0];
  });
  return sol;
};

Part2 = arrangeStacks2(input);
let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");
console.log("Part 1 :", Part1);
console.log("Part 2 :", Part2);
