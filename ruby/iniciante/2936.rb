curupira = gets.chomp.to_i
boitata = gets.chomp.to_i
boto = gets.chomp.to_i
mapinguari = gets.chomp.to_i
iara = gets.chomp.to_i
chica = 225

curupira = curupira * 300
boitata = boitata * 1500
boto = boto * 600
mapinguari = mapinguari * 1000
iara = iara * 150

total = curupira + boitata + boto + mapinguari + iara + chica
puts "#{total}"