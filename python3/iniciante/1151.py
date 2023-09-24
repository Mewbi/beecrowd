def fibonacci(n):
    a, b = 0, 1
    for i in range(n):
        a, b = b, a + b
    return a

n = int(input())
output = ""

for i in range(n):
    output += str(fibonacci(i))

    if i < n - 1:
        output += " "

print(output)
