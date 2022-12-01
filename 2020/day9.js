const fs = require('fs');
const input = fs.readFileSync('inputs/day9/input', 'utf-8').trim().split('\n').map(Number);
const ex = fs.readFileSync('inputs/day9/ex', 'utf-8').trim().split('\n').map(Number);

// Part 1
const noSumPrevious = (input, len) => {
    i = len;
    while (i < input.length) {
        foundSum = false;
        goalVal = input[i];
        numbers = input.slice(i - len, i);
        numbers.forEach((nb1) => {
            numbers.forEach((nb2) => {
                if ((nb1 + nb2 == goalVal) & (nb1 != nb2)) {
                    foundSum = true;
                }
            });
        });
        if (foundSum == false) {
            return goalVal;
        }
        i++;
    }
    return 'Did nit find !';
};

console.log('Part 1', noSumPrevious(input, 25));
//console.log(noSumPrevious(ex, 5));

// Part 2
const sumContiguousExt = (input, len_) => {
    goalVal = noSumPrevious(input, len_);
    for (let len = 2; len < input.length - 1; len++) {
        for (let i = 0; i < input.length - len; i++) {
            numbers = input.slice(i, i + len);
            if (len == 4) {
            }
            if (numbers.reduce((acc, v) => acc + v) == goalVal) {
                numbers.sort((a, b) => a - b);
                return numbers[0] + numbers.at(-1);
            }
        }
    }
    return 'Did not find';
};

console.log('Part 2', sumContiguousExt(input, 25));
