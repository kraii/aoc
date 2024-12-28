Mix.install([{:memoize, "~> 1.4"}])

defmodule Day21 do
  use Memoize

  defp make_numpad,
    do: %{
      "0" => {1, 3},
      "1" => {0, 2},
      "2" => {1, 2},
      "3" => {2, 2},
      "4" => {0, 1},
      "5" => {1, 1},
      "6" => {2, 1},
      "7" => {0, 0},
      "8" => {1, 0},
      "9" => {2, 0},
      "A" => {2, 3}
    }

  defp make_dir_pad,
    do: %{
      "^" => {1, 0},
      "A" => {2, 0},
      "<" => {0, 1},
      "v" => {1, 1},
      ">" => {2, 1}
    }

  defp parse(file) do
    File.stream!(file) |> Enum.map(&String.trim/1) |> Enum.to_list()
  end

  def solve(file, pads) do
    for code <- parse(file), reduce: 0 do
      acc ->
        code_pushes =
          Enum.chunk_every(["A" | String.graphemes(code)], 2, 1, :discard)
          |> Enum.reduce(0, fn [start, finish], acc ->
            acc + sequence(make_numpad(), start, finish, pads)
          end)

        acc + code_pushes * value(code)
    end
  end

  defp value(code), do: String.slice(code, 0, 3) |> String.to_integer()

  defmemo sequence(keypad, current, finish, 0) do
    # Base case is manhatten distance
    {cx, cy} = keypad[current]
    {fx, fy} = keypad[finish]
    abs(fx - cx) + abs(fy - cy) + 1
  end

  defmemo sequence(keypad, current, finish, robots) do
    pushes =
      for moves <- possible_moves(keypad, current, finish) do
        {_, _, pushes} =
          Stream.iterate({"A", moves, 0}, fn
            {_, [], _} ->
              false

            {current, [next | rest], acc} ->
              acc = acc + sequence(make_dir_pad(), current, next, robots - 1)
              {next, rest, acc}
          end)
          |> Stream.take_while(&Function.identity/1)
          |> Enum.at(-1)

        from_end = Enum.at(moves, -1) || "A"
        pushes + sequence(make_dir_pad(), from_end, "A", robots - 1)
      end

    Enum.min(pushes)
  end

  def possible_moves(keypad, from, to) do
    {cx, cy} = keypad[from]
    {fx, fy} = keypad[to]
    dx = fx - cx
    dy = fy - cy

    moves = []

    moves =
      cond do
        dy < 0 -> moves ++ List.duplicate("^", abs(dy))
        dy > 0 -> moves ++ List.duplicate("v", dy)
        true -> moves
      end

    moves =
      cond do
        dx < 0 -> moves ++ List.duplicate("<", abs(dx))
        dx > 0 -> moves ++ List.duplicate(">", dx)
        true -> moves
      end

    permutations(moves)
    |> Enum.filter(&all_valid?(keypad, from, &1))
    |> Enum.uniq()
  end

  defp all_valid?(keypad, from, moves) do
    valid_pos = Map.values(keypad)

    Enum.scan(moves, keypad[from], fn move, {cx, cy} ->
      case(move) do
        "^" -> {cx, cy - 1}
        "v" -> {cx, cy + 1}
        "<" -> {cx - 1, cy}
        ">" -> {cx + 1, cy}
      end
    end)
    |> Enum.all?(&(&1 in valid_pos))
  end

  defp permutations([]), do: [[]]

  defp permutations(list) do
    for elem <- list, rest <- permutations(list -- [elem]), do: [elem | rest]
  end
end

[file] = System.argv()

IO.puts("Part 1")
Day21.solve(file, 2) |> IO.inspect()
IO.puts("Part 2")
Day21.solve(file, 25) |> IO.inspect()
