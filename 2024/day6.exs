defmodule Day6 do
  defp load_grid() do
    File.stream!("test/day6.txt")
    |> Enum.to_list()
    |> Enum.map(&String.trim/1)
    |> List.to_tuple()
  end

  def part1() do
    grid = load_grid()
    start = find_start(grid)

    positions_until_leave(grid, start, {0, -1}, MapSet.new([start]))
    |> MapSet.size()
    |> IO.inspect()
  end

  defp find_start(grid) do
    find_start(grid, 0, 0)
  end

  defp find_start(grid, x, y) do
    cond do
      x > max_x(grid) ->
        find_start(grid, 0, y + 1)

      at(grid, x, y) == "^" ->
        {x, y}

      true ->
        find_start(grid, x + 1, y)
    end
  end

  defp max_x(grid) do
    String.length(elem(grid, 0)) - 1
  end

  defp max_y(grid), do: tuple_size(grid) - 1

  defp at(grid, x, y) do
    elem(grid, y) |> String.at(x)
  end

  defp positions_until_leave(grid, current, direction, visited) do
    newPos = add_coord(current, direction)
    {newX, newY} = newPos

    cond do
      newX < 0 || newX > max_x(grid) ->
        visited

      newY < 0 || newY > max_y(grid) ->
        visited

      at(grid, newX, newY) == "#" ->
        positions_until_leave(grid, current, turn_right(direction), visited)

      true ->
        positions_until_leave(grid, newPos, direction, MapSet.put(visited, newPos))
    end
  end

  defp add_coord({ax, ay}, {bx, by}) do
    {ax + bx, ay + by}
  end

  defp turn_right(coord) do
    case coord do
      # right to down
      {1, 0} -> {0, 1}
      # down to left
      {0, 1} -> {-1, 0}
      # left to up
      {-1, 0} -> {0, -1}
      # up to right
      {0, -1} -> {1, 0}
    end
  end

  def part2() do
    grid = load_grid()
    start = find_start(grid)

    potential_blockages =
      positions_until_leave(grid, start, {0, -1}, MapSet.new([start])) |> Enum.to_list()

    find_loops(grid, start, potential_blockages, 0) |> IO.inspect()
  end

  defp find_loops(_grid, _start, [], acc), do: acc

  defp find_loops(grid, start, [head | tail], acc) do
    cond do
      is_loop?(grid, start, head) -> find_loops(grid, start, tail, acc + 1)
      true -> find_loops(grid, start, tail, acc)
    end
  end

  defp is_loop?(grid, start, {x, y}) do
    current_line = elem(grid, y)

    new_line =
      current_line
      |> String.graphemes()
      |> put_in([Access.at(x)], "#")
      |> Enum.join()

    grid_with_extra = put_elem(grid, y, new_line)
    is_loop?(grid_with_extra, start, {0, -1}, MapSet.new([{start, {0, -1}}]))
  end

  defp is_loop?(grid, current, direction, visited) do
    newPos = add_coord(current, direction)
    {newX, newY} = newPos

    cond do
      newX < 0 || newX > max_x(grid) ->
        false

      newY < 0 || newY > max_y(grid) ->
        false

      Enum.member?(visited, {newPos, direction}) ->
        true

      at(grid, newX, newY) == "#" ->
        is_loop?(grid, current, turn_right(direction), visited)

      true ->
        is_loop?(grid, newPos, direction, MapSet.put(visited, {newPos, direction}))
    end
  end
end

IO.puts("part 1")
Day6.part1()
IO.puts("part 2")
Day6.part2()
