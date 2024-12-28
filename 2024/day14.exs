defmodule Day14 do
  defp parse(file) do
    File.stream!(file)
    |> Enum.map(&String.trim/1)
    |> Enum.map(fn line ->
      [_, x, y, dx, dy] = Regex.run(~r/p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)/, line)
      {{int(x), int(y)}, {int(dx), int(dy)}}
    end)
  end

  defp int(s), do: String.to_integer(s)

  def part1(file, w, h) do
    final_positions =
      for robot <- parse(file) do
        move(robot, 100, w, h)
      end

    half_h = Integer.floor_div(h - 1, 2)
    half_w = Integer.floor_div(w - 1, 2)

    %{:q1 => q1, :q2 => q2, :q3 => q3, :q4 => q4} =
      for {x, y} <- final_positions, reduce: %{} do
        acc ->
          cond do
            x < half_w && y < half_h -> Map.update(acc, :q1, 1, &(&1 + 1))
            x > half_w && y < half_h -> Map.update(acc, :q2, 1, &(&1 + 1))
            x < half_w && y > half_h -> Map.update(acc, :q3, 1, &(&1 + 1))
            x > half_w && y > half_h -> Map.update(acc, :q4, 1, &(&1 + 1))
            true -> acc
          end
      end

    q1 * q2 * q3 * q4
  end

  defp move({{start_x, start_y}, {dx, dy}}, seconds, w, h) do
    x = start_x + dx * seconds
    y = start_y + dy * seconds

    {wrap(x, w), wrap(y, h)}
  end

  defp wrap(coord, max) do
    c =
      if(coord < 0) do
        norm = Integer.floor_div(-coord, max) + 1
        coord + max * norm
      else
        coord
      end

    rem(c, max)
  end

  def part2(file, w, h) do
    robots = parse(file)
    find_line(robots, 1, 100_000, w, h)
  end

  defp find_line(robots, seconds, max, w, h) do
    if seconds > max do
      nil
    else
      by_y =
        for robot <- robots, reduce: %{} do
          acc ->
            {x, y} = move(robot, seconds, w, h)
            Map.update(acc, y, [x], &[x | &1])
        end

      found_line? = Map.values(by_y)
      |> Enum.any?(fn xs -> find_longest_line(xs) >= 10 end)

      if(found_line?) do
        seconds
      else
        find_line(robots, seconds + 1, max, w, h)
      end
    end
  end

  def find_longest_line(xs) do
    Enum.sort(xs)
    |> Enum.scan({0, -1}, fn x, {count, prevx} ->
      if x == prevx + 1 do
        {count + 1, x}
      else
        {1, x}
      end
    end)
    |> Enum.map(&elem(&1, 0))
    |> Enum.max()
  end
end

file = "test/day14.txt"
width = 101
height = 103

IO.puts("part 1")
Day14.part1(file, width, height) |> IO.inspect()
IO.puts("part 2")
Day14.part2(file, width, height) |> IO.inspect()
