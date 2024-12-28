defmodule Grid do
  def parse(file) do
    parse(file, &Function.identity/1)
  end

  def parse(file, element_fun) do
    File.stream!(file)
    |> parse_lines(element_fun)
  end

  def parse_lines(lines, element_fun) do
    Enum.map(lines, fn line ->
      String.trim(line)
      |> String.graphemes()
      |> Enum.map(element_fun)
      |> List.to_tuple()
    end)
    |> List.to_tuple()
  end

  def at(grid, {x, y}), do: elem(grid, y) |> elem(x)

  def max_x(grid), do: tuple_size(elem(grid, 0)) - 1

  def max_y(grid), do: tuple_size(grid) - 1

  def add({ax, ay}, {bx, by}), do: {ax + bx, ay + by}

  def contains?(grid, {x, y}), do: x >= 0 and x <= max_x(grid) && y >= 0 && y <= max_y(grid)

  def up(), do: {0, -1}
  def down(), do: {0, 1}
  def left(), do: {-1, 0}
  def right(), do: {1, 0}
  def up(pos), do: add(pos, up())
  def down(pos), do: add(pos, down())
  def left(pos), do: add(pos, left())
  def right(pos), do: add(pos, right())

  def neighbours(pos), do: [up(pos), down(pos), left(pos), right(pos)]

  def neighbours(grid, pos), do: neighbours(pos) |> Enum.filter(&contains?(grid, &1))

  def as_map(grid), do: as_map(grid, &Function.identity/1)

  def as_map(grid, terrain_fn) do
    for(
      x <- 0..Grid.max_x(grid),
      y <- 0..Grid.max_y(grid),
      into: %{}
    ) do
      {{x, y}, terrain_fn.(Grid.at(grid, {x, y}))}
    end
  end

  def find(grid, matcher) do
    do_find(Tuple.to_list(grid) |> Enum.with_index(), matcher)
  end

  defp do_find([], _), do: nil

  defp do_find([{row, y} | tail], matcher) do
    match = Tuple.to_list(row) |> Enum.with_index() |> Enum.find(&matcher.(elem(&1, 0)))

    if match do
      {_, x} = match
      {x, y}
    else
      do_find(tail, matcher)
    end
  end

  def set(grid, {x, y}, value) do
    new_row = elem(grid, y) |> put_elem(x, value)
    put_elem(grid, y, new_row)
  end

  def print(grid) do
    Tuple.to_list(grid)
    |> Enum.each(fn line ->
      Tuple.to_list(line) |> Enum.join() |> IO.puts()
    end)
  end
end
