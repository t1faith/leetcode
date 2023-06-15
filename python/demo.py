
list = [1,2,3,4]
target = 5

for i in range(len(list)):
    x = list[i]
    for j in range(len(list)):
        y = list[j]
        s = x + y
        if s == target:
            # print("x",x,"y",y,"s",s)
            print("output:",i,j)
