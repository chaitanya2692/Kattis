# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""

if __name__ == "__main__":

    base_dict = { 2: '2', 
                  3: '3', 
                  4: '4', 
                  5: '5', 
                  6: '6', 
                  7: '7', 
                  8: '8', 
                  9: '9',
                 10: 'a',
                 11: 'b',
                 12: 'c',
                 13: 'd',
                 14: 'e',
                 15: 'f',
                 16: 'g',
                 17: 'h',
                 18: 'i',
                 19: 'j',
                 20: 'k',
                 21: 'l',
                 22: 'm',
                 23: 'n',
                 24: 'o',
                 25: 'p',
                 26: 'q',
                 27: 'r',
                 28: 's',
                 29: 't',
                 30: 'u',
                 31: 'v',
                 32: 'w',
                 33: 'x',
                 34: 'y',
                 35: 'z',
                 36: '0'}
    

    first_line = int(input())
    
    for inputs in range(first_line):
        bases = []
        sample_input = input()
        myinput = sample_input.split()
        
        a,b,c = list(myinput[0]),list(myinput[2]), list(myinput[4])
    
        set_a, set_b, set_c = set(a), set(b), set(c)
    
        if len(set_a) == 1 and '1' in set_a:
            if len(set_b) == 1 and '1' in set_b:
                if len(set_c) == 1 and '1' in set_c:
                    sum_a = sum(list(map(int, a)))
                    sum_b = sum(list(map(int, b)))
                    sum_c = sum(list(map(int, c)))
                    if myinput[1] == '+':
                        if sum_c == sum_a + sum_b:
                            bases.append('1')
                    elif myinput[1] == '-':
                        if sum_c == sum_a - sum_b:
                            bases.append('1')
                    elif myinput[1] == '*':
                        if sum_c == sum_a * sum_b:
                            bases.append('1')
                    elif myinput[1] == '/':
                        if sum_c == sum_a / sum_b:
                            bases.append('1')
                            
        for i in range(2,37):
            try:
                a,b,c = int(myinput[0],i), int(myinput[2],i), int(myinput[4],i)
                if myinput[1] == '+':
                    if c == a + b:
                        bases.append(base_dict[i])
                elif myinput[1] == '-':
                    if c == a - b:
                        bases.append(base_dict[i])
                elif myinput[1] == '*':
                    if c == a * b:
                        bases.append(base_dict[i])
                elif myinput[1] == '/':
                    if c == a / b:
                        bases.append(base_dict[i])
            except:
                continue
            
        if not bases:
            print("invalid")
        else:
            print("".join(bases))
    