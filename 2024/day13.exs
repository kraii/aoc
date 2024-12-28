defmodule Day13 do
  defp lines() do
    File.stream!("test/day13.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.reject(&(&1 == ""))
  end

  defp parse([], _, acc), do: acc

  defp parse(lines, prize_ext, acc) do
    {[a_line, b_line, prize_line], rest} = Enum.split(lines, 3)

    {px, py} = get_x_y(prize_line)
    crane = %{a: get_x_y(a_line), b: get_x_y(b_line), prize: {px + prize_ext, py + prize_ext}}
    parse(rest, prize_ext, [crane | acc])
  end

  defp get_x_y(line) do
    [_, x, y] = Regex.run(~r/X.([0-9]+), Y.([0-9]+)/, line)
    {String.to_integer(x), String.to_integer(y)}
  end

  def solve(prize_ext) do
    cranes = parse(lines(), prize_ext, [])

    for crane <- cranes, reduce: 0 do
      acc ->
        acc + cost(crane)
    end
  end

  defp cost(%{:a => {ax, ay}, :b => {bx, by}, :prize => {px, py}}) do

    n = (px * by - py * bx) / (ax * by - bx * ay)
    m = (px - ax * n) / bx

    if whole(n) && whole(m) do
      trunc(n * 3 + m)
    else
      0
    end
  end

  defp whole(n), do: trunc(n) == n
end

IO.puts("part 1")
Day13.solve(0) |> IO.inspect()
IO.puts("part 2")
Day13.solve(10000000000000) |> IO.inspect()
