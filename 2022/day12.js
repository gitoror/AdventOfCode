const fs = require("fs");
const f = "day12/input.txt";
const input = fs
  .readFileSync(`./inputs/${f}`, "utf-8")
  .trim()
  .split("\n")
  .map((line) => {
    return line.split("");
  });
// console.log(input);

// Part 1
let start = performance.now();

const aplh = "abcdefghijklmnopqrstuvwxyz";
let H = {};
for (var l of aplh) {
  H[l] = Object.keys(H).length + 1;
}
H["S"] = 1;
H["E"] = 26;
let Inf = 1234654567898;

const dijkstra = (input) => {
  let V = new Set();
  let N = input[0].length * input.length;
  let d = [];
  let pred = [];
  for (let k = 0; k < N; k++) {
    V.add(k);
    d[k] = Inf;
    pred[k] = 0;
  }
  let sCoord = posInit(input, "S");
  let s = sCoord[0] * input[0].length + sCoord[1];
  let eCoord = posInit(input, "E");
  let e = eCoord[0] * input[0].length + eCoord[1];
  d[s] = 0;
  let found = false;
  while (V.size != 0 && !found) {
    let a = minD(V, d);
    V.delete(a);
    if (a == e) {
      found = true;
    } else {
      let voisins = findVoisins(input, a);
      let intersect = new Set([...V].filter((i) => voisins.has(i)));
      for (let b of intersect) {
        if (d[b] > d[a] + poids(a, b, input)) {
          d[b] = d[a] + poids(a, b, input);
          pred[b] = a;
        }
      }
    }
  }
  let chemin = [];
  knot = e;
  while (knot != s) {
    chemin.unshift(knot);
    knot = pred[knot];
  }
  chemin.unshift(s);
  return chemin.length - 1;
};

function giveCoord(input, a) {
  let ai = Math.trunc(a / input[0].length);
  let aj = a % input[0].length;
  return [ai, aj];
}

function poids(a, b, input) {
  let ai = Math.trunc(a / input[0].length);
  let aj = a % input[0].length;
  let bi = Math.trunc(b / input[0].length);
  let bj = b % input[0].length;
  let d = H[input[bi][bj]] - H[input[ai][aj]];
  if (d <= 1) {
    return 1;
  } else {
    return Inf;
  }
}

function posInit(input, letter) {
  input.forEach((line, i) => {
    line.forEach((char, j) => {
      if (char == letter) {
        R = [i, j];
      }
    });
  });
  return R;
}

function minD(V, d) {
  let min = 502165464500;
  let sMin = -1;
  V.forEach((v) => {
    if (d[v] < min) {
      min = d[v];
      sMin = v;
    }
  });
  return sMin;
}

let edgeU = new Set();
let edgeD = new Set();
let edgeL = new Set();
let edgeR = new Set();
let cornerUL = 0;
let cornerUR = input[0].length - 1;
let cornerDL = input[0].length * (input[0].length - 1);
let cornerDR = input[0].length * input[0].length - 1;
for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[0].length; j++) {
    if (j == 0 && (i != 0 || i != input.length - 1)) {
      edgeL.add(j + i * input[0].length);
    }
    if (j == input[0].length - 1 && (i != 0 || i != input.length - 1)) {
      edgeR.add(j + i * input[0].length);
    }
    if (i == 0 && (j != 0 || i != input[0].length - 1)) {
      edgeU.add(j + i * input[0].length);
    }
    if (i == input.length - 1 && (j != 0 || i != input[0].length - 1)) {
      edgeD.add(j + i * input[0].length);
    }
  }
}

function findVoisins(input, a) {
  let voisins = new Set();
  if (a == cornerUL) {
    voisins.add(1);
    voisins.add(input[0].length);
  } else if (a == cornerUR) {
    voisins.add(input[0].length - 2);
    voisins.add(2 * input[0].length - 1);
  } else if (a == cornerDL) {
    voisins.add(input[0].length * (input[0].length - 2));
    voisins.add(input[0].length * (input[0].length - 1) + 1);
  } else if (a == cornerDR) {
    voisins.add(input[0].length * (input[0].length - 1) - 1);
    voisins.add(input[0].length * input[0].length - 2);
  } else if (edgeD.has(a)) {
    voisins.add(a - 1);
    voisins.add(a + 1);
    voisins.add(a - input[0].length);
  } else if (edgeU.has(a)) {
    voisins.add(a - 1);
    voisins.add(a + 1);
    voisins.add(a + input[0].length);
  } else if (edgeR.has(a)) {
    voisins.add(a - input[0].length);
    voisins.add(a + input[0].length);
    voisins.add(a - 1);
  } else if (edgeL.has(a)) {
    voisins.add(a - input[0].length);
    voisins.add(a + input[0].length);
    voisins.add(a + 1);
  } else {
    voisins.add(a + 1);
    voisins.add(a - 1);
    voisins.add(a - input[0].length);
    voisins.add(a + input[0].length);
  }
  return voisins;
}

let part1 = dijkstra(input);

// Part 2

const dijkstra2 = (input) => {
  let V = new Set();
  let N = input[0].length * input.length;
  let d = [];
  let pred = [];
  for (let k = 0; k < N; k++) {
    V.add(k);
    d[k] = Inf;
    pred[k] = 0;
  }
  let eCoord = posInit(input, "E");
  let e = eCoord[0] * input[0].length + eCoord[1];
  // Part 2
  s = e;

  d[s] = 0;
  let found = false;

  while (V.size != 0 && !found) {
    let a = minD(V, d);
    V.delete(a);
    // part 2
    if (input[giveCoord(input, a)[0]][giveCoord(input, a)[1]] == "a") {
      found = true;
      e = a;
    } else {
      let voisins = findVoisins(input, a);
      let intersect = new Set([...V].filter((i) => voisins.has(i)));
      for (let b of intersect) {
        // Part 2 exchange a,b
        if (d[b] > d[a] + poids(b, a, input)) {
          d[b] = d[a] + poids(b, a, input);
          pred[b] = a;
        }
      }
    }
  }
  let chemin = [];
  knot = e;
  while (knot != s) {
    chemin.unshift(knot);
    knot = pred[knot];
  }
  chemin.unshift(s);
  return chemin.length - 1;
};

part2 = dijkstra2(input);

let end = performance.now();
console.log("Execution time :", (end - start).toFixed(2), "ms");

console.log("Part 1 :", part1);
console.log("Part 2 :", part2);
