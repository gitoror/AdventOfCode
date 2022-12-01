f = 'day7/input.txt';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n');

// Part 1
const hasGoldDirect = (line) => {
    container = line.split(' bags contain')[0];
    contained = line.split(' bags contain ')[1].split(', ');
    return (
        contained.filter((el) => {
            bags = el.split(' ');
            if (`${bags[1]} ` + `${bags[2]}` == 'shiny gold') {
                return true;
            }
        }).length != 0
    );
};

countTotBags = (input) => {
    goldBags = input.filter(hasGoldDirect).map((line) => line.split(' bags contain')[0]);
    a = input.filter(hasGoldDirect).length;
    otherBags = [];
    size = 0;
    prevsize = -1;
    while (size != prevsize) {
        prevsize = size;
        otherBags = input
            .filter((line) => {
                container = line.split(' bags contain')[0];
                contained = line.split(' bags contain ')[1].split(', ');
                return (
                    contained.filter((el) => {
                        bags = el.split(' ');
                        if (goldBags.includes(`${bags[1]} ` + `${bags[2]}`) & !goldBags.includes(container)) {
                            return true;
                        }
                        if (otherBags.includes(`${bags[1]} ` + `${bags[2]}`) & !goldBags.includes(container)) {
                            return true;
                        }
                        return false;
                    }).length != 0
                );
            })
            .map((line) => line.split(' bags contain')[0]);
        size = otherBags.length;
    }
    b = otherBags.length;
    return a + b;
};

ex = fs.readFileSync(`./inputs/day7/ex`, 'utf-8').trim().split('\n');
console.log('Result Part 1 :', countTotBags(input));

// Part 2 - recursive
const countTotContainedBags = (bag, input) => {
    bagContent = input.filter((line) => {
        if (line.split(' bags contain')[0] == bag) {
            return true;
        }
    })[0];
    bagContent = bagContent.split(' bags contain ')[1].split(', ');
    if (bagContent == 'no other bags.') {
        return 0;
    }
    count = 0;
    bagContent.forEach((bag) => {
        bag = bag.split(' ');
        bagName = `${bag[1]} ` + `${bag[2]}`;
        bagCoef = Number(bag[0]);
        count += bagCoef + bagCoef * countTotContainedBags(bagName, input);
    });
    return count;
};

ex2 = fs.readFileSync(`./inputs/day7/ex2`, 'utf-8').trim().split('\n');
console.log('Result Part 2 :', countTotContainedBags('shiny gold', input));

// Part 1 - recursive, but too long
/*
const containsGoldBag = (line, input) => {
    bagContent = line.split(' bags contain ')[1].split(', ');
    if (bagContent == 'no other bags.') {
        return 0;
    }
    for (let bag of bagContent) {
        bag = bag.split(' ');
        bagName = `${bag[1]} ` + `${bag[2]}`;
        if (bagName == 'shiny gold') {
            return 1;
        } else {
            bagLine = input.filter((line) => {
                if (line.split(' bags contain')[0] == bagName) {
                    return true;
                }
            })[0];
            count = containsGoldBag(bagLine, input);
            if (count == 1) {
                return 1;
            }
        }
    }
    return 0;
};

const countContainGoldBag = (input) => {
    c = 0;
    input.forEach((line) => {
        c += containsGoldBag(line, input);
    });
    return c;
};
console.log('Result Part 1 :', countContainGoldBag(input));
*/

L1 = [1, 2, 3];
L2 = L1;
//L2[0] = 9;
const change = (arr, i) => {
    arr[i] = 5;
};
change(L1, 0);
console.log(L1);
