import math

def bhaskara(a, b, c):
	x  = (b ** 2) - (4 * a * c)
	x  = math.sqrt(x)
	x1 = (-b + x) / (2 * a)
	x2 = (-b - x) / (2 * a)

	xr = x1 if x1 > 0 else x2

	return xr

def calc(h, p1, p2, a, v, pi, g):
	# vx = v * cos ang
	# vy = v * sin ang
	vy = v * math.sin(a * pi / 180)
	vx = v * math.cos(a * pi / 180)

	# 0 = h + (vy * t) + ((g / 2) * t ** 2) # Bhaskara
	t = bhaskara(-g / 2, vy, h)

	# Alcance = vx . t
	A = vx * t

	# Verifica se acertou caraio
	if (A >= p1) and (A <= p2):
		print('{:.5f} -> DUCK'.format(A))
	else: 
		print('{:.5f} -> NUCK'.format(A))

while True:
	try:
		# Altura inicial
		h = float(input())
		# Inicio e fim do alvo
		p = [int(n) for n in input().split()]
		# Tentativas
		n = int(input())

		for i in range(n):
			V = [float(x) for x in input().split()]

			calc(h, p[0], p[1], V[0], V[1], 3.14159, 9.80665)
	except EOFError:
		break