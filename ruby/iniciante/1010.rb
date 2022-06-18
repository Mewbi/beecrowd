num1 = gets.chomp.split(' ')
p1 = num1[1].to_i
v1 = num1[2].to_f
num2 = gets.chomp.split(' ')
p2 = num2[1].to_i
v2 = num2[2].to_f
puts "VALOR A PAGAR: R$ "'%.2f' % "#{(p1 * v1) + (p2 * v2)}"