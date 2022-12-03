f = "day2/input.txt";
const fs = require("fs");
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");
let start = performance.now();
// Part 1
const points = { A: 0, B: 1, C: 2, X: 0, Y: 1, Z: 2 };

const countPoints = (input) => {
  victoryPoints = 0;
  choicePoints = 0;
  input.forEach((line) => {
    advChoice = line.split(" ")[0];
    yourChoice = line.split(" ")[1];
    choicePoints += points[yourChoice] + 1;
    if (points[advChoice] == points[yourChoice]) {
      victoryPoints += 3;
    } else if (points[yourChoice] == (((points[advChoice] + 1) % 3) + 3) % 3) {
      victoryPoints += 6;
    } else {
      victoryPoints += 0;
    }
  });
  return victoryPoints + choicePoints;
};

console.log(countPoints(input));

// Part 2
const countPoints2 = (input) => {
  victoryPoints = 0;
  choicePoints = 0;
  input.forEach((line) => {
    advChoice = line.split(" ")[0];
    yourChoice = line.split(" ")[1];
    if (yourChoice == "X") {
      victoryPoints += 0;
      choicePoints += ((((points[advChoice] - 1) % 3) + 3) % 3) + 1;
    }
    if (yourChoice == "Y") {
      victoryPoints += 3;
      choicePoints += points[advChoice] + 1;
    }
    if (yourChoice == "Z") {
      victoryPoints += 6;
      choicePoints += ((((points[advChoice] + 1) % 3) + 3) % 3) + 1;
    }
  });
  return victoryPoints + choicePoints;
};
console.log(countPoints2(input));
let end = performance.now();
console.log((end - start).toFixed(2));
