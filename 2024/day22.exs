defmodule Day22 do
  defp parse(file) do
    File.stream!(file) |> Enum.map(&String.trim/1) |> Enum.map(&String.to_integer/1)
  end

  def part1(file) do
    for secret <- parse(file), reduce: 0 do
      acc ->
        {secret, _} =
          Stream.iterate(secret, &evolve/1)
          |> Stream.with_index()
          |> Enum.find(fn {_, i} -> i == 2000 end)

        acc + secret
    end
  end

  def part2(file) do
    for secret <- parse(file) do
      Stream.iterate({secret, nil}, fn {secret, _} ->
        new_secret = evolve(secret)
        diff = last_dig(new_secret) - last_dig(secret)
        {new_secret, diff}
      end)
      |> Enum.take(2000)
      |> Enum.drop(1)
      |> Enum.chunk_every(4, 1, :discard)
      |> Enum.map(fn [_, _, _, {sec, _}] = chunk ->
        changes = Enum.map(chunk, &elem(&1, 1))
        {changes, last_dig(sec)}
      end)
      |> Enum.uniq_by(fn {changes, _} -> changes end)
    end
    |> List.flatten()
    |> Enum.group_by(&elem(&1, 0), &elem(&1, 1))
    |> Map.values()
    |> Enum.map(&Enum.sum(&1))
    |> Enum.max()
  end

  def last_dig(n), do: Integer.mod(n, 10)

  def evolve(secret) do
    secret = Bitwise.bxor(secret * 64, secret) |> prune
    secret = Bitwise.bxor(Integer.floor_div(secret, 32), secret) |> prune
    Bitwise.bxor(secret * 2048, secret) |> prune
  end

  defp prune(number), do: Integer.mod(number, 16_777_216)
end

[file] = System.argv()
IO.puts("Part 1")
Day22.part1(file) |> IO.inspect()
IO.puts("Part 2")
Day22.part2(file) |> IO.inspect(charlists: :as_lists)
