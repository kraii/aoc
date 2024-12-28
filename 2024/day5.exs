defmodule Day5 do
  defp file(), do: "test/day5.txt"

  def part1() do
    {updates, rules} = parse()
    Enum.reduce(updates, 0, countMiddleIfMatchesRules(rules))
  end

  defp parse() do
    {rulesData, updatesWithSep} =
      File.stream!(file())
      |> Stream.map(&String.trim/1)
      |> Enum.split_while(fn line -> String.length(line) > 0 end)

    updates =
      Enum.drop(updatesWithSep, 1)
      |> Enum.map(&String.split(&1, ","))
      |> Enum.map(fn update -> Enum.map(update, &String.to_integer/1) end)

    rules =
      Enum.map(rulesData, fn rule ->
        [first, second] = String.split(rule, "|")
        {String.to_integer(first), String.to_integer(second)}
      end)

    {updates, rules}
  end

  defp countMiddleIfMatchesRules(rules) do
    fn update, acc ->
      if(matchesRules(rules, update)) do
        acc + mid(update)
      else
        acc
      end
    end
  end

  defp mid(update) do
    Enum.at(update, Integer.floor_div(length(update), 2))
  end

  defp matchesRules(rules, update) do
    index = Enum.with_index(update) |> Enum.into(%{})

    Enum.all?(rules, fn {l, r} ->
      il = Map.get(index, l)
      ir = Map.get(index, r)
      il == nil || ir == nil || il < ir
    end)
  end

  def part2() do
    {updates, rules} = parse()
    Enum.reduce(updates, 0, countMiddleOfSorted(rules))
  end

  defp countMiddleOfSorted(rules) do
    fn update, acc ->
      if matchesRules(rules, update) do
        acc
      else
        sorted = Enum.sort(update, sorter(rules))
        acc + mid(sorted)
      end
    end
  end

  defp sorter(rules) do
    fn a, b ->
      Enum.any?(rules, fn {l, r} -> a == r && b == l end)
    end
  end
end

IO.puts("Part 1 -> " <> Integer.to_string(Day5.part1()))
IO.puts("Part 2 -> " <> Integer.to_string(Day5.part2()))
