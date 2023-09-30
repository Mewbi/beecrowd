function NumCarryAdd(n1, n2)
    local c = 0
    local t = 0
    repeat
        local a = n1 % 10
        local b = n2 % 10
        n1 = math.floor(n1 / 10)
        n2 = math.floor(n2 / 10)
        if (a + b + c) >= 10 then
            t = t + 1
            c = 1
        else
            c = 0
        end
    until n1 == 0 and n2 == 0
    return t
end

while true do
    local x, y = io.read("*n", "*n")
    if x == 0 and y == 0 then
        break
    end
    local carry = NumCarryAdd(x, y)
    if carry == 0 then
        print("No carry operation.")
    elseif carry == 1 then
        print("1 carry operation.")
    else
        print(string.format("%d carry operations.", carry))
    end
end
