f = 'day4/input.txt';
const fs = require('fs');
const input = fs.readFileSync(`./inputs/${f}`, 'utf-8').trim().split('\n\n');

validDoc = (doc, index) => {
    doc = doc.replaceAll('\n', ' ');
    fields = doc.split(' ');
    fieldsReq = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']; // no cid bcs optional
    nbFields = fields.filter((field) => {
        fieldName = field.split(':')[0];
        if (fieldsReq.includes(fieldName)) {
            fieldsReq.filter((field) => field != fieldName);
        }
        return fieldsReq.includes(fieldName);
    }).length;
    return nbFields == 7;
};

ex = ['ecl:gry\npid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm'];
ex = [ex[0].replace('\n', ' ')];

console.log(input.filter(validDoc).length);

// Part 2
validDoc2 = (doc, index) => {
    doc = doc.replaceAll('\n', ' ');
    fields = doc.split(' ');
    fieldsReq = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']; // no cid bcs optional
    correctFields = fields.filter((field) => {
        fieldName = field.split(':')[0];
        fieldContent = field.split(':')[1];
        if (fieldsReq.includes(fieldName)) {
            fieldsReq.filter((field) => field != fieldName); // Be sure that all fields don't appear sevral times
            if (fieldName == 'byr') {
                if ((1920 <= Number(fieldContent)) & (Number(fieldContent) <= 2002)) {
                    return true;
                }
            }
            if (fieldName == 'iyr') {
                if ((2010 <= Number(fieldContent)) & (Number(fieldContent) <= 2020)) {
                    return true;
                }
            }
            if (fieldName == 'eyr') {
                if ((2020 <= Number(fieldContent)) & (Number(fieldContent) <= 2030)) {
                    return true;
                }
            }
            if (fieldName == 'hgt') {
                content = fieldContent.split('cm');
                if (content.length == 2) {
                    if ((150 <= Number(content[0])) & (Number(content[0]) <= 193)) {
                        return true;
                    }
                } else {
                    content = fieldContent.split('in');
                    if ((59 <= Number(content[0])) & (Number(content[0]) <= 76)) {
                        return true;
                    }
                }
            }
            if (fieldName == 'hcl') {
                if (fieldContent[0] == '#') {
                    content = fieldContent.split('#')[1];
                    if (content.length == 6) {
                        if (
                            content
                                .split('')
                                .filter((x) =>
                                    [
                                        '0',
                                        '1',
                                        '2',
                                        '3',
                                        '4',
                                        '5',
                                        '6',
                                        '7',
                                        '8',
                                        '9',
                                        'a',
                                        'b',
                                        'c',
                                        'd',
                                        'e',
                                        'f',
                                    ].includes(x)
                                ).length == 6
                        ) {
                            return true;
                        }
                    }
                }
            }
            if (fieldName == 'ecl') {
                if (['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].includes(fieldContent)) {
                    return true;
                }
            }
            if (fieldName == 'pid') {
                if (fieldContent.length == 9) {
                    return true;
                }
            }
        }
        return false;
    });
    nbFields = correctFields.length;
    return nbFields == 7;
};

console.log(input.filter(validDoc2).length);
