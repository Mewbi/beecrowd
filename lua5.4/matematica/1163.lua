function bhaskara(a, b, c)
    local x = math.sqrt(math.pow(b, 2) - (4 * a * c))
    local x1 = (-b + x) / (2 * a)
    local x2 = (-b - x) / (2 * a)

    if x1 > 0 then
        return x1
    end

    return x2
end

function calc(h, p1, p2, a, v)
    local pi = 3.14159
    local g = 9.80665

    local vy = v * math.sin(a * pi / 180)
    local vx = v * math.cos(a * pi / 180)

    local t = bhaskara(-g / 2, vy, h)

    local alc = vx * t

    local res = "NUCK"
    if alc >= p1 and alc <= p2 then
        res = "DUCK"
    end

    print(string.format("%.5f -> %s", alc, res))
end

while true do
    local h = io.read("*n")
    if h == nil then
        break
    end

    local p1, p2 = io.read("*n", "*n")
    local n = io.read("*n")

    for i = 1, n do
        local a, v = io.read("*n", "*n")
        calc(h, p1, p2, a, v)
    end
end

