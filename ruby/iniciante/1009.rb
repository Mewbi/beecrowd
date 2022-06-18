nome = gets.chomp
salario = gets.chomp.to_f
vendas = gets.chomp.to_f
puts "TOTAL = R$ "'%.2f' % "#{salario + (vendas * 0.15)}"