function nonzero(n)
    if n == 0 or n == 1 then
        return 1
    end

    local x = math.floor(n / 5)
    local y = n % 5

    local potX = x % 4
    if potX == 0 then
        potX = 1
    elseif potX == 1 then
        potX = 2
    elseif potX == 2 then
        potX = 4
    elseif potX == 3 then
        potX = 8
    end

    local lastY = 0
    if y == 0 then
        lastY = 1
    elseif y == 1 then
        lastY = 1
    elseif y == 2 then
        lastY = 2
    elseif y == 3 then
        lastY = 6
    elseif y == 4 then
        lastY = 4
    end

    return potX * nonzero(x) * lastY
end

local instancia = 1
while true do
    local n = io.read("*n")
    if not n then
        break
    end
    print("Instancia " .. instancia)
    instancia = instancia + 1

    local result = nonzero(n)
    while result >= 10 do
        result = result % 10
    end

    print(result)
    print()
end
