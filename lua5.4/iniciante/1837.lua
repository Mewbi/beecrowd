function euclideanDivision(a, b)
    local limit = b < 0 and -b or b
    local r, q = 0, 0

    for r = 0, limit - 1 do
        if (a - r) % b == 0 then
            q = (a - r) // b
            return q, r
        end
    end

    return q, r
end

local a, b = io.read("*n", "*n")
local q, r = euclideanDivision(a, b)
print(q .. " " .. r)
