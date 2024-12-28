Code.require_file("./grid.ex")

defmodule Day10 do
  defp file(), do: "test/day10.txt"

  def part1() do
    grid = Grid.parse(file(), &String.to_integer/1)

    for start <- find_start_points(grid) do
      find_trails(grid, start, [], MapSet.new())
      |> List.flatten()
      |> Enum.uniq()
      |> Enum.count()
    end
    |> Enum.sum()
  end

  def part2() do
    grid = Grid.parse(file(), &String.to_integer/1)

    for start <- find_start_points(grid) do
      find_trails(grid, start, [], MapSet.new())
      |> List.flatten()
      |> Enum.count()
    end
    |> Enum.sum()
  end

  defp find_trails(grid, current, acc, visited) do
    current_tile = Grid.at(grid, current)

    if current_tile == 9 do
      [current | acc]
    else
      new_visited = MapSet.put(visited, current)

      for move <- find_moves(grid, current, new_visited) do
        find_trails(grid, move, acc, new_visited) |> List.flatten()
      end ++ acc
    end
  end

  defp find_moves(grid, current, visited) do
    for move <- [Grid.up(), Grid.down(), Grid.left(), Grid.right()] do
      new_pos = Grid.add(current, move)

      if !MapSet.member?(visited, new_pos) && Grid.contains?(grid, new_pos) &&
           Grid.at(grid, new_pos) == Grid.at(grid, current) + 1 do
        new_pos
      end
    end
    |> Enum.filter(&Function.identity/1)
  end

  defp find_start_points(grid) do
    for {row, y} <- Tuple.to_list(grid) |> Enum.with_index(),
        {elem, x} <- Tuple.to_list(row) |> Enum.with_index() do
      if elem == 0 do
        {x, y}
      end
    end
    |> Enum.filter(&Function.identity/1)
  end
end

IO.puts("part 1")
Day10.part1() |> IO.inspect()
IO.puts("part 2")
Day10.part2() |> IO.inspect()
