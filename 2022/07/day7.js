const fs = require("fs");
const f = "day7/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");
let start = performance.now();

// Part 1
const buildFileTree = (input) => {
  tree = { "/": {} };
  path = [];
  input.forEach((line) => {
    line = line.split(" ");
    if (line[0] == "$") {
      if (line[1] == "cd") {
        if (line[2] == "..") {
          path.pop();
        } else {
          path.push(line[2]);
        }
      }
    } else {
      treeToPush = tree;
      path.forEach((dir) => {
        treeToPush = treeToPush[dir];
      });
      if (line[0] == "dir") {
        treeToPush[line[1]] = {};
      } else {
        treeToPush[line[1]] = Number(line[0]);
      }
    }
  });
  return tree;
};

function sizeDir(tree, possibleDir = [], minDelete = 0) {
  let s = 0;
  let tot = 0;
  Object.keys(tree).forEach((key) => {
    if (typeof tree[key] == "object") {
      mem = sizeDir(tree[key], possibleDir, minDelete);
      newS = mem[0];
      newTot = mem[1];
      if (newS < 100000) {
        tot += newS;
      }
      s += newS;
      tot += newTot;
      // part 2
      if (newS >= minDelete) {
        possibleDir.push(newS);
      }
    } else {
      s += tree[key];
    }
  });
  return [s, tot];
}

// Part 2
function findDeleteDir(tree) {
  possibleDir = [];
  diskSize = part1[0];
  minDelete = diskSize - 40000000;
  sizeDir(tree, possibleDir, minDelete);
  return possibleDir.sort((a, b) => a - b)[0];
}
tree = buildFileTree(input)["/"];
part1 = sizeDir(tree);
part2 = findDeleteDir(tree);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1[1]);
console.log("Part 2 :", part2);
