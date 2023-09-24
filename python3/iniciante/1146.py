while True:
    x = int(input())
    if x == 0:
        break

    list_str = ""
    for i in range(1, x + 1):
        list_str += str(i)
        if i < x:
            list_str += " "

    print(list_str)
