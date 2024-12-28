defmodule Day8 do
  defp load_grid() do
    File.stream!("test/day8.txt")
    |> Enum.to_list()
    |> Enum.map(&String.trim/1)
  end

  defp build_point_map(grid) do
    for {row, y} <- Enum.with_index(grid),
        {val, x} <- String.graphemes(row) |> Enum.with_index(),
        reduce: %{} do
      acc ->
        cond do
          val == "." ->
            acc

          Map.has_key?(acc, val) ->
            elem(Map.get_and_update(acc, val, fn coords -> {coords, coords ++ [{x, y}]} end), 1)

          true ->
            Map.put(acc, val, [{x, y}])
        end
    end
  end

  def part1() do
    grid = load_grid()
    types_to_points = build_point_map(grid)
    maxX = max_x(grid)
    maxY = max_y(grid)

    solve(types_to_points, fn {_type, pairs} ->
      for {{ax, ay}, {bx, by}} <- pairs do
        diffX = ax - bx
        diffY = ay - by

        [{ax + diffX, ay + diffY}, {bx - diffX, by - diffY}]
        |> Enum.reject(fn {x, y} -> x < 0 || x > maxX || y < 0 || y > maxY end)
      end
    end)
  end

  def part2() do
    grid = load_grid()
    types_to_points = build_point_map(grid)
    maxX = max_x(grid)
    maxY = max_y(grid)

    solve(types_to_points, fn {_type, pairs} ->
      for {{ax, ay}, {bx, by}} <- pairs do
        diffX = ax - bx
        diffY = ay - by

        find_nodes({ax, ay}, {diffX, diffY}, maxX, maxY, []) ++
          find_nodes({bx, by}, {-diffX, -diffY}, maxX, maxY, [])
      end
    end)
  end

  defp solve(types_to_points, calcfn) do
    Map.to_list(types_to_points)
    |> Enum.map(fn {type, locs} ->
      {type, all_pairs(locs)}
    end)
    |> Enum.flat_map(calcfn)
    |> List.flatten()
    |> Enum.uniq()
    |> Enum.count()
    |> IO.inspect()
  end

  defp find_nodes({x, y}, {diffX, diffY}, maxX, maxY, acc) do
    cond do
      x < 0 || x > maxX -> acc
      y < 0 || y > maxY -> acc
      true -> find_nodes({x + diffX, y + diffY}, {diffX, diffY}, maxX, maxY, [{x, y}] ++ acc)
    end
  end

  defp all_pairs(locs) do
    for i <- locs,
        j <- locs do
      if(i != j) do
        {i, j}
      end
    end
  end

  defp max_x(grid) do
    String.length(Enum.at(grid, 0)) - 1
  end

  defp max_y(grid), do: length(grid) - 1
end

IO.puts("Part 1")
Day8.part1()

IO.puts("Part 2")
Day8.part2()
