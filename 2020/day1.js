f = 'day1/input.txt';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n').map(Number);
// Part 1
var p = 0;
let findP1 = input.forEach((el1) => {
    input.forEach((el2) => {
        if (el1 + el2 == 2020) p = el1 * el2;
    });
});

function findP2(input) {
    var e = input[0];
    var p;
    var tab = input.slice(1);
    tab.forEach((el) => {
        if (e + el == 2020) p = e * el;
    });
    if (p != undefined) return p;
    return findP2(input.slice(1));
}

console.log(findP2(input));

// Part 2
var p = 0;
let findP3 = input.forEach((el1) => {
    input.forEach((el2) => {
        input.forEach((el3) => {
            if (el1 + el2 + el3 == 2020) p = el1 * el2 * el3;
        });
    });
});
console.log(p);
