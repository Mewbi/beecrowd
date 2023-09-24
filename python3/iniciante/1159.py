while True:
    x = int(input())

    if x == 0:
        break

    sum_val, count = 0, 0

    while True:
        if x % 2 == 0:
            count += 1
            sum_val += x

        x += 1

        if count == 5:
            break

    print(sum_val)
