def bhaskara(a, b, c)
  x = Math.sqrt(b**2 - 4 * a * c)
  x1 = (-b + x) / (2 * a)
  x2 = (-b - x) / (2 * a)

  x1 > 0 ? x1 : x2
end

def calc(h, p1, p2, a, v)
  pi = 3.14159
  g = 9.80665

  vy = v * Math.sin(a * pi / 180)
  vx = v * Math.cos(a * pi / 180)

  t = bhaskara(-g / 2, vy, h)

  alc = vx * t

  res = alc >= p1 && alc <= p2 ? "DUCK" : "NUCK"

  puts "%.5f -> %s" % [alc, res]
end

loop do
  input = gets
  break unless input

  h = input.chomp.to_f
  p1, p2 = gets.split.map(&:to_f)
  n = gets.chomp.to_i

  n.to_i.times do
    input = gets
    a, v = input.split.map(&:to_f)
    calc(h, p1, p2, a, v)
  end
end
