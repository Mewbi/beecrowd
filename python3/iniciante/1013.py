def M(a, b, c):
	m = (a + b + abs(a - b)) / 2
	m = (m + c + abs(m - c)) / 2
	return m

n = input().split(' ')

# Funciona no py 3.8, mas não no py 3.4
# print(f'{M(int(n[0]), int(n[1]), int(n[2])):.0f} eh o maior')

# Funciona no py 3.4
print('{:.0f} eh o maior'.format(M(int(n[0]), int(n[1]), int(n[2]))))