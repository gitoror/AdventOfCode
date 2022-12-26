const fs = require("fs");
const f = "day8/input.txt";
const input = fs.readFileSync(`./inputs/${f}`, "utf-8").trim().split("\n");
for (i = 0; i < input.length; i++) {
  line = input[i];
  input[i] = line.split("").map(Number);
}
let start = performance.now();

// Part 1
function isVisible(input, i, j) {
  v = input[i][j];
  vU = true;
  vD = true;
  vL = true;
  vR = true;
  for (u = i - 1; u >= 0; u--) {
    if (input[u][j] >= v) {
      vU = false;
      break;
    }
  }
  for (d = i + 1; d < input.length; d++) {
    if (input[d][j] >= v) {
      vD = false;
      break;
    }
  }
  for (l = j - 1; l >= 0; l--) {
    if (input[i][l] >= v) {
      vL = false;
      break;
    }
  }
  for (r = j + 1; r < input[0].length; r++) {
    if (input[i][r] >= v) {
      vR = false;
      break;
    }
  }
  return vU || vD || vL || vR;
}

function countVisible(input) {
  c = 2 * input.length + 2 * (input[0].length - 2);
  for (let i = 1; i < input.length - 1; i++) {
    for (let j = 1; j < input[0].length - 1; j++) {
      if (isVisible(input, i, j)) {
        c++;
      }
    }
  }
  return c;
}

part1 = countVisible(input);

// Part 2
function scenicScore(input, i, j) {
  v = input[i][j];
  sU = 0;
  sD = 0;
  sL = 0;
  sR = 0;
  for (u = i - 1; u >= 0; u--) {
    sU += 1;
    if (input[u][j] >= v) {
      break;
    }
  }
  for (d = i + 1; d < input.length; d++) {
    sD += 1;
    if (input[d][j] >= v) {
      break;
    }
  }
  for (l = j - 1; l >= 0; l--) {
    sL += 1;
    if (input[i][l] >= v) {
      break;
    }
  }
  for (r = j + 1; r < input[0].length; r++) {
    sR += 1;
    if (input[i][r] >= v) {
      break;
    }
  }
  return sR * sL * sU * sD;
}

function bestScenicScore(input) {
  m = 0;
  for (let i = 1; i < input.length - 1; i++) {
    for (let j = 1; j < input[0].length - 1; j++) {
      s = scenicScore(input, i, j);
      if (s > m) {
        m = s;
      }
    }
  }
  return m;
}
part2 = bestScenicScore(input);
let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);
console.log("Part 2 :", part2);
