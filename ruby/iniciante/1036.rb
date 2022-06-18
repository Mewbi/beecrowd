coef = gets.chomp.split(' ')
A = coef[0].to_f
B = coef[1].to_f
C = coef[2].to_f
Delta = (B ** 2) + (-4 * A * C)

if Delta <= 0 || A == 0
	puts "Impossivel calcular"
else
	R1 = (- B + Math.sqrt(Delta)) / (2 * A)
	R2 = (- B - Math.sqrt(Delta)) / (2 * A)
	puts "R1 = "'%.5f' % R1
	puts "R2 = "'%.5f' % R2
end