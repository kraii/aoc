defmodule Day11 do
  defp parse() do
    File.read!("test/day11.txt")
    |> String.trim()
    |> String.split()
    |> Enum.map(&String.to_integer/1)
  end

  def solve(blinks) do
    terms = parse()

    start =
      for term <- terms, into: %{} do
        {term, 1}
      end

    stone_counts =
      for _i <- 0..(blinks - 1), reduce: start do
        acc ->
          blink(acc)
      end

    Map.values(stone_counts) |> Enum.sum()
  end

  defp blink(stones) do
    for {stone, count} <- Map.to_list(stones), reduce: %{} do
      acc ->
        cond do
          stone == 0 ->
            Map.update(acc, 1, count, &(&1 + count))

          rem(count_digits(stone), 2) == 0 ->
            digits = Integer.digits(stone)

            {first_digits, second_digits} =
              Enum.split(digits, Integer.floor_div(Enum.count(digits), 2))

            Map.update(acc, Integer.undigits(first_digits), count, &(&1 + count))
            |> Map.update(Integer.undigits(second_digits), count, &(&1 + count))

          true ->
            Map.update(acc, stone * 2024, count, &(&1 + count))
        end
    end
  end

  defp count_digits(i), do: Integer.digits(i) |> Enum.count()
end

IO.puts("part 1")
Day11.solve(25) |> IO.inspect()
IO.puts("part 2")
Day11.solve(75) |> IO.inspect()
