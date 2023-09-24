d = int(input())

a = d // 365
d = d % 365

m = d // 30
d = d % 30

print("{} ano(s)".format(a))
print("{} mes(es)".format(m))
print("{} dia(s)".format(d))
