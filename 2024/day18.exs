Code.require_file("./grid.ex")

defmodule Day18 do
  defmodule Move do
    defstruct [:pos, :score]
  end

  def parse(file) do
    File.stream!(file)
    |> Enum.map(fn line ->
      [x, y] = String.trim(line) |> String.split(",")
      {String.to_integer(x), String.to_integer(y)}
    end)
  end

  def part1(file, bytes, size) do
    corrupted = parse(file) |> Enum.take(bytes) |> corrupted_map()

    search(corrupted, {size, size}, MapSet.new(), [%Move{pos: {0, 0}, score: 0}])
  end

  defp corrupted_map(bytes) do
    for point <- bytes, into: %{} do
      {point, true}
    end
  end

  defp search(_, _, _, []), do: -1

  defp search(corrupted, finish, visited, [current | move_queue]) do
    if(finish == current.pos) do
      current.score
    else
      possible_moves =
        for dir <- [Grid.up(), Grid.down(), Grid.left(), Grid.right()] do
          %Move{pos: Grid.add(current.pos, dir), score: current.score + 1}
        end

      possible_moves =
        Enum.reject(possible_moves, fn move ->
          move.pos in visited || corrupted[move.pos] ||
            outside?(move.pos, finish)
        end)

      queue =
        (possible_moves ++ move_queue)
        |> Enum.sort_by(& &1.score)

      visited =
        for move <- possible_moves, reduce: visited do
          acc -> MapSet.put(acc, move.pos)
        end

      search(corrupted, finish, visited, queue)
    end
  end

  defp outside?({x, y}, {max_x, max_y}), do: x < 0 || y < 0 || x > max_x || y > max_y

  def part2(file, start_byte, size) do
    falling_bytes = parse(file)

    {i, _} =
      Stream.iterate({start_byte, length(falling_bytes)}, fn {low, high} ->
        mid = Integer.floor_div(low + high, 2)
        corrupted = Enum.take(falling_bytes, mid) |> corrupted_map()

        found = search(corrupted, {size, size}, MapSet.new(), [%Move{pos: {0, 0}, score: 0}])

        if(found > 0) do
          {mid + 1, high}
        else
          {low, mid - 1}
        end
      end)
      |> Enum.find(fn {low, high} -> low > high end)

    Enum.at(falling_bytes, i - 1)
  end
end

[file, bytes, size] = System.argv()
IO.puts("Part 1")
Day18.part1(file, String.to_integer(bytes), String.to_integer(size)) |> IO.inspect()
IO.puts("Part 2")
Day18.part2(file, String.to_integer(bytes), String.to_integer(size)) |> IO.inspect()
