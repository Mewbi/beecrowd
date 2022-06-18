num = gets.chomp.split(' ')
A = num[0].to_f
B = num[1].to_f
C = num[2].to_f
pi = 3.14159
puts "TRIANGULO: "'%.3f' % "#{(A * C) / 2}"
puts "CIRCULO: "'%.3f' % "#{(pi) * (C ** 2)}"
puts "TRAPEZIO: "'%.3f' % "#{((A + B) * C) / 2}"
puts "QUADRADO: "'%.3f' % "#{B * B}"
puts "RETANGULO: "'%.3f' % "#{A * B}"