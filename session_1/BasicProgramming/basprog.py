if __name__ == "__main__":
    
    line_1 = input()
    array_len = int(line_1.split()[0])
    ops = line_1.split()[1]
    
    input_list = sorted(input().split())
    
    if ops == '1':
        flag = 0        
        int_list = list(map(int, input_list))
        filter_list = [x for x in int_list if x <= 7777]
        
        if filter_list:
            for num in filter_list: 
                if 7777-num in filter_list:
                    print("Yes")
                    flag = 1
                    break
            if flag == 0:
                print("No")
        else:
            print("No")
    
    if ops == '2':
        set1 = set(input_list)
        if sorted(list(set1)) == input_list:
            print("Unique")
        else:
            print("Contains duplicate")
            
    if ops == '3':
        myset = sorted(list(set(input_list)))
        if myset == input_list:
            print("-1")
        else:
            for element in myset:
                if input_list.count(element) > array_len/2:
                    print(element)
                    break
            print("-1")
    
    if ops == '4':
        if array_len%2 == 0:
            median1 = input_list[array_len//2 - 1]
            median2 = input_list[array_len//2]
            print(median1,median2)
        else:
            median = input_list[array_len//2]
            print(median)
                
    if ops == '5':
        int_list = filter(lambda x: int(x)>99 and int(x)<1000, input_list)
        print(*int_list, sep=' ')