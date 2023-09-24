n, sum, count = 0, 0, 0

while True:
    n = int(input())
    
    if n >= 0:
        count += 1
        sum += n
    else:
        break

if count > 0:
    average = sum / count
    print("{:.2f}".format(average))
