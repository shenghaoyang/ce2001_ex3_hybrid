#!/usr/bin/env python3

import pathlib
import sys
import re

if len(sys.argv) < 2:
    print('No folder specified.')   
    exit(1) 

folder = pathlib.Path(sys.argv[1])
folder.resolve()

pattern_sz = re.compile(r'input size (\d+)')
pattern_threshold = re.compile(r'threshold (\d+) [(](\d+) loops')
pattern_keycomps = re.compile(r'comparisons: (\d+)')
pattern_time = re.compile(r'Average time: (\S+)')
pattern_stddev = re.compile(r'Standard deviation: (\S+)')

print('size, threshold, loops, keycomps, time, stddev')

for f in folder.iterdir():
    name = f.name
    content = f.read_text()
    sz = int(pattern_sz.search(content)[1])
    threshold = int(pattern_threshold.search(content)[1])
    loops = int(pattern_threshold.search(content)[2])
    keycomps = int(pattern_keycomps.search(content)[1])
    time = float(pattern_time.search(content)[1])
    stddev = float(pattern_stddev.search(content)[1])
    
    print(sz, threshold, loops, keycomps, time, stddev, sep=', ')


