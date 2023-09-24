n = float(input())

adjusts = [0.15, 0.12, 0.10, 0.07, 0.04]
percent = ""
ajust = 0

if n <= 400.00:
    ajust = n * adjusts[0]
    n += ajust
    percent = "15"
elif 400.01 <= n <= 800.00:
    ajust = n * adjusts[1]
    n += ajust
    percent = "12"
elif 800.01 <= n <= 1200.00:
    ajust = n * adjusts[2]
    n += ajust
    percent = "10"
elif 1200.01 <= n <= 2000.00:
    ajust = n * adjusts[3]
    n += ajust
    percent = "7"
elif n > 2000.00:
    ajust = n * adjusts[4]
    n += ajust
    percent = "4"

print("Novo salario: {:.2f}".format(n))
print("Reajuste ganho: {:.2f}".format(ajust))
print("Em percentual: {} %".format(percent))
