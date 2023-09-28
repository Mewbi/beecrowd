while true do
    local n = io.read("*n")
    if not n then
        break
    end

    local out = "vai ter copa!"
    if n > 0 then
        out = "vai ter duas!"
    end
    print(out)
end

