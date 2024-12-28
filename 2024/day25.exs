defmodule Day25 do
  defp parse(file) do
    File.stream!(file)
    |> Stream.map(&String.trim/1)
    |> Stream.chunk_by(&(&1 == ""))
    |> Stream.reject(&(&1 == [""]))
    |> Enum.split_with(&(hd(&1) == "#####"))
  end

  def part1(file) do
    {locks, keys} = parse(file)

    for lock <- locks, key <- keys, reduce: 0 do
      acc ->
        if fits(lock, key) do
          acc + 1
        else
          acc
        end
    end
  end

  defp fits(lock, key) do
    overlap =
      Enum.zip(lock, key)
      |> Enum.any?(fn {l, k} ->
        Enum.zip(String.graphemes(l), String.graphemes(k))
        |> Enum.any?(fn {ll, kk} -> ll == "#" && kk == "#" end)
      end)

    !overlap
  end
end

[file] = System.argv()
Day25.part1(file) |> IO.inspect()
