#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Thu Jan 21 16:15:56 2021

@author: ekudgan
"""

dom = {'A':11, 'K':4, 'Q':3, 'J':20, 'T':10, '9':14, '8': 0, '7': 0}
not_dom = {'A':11, 'K':4, 'Q':3, 'J':2, 'T':10, '9':0, '8': 0, '7': 0}
total_score = 0

firstline = input()
hands = int(firstline.split()[0])
dom_suit = firstline.split()[1]

for i in range(4*hands):
    content = list(map(str, input())) 
    # print(content)
    if content[1] == dom_suit:
        score = dom[content[0]]
    else:
        score = not_dom[content[0]]
    
    total_score = total_score + score

print(total_score)