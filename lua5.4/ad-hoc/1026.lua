while true do
    local x, y = io.read("*n", "*n")
    if x == nil or y == nil then
        break
    end
    print(x ~ y)
end
