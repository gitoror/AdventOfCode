f = 'day3/input';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n');
// Part 1
isTree = (line, indexLine) => {
    if (line.split('')[(3 * indexLine) % line.length] == '#') {
        return true;
    }
    return false;
};

console.log(input.filter(isTree).length);

// Part 2
countTrees = (input, right, down) => {
    let count = 0;
    lenLine = input[0].length;
    for (let i = 0; i < input.length; i += down) {
        if (input[i].split('')[((right * i) / down) % lenLine] == '#') {
            count++;
        }
    }
    return count;
};
console.log(
    countTrees(input, 1, 1) *
        countTrees(input, 3, 1) *
        countTrees(input, 5, 1) *
        countTrees(input, 7, 1) *
        countTrees(input, 1, 2)
);
/*
1 1
3 1
5 1 
7 1
1 2

*/