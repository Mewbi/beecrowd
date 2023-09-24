n = int(input())
init = 1

for i in range(n):
    for j in range(init, init + 3):
        print(j, end=" ")
    print("PUM")
    init += 4
