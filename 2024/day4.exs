defmodule Day4 do
  defp load_grid() do
    File.stream!("test/day4.txt")
    |> Enum.to_list()
    |> List.to_tuple()
  end

  def part1() do
    grid = load_grid()
    IO.inspect(grid)
    countXmas(grid, 0, 0, 0)
  end

  defp at(grid, x, y) do
    elem(grid, y) |> String.at(x)
  end

  defp countXmas(grid, x, y, acc) do
    maxY = tuple_size(grid) - 1
    maxX = String.length(elem(grid, 0)) - 1

    cond do
      y > maxY ->
        acc

      x > maxX ->
        countXmas(grid, 0, y + 1, acc)

      true ->
        countFromPos =
          left(grid, x, y) + right(grid, maxX, x, y) + up(grid, x, y) + down(grid, maxY, x, y) +
            up_right(grid, maxX, x, y) + down_right(grid, maxX, maxY, x, y) + up_left(grid, x, y) +
            down_left(grid, maxY, x, y)

        countXmas(grid, x + 1, y, acc + countFromPos)
    end
  end

  defp right(grid, maxX, x, y) do
    cond do
      x + 3 > maxX -> 0
      String.slice(elem(grid, y), x..(x + 3)) == "XMAS" -> 1
      true -> 0
    end
  end

  defp left(grid, x, y) do
    cond do
      x - 3 < 0 -> 0
      String.reverse(String.slice(elem(grid, y), (x - 3)..x)) == "XMAS" -> 1
      true -> 0
    end
  end

  defp up(grid, x, y) do
    cond do
      y - 3 < 0 ->
        0

      at(grid, x, y) <> at(grid, x, y - 1) <> at(grid, x, y - 2) <> at(grid, x, y - 3) == "XMAS" ->
        1

      true ->
        0
    end
  end

  defp down(grid, maxY, x, y) do
    cond do
      y + 3 > maxY ->
        0

      at(grid, x, y) <> at(grid, x, y + 1) <> at(grid, x, y + 2) <> at(grid, x, y + 3) == "XMAS" ->
        1

      true ->
        0
    end
  end

  defp up_right(grid, maxX, x, y) do
    cond do
      x + 3 > maxX || y - 3 < 0 ->
        0

      at(grid, x, y) <> at(grid, x + 1, y - 1) <> at(grid, x + 2, y - 2) <> at(grid, x + 3, y - 3) ==
          "XMAS" ->
        1

      true ->
        0
    end
  end

  defp down_right(grid, maxX, maxY, x, y) do
    cond do
      x + 3 > maxX || y + 3 > maxY ->
        0

      at(grid, x, y) <> at(grid, x + 1, y + 1) <> at(grid, x + 2, y + 2) <> at(grid, x + 3, y + 3) ==
          "XMAS" ->
        1

      true ->
        0
    end
  end

  defp down_left(grid, maxY, x, y) do
    cond do
      x - 3 < 0 || y + 3 > maxY ->
        0

      at(grid, x, y) <> at(grid, x - 1, y + 1) <> at(grid, x - 2, y + 2) <> at(grid, x - 3, y + 3) ==
          "XMAS" ->
        1

      true ->
        0
    end
  end

  defp up_left(grid, x, y) do
    cond do
      x - 3 < 0 || y - 3 < 0 ->
        0

      at(grid, x, y) <> at(grid, x - 1, y - 1) <> at(grid, x - 2, y - 2) <> at(grid, x - 3, y - 3) ==
          "XMAS" ->
        1

      true ->
        0
    end
  end

  def part2() do
    grid = load_grid()
    countMas(grid, 0, 0, 0)
  end

  defp countMas(grid, x, y, acc) do
    maxY = tuple_size(grid) - 1
    maxX = String.length(elem(grid, 0)) - 1

    cond do
      y + 2 > maxY -> acc
      x + 2 > maxX -> countMas(grid, 0, y + 1, acc)
      xMas?(grid, x, y) -> countMas(grid, x + 1, y, acc + 1)
      true -> countMas(grid, x + 1, y, acc)
    end
  end

  defp xMas?(grid, x, y) do
    topLeft = at(grid, x, y)
    topRight = at(grid, x + 2, y)
    middle = at(grid, x + 1, y + 1)
    bottomLeft = at(grid, x, y + 2)
    bottomRight = at(grid, x + 2, y + 2)

    cross1 = topLeft <> middle <> bottomRight
    cross2 = topRight <> middle <> bottomLeft

    (cross1 == "MAS" || cross1 == "SAM") && (cross2 == "MAS" || cross2 == "SAM")
  end
end

IO.puts("Part 1 - " <> Integer.to_string(Day4.part1()))
IO.puts("Part 2 - " <> Integer.to_string(Day4.part2()))
