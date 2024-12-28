Code.require_file("./grid.ex")

defmodule Day12 do
  defp file(), do: "test/day12.txt"

  def part1() do
    grid = Grid.parse(file())

    regions(Grid.as_map(grid), [])
    |> Enum.map(fn region ->
      region_set = MapSet.new(region)
      MapSet.size(region_set) * perimiter(region)
    end)
    |> Enum.sum()
  end

  def part2() do
    grid = Grid.parse(file())

    regions(Grid.as_map(grid), [])
    |> Enum.map(fn region ->
      region_set = MapSet.new(region)
      MapSet.size(region_set) * sides(region)
    end)
    |> Enum.sum()
  end

  defp regions(tiles, acc) when map_size(tiles) == 0, do: acc

  defp regions(tiles, acc) do
    {pos, plant} = Enum.at(tiles, 0)
    {rest_tiles, region} = fill_region(tiles, plant, pos, [])
    regions(rest_tiles, [region | acc])
  end

  defp fill_region(tiles, plant, pos, acc) do
    if tiles[pos] == plant do
      acc = [pos | acc]
      tiles = Map.delete(tiles, pos)

      for neighb <- Grid.neighbours(pos), reduce: {tiles, acc} do
        {tiles, acc} -> fill_region(tiles, plant, neighb, acc)
      end
    else
      {tiles, acc}
    end
  end

  defp perimiter(region) do
    border_tiles(region)
    |> Enum.map(fn tile ->
      Grid.neighbours(tile)
      |> Enum.count(&(&1 not in region))
    end)
    |> Enum.sum()
  end

  def border_tiles(region) do
    Enum.filter(region, fn tile ->
      Grid.neighbours(tile)
      |> Enum.any?(&(&1 not in region))
    end)
  end

  def sides(region) do
    borders =
      border_tiles(region)
      |> Enum.flat_map(fn tile ->
        acc = (Grid.up(tile) in region && []) || [{tile, :up}]
        acc = (Grid.down(tile) in region && acc) || [{tile, :down} | acc]
        acc = (Grid.left(tile) in region && acc) || [{tile, :left} | acc]
        (Grid.right(tile) in region && acc) || [{tile, :right} | acc]
      end)
      |> MapSet.new()

    count_sides(borders, 0)
  end

  defp remove_border(border_tiles, {tile, dir} = entry) when dir in [:up, :down] do
    if entry in border_tiles do
      border_tiles = MapSet.delete(border_tiles, entry)
      remove_border(border_tiles, {Grid.right(tile), dir})
    else
      border_tiles
    end
  end

  defp remove_border(border_tiles, {tile, dir} = entry) when dir in [:left, :right] do
    if entry in border_tiles do
      border_tiles = MapSet.delete(border_tiles, entry)
      remove_border(border_tiles, {Grid.down(tile), dir})
    else
      border_tiles
    end
  end

  defp count_sides(borders, acc) do
    if MapSet.size(borders) == 0 do
      acc
    else
      borders = remove_border(borders, Enum.min(borders))
      count_sides(borders, acc + 1)
    end
  end
end

IO.puts("part 1")
Day12.part1() |> IO.inspect()
IO.puts("part 2")
Day12.part2() |> IO.inspect()
