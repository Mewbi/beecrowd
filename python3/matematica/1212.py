def NumCarryAdd(n1, n2):
    c = 0
    t = 0
    while True:
        a = n1 % 10
        b = n2 % 10
        n1 = n1 // 10
        n2 = n2 // 10
        if (a + b + c) >= 10:
            t += 1
            c = 1
        else:
            c = 0
        if n1 == 0 and n2 == 0:
            break
    return t

while True:
    x, y = map(int, input().split())
    if x == 0 and y == 0:
        break
    carry = NumCarryAdd(x, y)
    if carry == 0:
        print("No carry operation.")
    elif carry == 1:
        print("1 carry operation.")
    else:
        print("{} carry operations.".format(carry))
