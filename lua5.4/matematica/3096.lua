local function kamenetsky(n)
    if n < 0 then
        return 0
    end

    if n <= 1 then
        return 1
    end

    local x = ((n * math.log10(n/math.exp(1))) +
                (math.log10(2*math.pi*n) / 2.0))

    return math.floor(x) + 1
end

local n = io.read("*n")
print(kamenetsky(n))
