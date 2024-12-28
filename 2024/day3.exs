defmodule Day3 do
  defp file(), do: "test/day3.txt"

  def part1() do
    data = File.read!(file())

    sumMul(data)
  end

  defp sumMul(data) do
    Regex.scan(~r/mul\(([0-9]+),([0-9]+)\)/, data, capture: :all_but_first)
    |> Enum.map(fn [x, y] ->
      String.to_integer(x) * String.to_integer(y)
    end)
    |> Enum.sum()
  end

  def part2() do
    data = File.read!(file())

    matches =
      Regex.scan(~r/mul\(([0-9]+),([0-9]+)\)|(don't\(\))|(do\(\))/, data, capture: :all_but_first)

    sum(matches, true, 0)
  end

  defp sum([], _take?, acc), do: acc

  defp sum([match | tail], true, acc) do
    case match do
      [x, y] when length(match) == 2 ->
        sum(tail, true, acc + String.to_integer(x) * String.to_integer(y))

      ["", "", "", "do()"] ->
        sum(tail, true, acc)

      ["", "", "don't()"] ->
        sum(tail, false, acc)
    end
  end

  defp sum([match | tail], false, acc) do
    case match do
      ["", "", "", "do()"] -> sum(tail, true, acc)
      _ -> sum(tail, false, acc)
    end
  end
end

IO.puts("Part 1 - " <> Integer.to_string(Day3.part1()))

IO.puts("Part 2 - " <> Integer.to_string(Day3.part2()))
