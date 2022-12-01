f = 'day5/input';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n');

calcRow = (code) => {
    code = code.slice(0, 7);
    return code.split('').reduce((row, currentValue, index) => {
        if (currentValue == 'B') {
            return row + 2 ** (6 - index);
        }
        return row;
    }, 0);
};

calcCol = (code) => {
    code = code.slice(7, 10);
    return code.split('').reduce((col, currentValue, index) => {
        if (currentValue == 'R') {
            return col + 2 ** (2 - index);
        }
        return col;
    }, 0);
};

calcId = (code) => {
    return calcCol(code) + 8 * calcRow(code);
};

console.log(
    input.map(calcId).reduce((a, b) => {
        return Math.max(a, b);
    })
);

// Part 2
hideId = (ids) => {
    ids = ids.filter((id) => !ids.includes(id + 1));
    return Math.min.apply(null, ids) + 1;
};
console.log(hideId(input.map(calcId)));

S = new Set();
S.add('a');
Q = new Set();
Q.add('b');
Q.add('a');
Z = new Set([...S].filter((i) => Q.has(i)));
console.log([...S], Q, Q.has('a'));
