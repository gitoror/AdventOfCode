const fs = require("fs");
const f = "day13/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n\n");

// Part 1
let start = performance.now();
const findPairsOrdered = (input) => {
  let sum = 0;
  input.forEach((pair, ind) => {
    pair = pair.split("\n");
    let left = pair[0];
    let right = pair[1];
    if (isOrdered(toArray(left), toArray(right))) {
      sum += ind + 1;
    }
  });
  return sum;
};

const isOrdered = (left, right) => {
  if (left.length == 0 && right.length == 0) {
    return 0;
  } else if (left.length == 0 && right.length != 0) {
    return true;
  } else if (right.length == 0 && left.length != 0) {
    return false;
  } else {
    if (left[0] instanceof Array && right[0] instanceof Array) {
      let b = isOrdered([...left[0]], [...right[0]]);
      if (b === true) {
        return true;
      } else if (b === false) {
        return false;
      }
    } else if (left[0] instanceof Array) {
      let b = isOrdered([...left[0]], [right[0]]);
      if (b === true) {
        return true;
      } else if (b === false) {
        return false;
      }
    } else if (right[0] instanceof Array) {
      let b = isOrdered([left[0]], [...right[0]]);
      if (b === true) {
        return true;
      } else if (b === false) {
        return false;
      }
    } else if (left[0] > right[0]) {
      return false;
    } else if (left[0] < right[0]) {
      return true;
    }
    return isOrdered(left.slice(1), right.slice(1));
  }
};

const toArray = (arr) => {
  return JSON.parse(arr);
};

part1 = findPairsOrdered(input);

// Part 2
L = [[[2]], [[6]]];
input.forEach((pair) => {
  pair = pair.split("\n");
  L.push(toArray(pair[0]));
  L.push(toArray(pair[1]));
});

L = L.sort((A, B) => {
  if (isOrdered(A, B)) {
    return -1;
  } else {
    return 1;
  }
});

L.forEach((l, i) => {
  if (l.toString() == [[2]]) {
    index2 = i + 1;
  }
  if (l.toString() == [[6]]) {
    index6 = i + 1;
  }
});

part2 = index2 * index6;

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);
console.log("Part 2 :", part2);
