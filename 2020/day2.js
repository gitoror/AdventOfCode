f = 'day2/input';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n');

// Part 1
const validPwd1 = (line) => {
    s = line.split(' ');
    min = s[0].split('-')[0];
    max = s[0].split('-')[1];
    letter = s[1][0].split('')[0];
    word = s[2].split('');
    countLett = word.reduce((count, currentLetter) => {
        if (currentLetter == letter) {
            return count + 1;
        }
        return count;
    }, 0);
    return (min <= countLett) & (countLett <= max);
};

console.log(input.filter(validPwd1).length);

// Part2
const validPwd2 = (line) => {
    s = line.split(' ');
    pos1 = s[0].split('-')[0];
    pos2 = s[0].split('-')[1];
    letter = s[1][0].split('')[0];
    word = s[2].split('');
    if ((word[`${pos1}` - 1] == letter) ^ (word[`${pos2}` - 1] == letter)) {
        return true;
    }
    return false;
};

console.log(input.filter(validPwd2).length);
