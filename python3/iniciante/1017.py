def qnt(t, v):
	d = t * v
	return d / 12

t = int(input())
v = int(input())

print('{:.3f}'.format(qnt(t, v)))