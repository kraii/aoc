defmodule Day2 do
  defp file(), do: "test/day2.txt"

  defp parse_lines() do
    File.stream!(file())
    |> Stream.map(&String.trim/1)
    |> Stream.map(fn line ->
      Enum.map(String.split(line), fn c -> String.to_integer(c) end)
    end)
    |> Enum.to_list()
  end

  def part1() do
    lines = parse_lines()

    Enum.count(lines, fn line -> safe?(line) end)
  end

  defp safe?(line) do
    [first, second | _rest] = line

    cond do
      first > second -> safe?(line, fn a, b -> a - b end)
      first < second -> safe?(line, fn a, b -> b - a end)
      true -> false
    end
  end

  defp safe?([_head | []], _differ) do
    true
  end

  defp safe?([head | tail], differ) do
    diff = differ.(head, Enum.at(tail, 0))

    cond do
      diff > 0 and diff < 4 -> safe?(tail, differ)
      true -> false
    end
  end

  def part2() do
    lines = parse_lines()

    Enum.count(lines, fn line ->
      if safe?(line) do
        true
      else
        Enum.with_index(line) |> Enum.any?(fn {_n, i} -> safe?(List.delete_at(line, i)) end)
      end
    end)
  end
end

IO.puts("part 1 - " <> Integer.to_string(Day2.part1()))
IO.puts("part 2 - " <> Integer.to_string(Day2.part2()))
