defmodule M do
    def kamenetsky(0) do
        0
    end
    def kamenetsky(1) do
        1
    end
    def kamenetsky(n) do
      x = (n * :math.log10(n / :math.exp(1.0)) +
            ( :math.log10(2.0 * :math.pi * n) / 2.0) )
      :math.floor(x) + 1
    end
    
    def main do
        {n,_} = IO.gets("") |> String.trim |> Float.parse
        result = kamenetsky(n)
        IO.puts trunc(result)
    end
end

M.main()
