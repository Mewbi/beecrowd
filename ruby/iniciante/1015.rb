p1 = gets.chomp.split(' ')
p2 = gets.chomp.split(' ')
x1 = p1[0].to_f
y1 = p1[1].to_f
x2 = p2[0].to_f
y2 = p2[1].to_f
puts '%.4f' % "#{Math.sqrt(((x2 - x1) ** 2)  + ((y2 - y1) ** 2))}"