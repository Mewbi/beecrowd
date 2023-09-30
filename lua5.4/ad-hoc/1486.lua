while true do
    local n, m, p = io.read("*n", "*n", "*n")
    if n == 0 and m == 0 and p == 0 then
        break
    end

    local vetor = {}
    for i = 1, m do
        vetor[i] = {}
        for j = 1, n do
            vetor[i][j] = tonumber(io.read("*n"))
        end
    end

    local resp = 0
    for i = 1, n do
        local jaforam = 0
        for j = 1, m do
            if vetor[j][i] == 1 then
                jaforam = jaforam + 1
            else
                if jaforam >= p then
                    resp = resp + 1
                end
                jaforam = 0
            end
        end
        if jaforam >= p then
            resp = resp + 1
        end
    end

    print(resp)
end
