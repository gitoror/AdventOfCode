f = 'day6/input';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n\n');

groupPoints = (group) => {
    group = group.replaceAll('\n', '').split('');
    answerLetter = [];
    group.forEach((letter) => {
        if (!answerLetter.includes(letter)) {
            answerLetter.unshift(letter);
        }
    });
    return answerLetter.length;
};

ex = 'abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb'.split('\n\n');

console.log(
    input.reduce((count, currentGroup) => {
        return count + groupPoints(currentGroup);
    }, 0)
);

// Part 2

groupPoints2 = (group) => {
    individuals = group.split('\n');
    letters = new Set();
    individuals.forEach((individual, index) => {
        answerLetter = new Set();
        individual = individual.split('');
        individual.forEach((letter) => {
            answerLetter.add(letter);
        });
        if (index == 0) {
            letters = answerLetter;
        }
        letters = new Set([...letters].filter((i) => answerLetter.has(i)));
    });
    return letters.size;
};

console.log(
    input.reduce((count, currentGroup) => {
        return count + groupPoints2(currentGroup);
    }, 0)
);