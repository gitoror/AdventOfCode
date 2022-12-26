#!/usr/bin/env python3
import time
import os
import re
import argparse
from sympy.solvers import solve
from sympy.core import sympify
from sympy import Symbol
# import numpy as np


def doArgs(argList, name):
    parser = argparse.ArgumentParser(description=name)

    parser.add_argument('-v', "--verbose", action="store_true",
                        help="Enable verbose debugging", default=False)
    parser.add_argument('--input', action="store", dest="inputFn",
                        type=str, help="Input file name", required=True)
    parser.add_argument('--output', action="store", dest="outputFn",
                        type=str, help="Output file name", required=True)

    return parser.parse_args(argList)


monkeyMap = {}
b = ''

def parseData(filename):
    with open(filename) as f:
        #        lines = f.readlines()
        for l in f.readlines():
            l = l.rstrip()
            ll = l.split(': ')
            partOne = ll[0]
            partTwo = ll[1].split(" ")
            monkeyMap[partOne] = partTwo
            monkeyMap['root'][1] = '='
    print(monkeyMap)
    return




def recurseMonkey(monkeyName):
    if len(monkeyMap[monkeyName]) == 1:
        return int(monkeyMap[monkeyName][0])

    op = monkeyMap[monkeyName][1]

    if op == "+":
        return recurseMonkey(monkeyMap[monkeyName][0]) + recurseMonkey(monkeyMap[monkeyName][2])
    if op == "*":
        return recurseMonkey(monkeyMap[monkeyName][0]) * recurseMonkey(monkeyMap[monkeyName][2])
    if op == "-":
        return recurseMonkey(monkeyMap[monkeyName][0]) - recurseMonkey(monkeyMap[monkeyName][2])
    if op == "/":
        return recurseMonkey(monkeyMap[monkeyName][0]) / recurseMonkey(monkeyMap[monkeyName][2])

    return 0


def recurseMonkeyP2(monkeyName):
    if len(monkeyMap[monkeyName]) == 1:
        return monkeyMap[monkeyName][0]

    op = monkeyMap[monkeyName][1]
    print(b)
    if op == "+":
        b = b + "(" + recurseMonkeyP2(monkeyMap[monkeyName][0]) + \
            "+" + recurseMonkeyP2(monkeyMap[monkeyName][2]) + ")"
    if op == "*":
        b = b + "(" + recurseMonkeyP2(monkeyMap[monkeyName][0]) + \
            "*" + recurseMonkeyP2(monkeyMap[monkeyName][2]) + ")"
    if op == "-":
        b = b + "(" + recurseMonkeyP2(monkeyMap[monkeyName][0]) + \
            "-" + recurseMonkeyP2(monkeyMap[monkeyName][2]) + ")"
    if op == "/":
        b = b + "(" + recurseMonkeyP2(monkeyMap[monkeyName][0]) + \
            "/" + recurseMonkeyP2(monkeyMap[monkeyName][2]) + ")"
    if op == "=":
        b = b + "(" + recurseMonkeyP2(monkeyMap[monkeyName][0]) + \
            "=" + recurseMonkeyP2(monkeyMap[monkeyName][2]) + ")"
    return ''


def main():

    print("Starting")
    startTime = float(time.time())

    parseData("../inputs/day21/ex.txt")
    b = " "
    a = recurseMonkeyP2('root')
    print(a)
    print(b)
    hum = Symbol('humn')
    eq = "hum * 3 - 1 = 4 * 3 + 2"
    sympyEq = sympify("Eq(" + eq.replace("=", ",") + ")")
    sol = solve(sympyEq)
    print(sol)
    print("Finished in %0.4f seconds" % (time.time() - startTime))
    return


if __name__ == '__main__':
    # sys.argv = ["programName.py","--input","test.txt","--output","tmp/test.txt"]
    main()
