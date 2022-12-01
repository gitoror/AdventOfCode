const fs = require('fs');
const input = fs.readFileSync('inputs/day8/input', 'utf-8').trim().split('\n');
const ex = fs.readFileSync('inputs/day8/ex', 'utf-8').trim().split('\n');

// Part 1
const accBeforeLoop = (input) => {
    i = 0;
    readLines = [];
    acc = 0;
    while ((i < input.length) & !readLines.includes(i)) {
        line = input[i].split(' ');
        readLines.push(i);
        if (line[0] == 'acc') {
            op = line[1][0];
            nb = Number(line[1].slice(1));
            if (op == '+') {
                acc += nb;
            }
            if (op == '-') {
                acc -= nb;
            }
            i++;
        }
        if (line[0] == 'nop') {
            i++;
        }
        if (line[0] == 'jmp') {
            op = line[1][0];
            nb = Number(line[1].slice(1));
            if (op == '+') {
                i += nb;
            }
            if (op == '-') {
                i -= nb;
            }
        }
    }
    return acc;
};

console.log(accBeforeLoop(input));

// Part 2
const accCorrectProgram = (input) => {
    listNopJmp = [];
    input.forEach((line, i) => {
        if (['nop', 'jmp'].includes(line.split(' ')[0])) {
            listNopJmp.push(i);
        }
    });
    while (listNopJmp.length > 0) {
        p = listNopJmp.pop();
        changeNopJmp(input, p);
        i = 0;
        readLines = [];
        acc = 0;
        while ((i < input.length) & !readLines.includes(i)) {
            line = input[i].split(' ');
            readLines.push(i);
            if (line[0] == 'acc') {
                op = line[1][0];
                nb = Number(line[1].slice(1));
                if (op == '+') {
                    acc += nb;
                }
                if (op == '-') {
                    acc -= nb;
                }
                i++;
            }
            if (line[0] == 'nop') {
                i++;
            }
            if (line[0] == 'jmp') {
                op = line[1][0];
                nb = Number(line[1].slice(1));
                if (op == '+') {
                    i += nb;
                }
                if (op == '-') {
                    i -= nb;
                }
            }
        }
        if (i >= input.length) {
            return acc;
        }
        changeNopJmp(input, p);
    }
    return 'NoCorrectProgram';
};

const changeNopJmp = (input, index) => {
    line = input[index];
    cmd = line.split(' ')[0];
    if (cmd == 'nop') {
        input[index] = line.replace('nop', 'jmp');
    }
    if (cmd == 'jmp') {
        input[index] = line.replace('jmp', 'nop');
    }
};

console.log(accCorrectProgram(input));
