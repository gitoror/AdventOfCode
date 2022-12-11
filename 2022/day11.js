const fs = require("fs");
const f = "day11/input.txt";
const input = fs
  .readFileSync(`./inputs/${f}`, "utf-8")
  .trim()
  .split("\n\n")
  .map((monkeyInfo) => {
    monkeyInfo = monkeyInfo.split("\n").map((line) => {
      return line.trim();
    });
    return monkeyInfo;
  });

// Part 1
let start = performance.now();

function monkeysInit(input) {
  let monkeys = [];
  let monkeysActivity = [];
  let divAll = 1;
  for (m = 0; m < input.length; m++) {
    monkeysActivity.push(0);
    monkey = input[m];
    startingItems = monkey[1]
      .split(": ")[1]
      .split(", ")
      .map((n) => {
        return Number(n);
      });
    operation = monkey[2].split(": ")[1].split("old ")[1].split(" ");
    divideBy = Number(monkey[3].split(": ")[1].split("by ")[1]);
    divAll *= divideBy;
    throwTrue = Number(monkey[4].split(": ")[1].split("monkey ")[1]);
    throwFalse = Number(monkey[5].split(": ")[1].split("monkey ")[1]);
    monkeys.push({
      items: startingItems,
      op: operation,
      divideBy: divideBy,
      throwTrue: throwTrue,
      throwFalse: throwFalse,
    });
  }
  return [monkeys, monkeysActivity, divAll];
}

function monkeyOp(op, v, worry) {
  if (op == "+") {
    if (v == "old") {
      worry += worry;
    } else {
      worry += Number(v);
    }
  }
  if (op == "*") {
    if (v == "old") {
      worry = worry * worry;
    } else {
      worry = worry * Number(v);
    }
  }
  return worry;
}

function monkeyThrow(src, dest, worry, monkeys) {
  monkeys[dest].items.push(worry);
}

const monkeyBusiness = (input, part) => {
  let [monkeys, monkeysActivity, divAll] = monkeysInit(input);
  for (let round = 0; round < 10000; round++) {
    monkeys.forEach((monkey, src, monkeys) => {
      items = [...monkey.items];
      items.forEach((worryItem) => {
        worry = worryItem;
        op = monkey.op[0];
        v = monkey.op[1];
        worry = monkeyOp(op, v, worry);
        monkeysActivity[src]++;
        if (part == 1) {
          worry = Math.trunc(worry / 3);
          if (worry % monkey.divideBy == 0) {
            monkeyThrow(src, monkey.throwTrue, worry, monkeys);
          } else {
            monkeyThrow(src, monkey.throwFalse, worry, monkeys);
          }
        } else if (part == 2) {
          if (worry % monkey.divideBy == 0) {
            monkeyThrow(src, monkey.throwTrue, worry % divAll, monkeys);
          } else {
            monkeyThrow(src, monkey.throwFalse, worry % divAll, monkeys);
          }
        }
      });
      monkeys[src].items = [];
    });
  }
  monkeysActivity = monkeysActivity.sort((a, b) => b - a);
  let monkeyBusiness = monkeysActivity[0] * monkeysActivity[1];
  return monkeyBusiness;
};

part1 = monkeyBusiness(input, 1);
part2 = monkeyBusiness(input, 2);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);
console.log("Part 2 :", part2);
