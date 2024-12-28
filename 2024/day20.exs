Code.require_file("./grid.ex")

defmodule Day20 do
  defmodule Move do
    defstruct [:pos, :score]
  end

  def solve(file, cheat_amount, target) do
    grid = Grid.parse(file)
    start = Grid.find(grid, &(&1 == "S"))
    finish = Grid.find(grid, &(&1 == "E"))
    track = Grid.as_map(grid, &(&1 == "." || &1 == "S" || &1 == "E"))
    track = Map.put(track, :max_x, Grid.max_x(grid))
    track = Map.put(track, :max_y, Grid.max_y(grid))

    distances =
      search(track, finish, %{start => 0}, :queue.from_list([%Move{pos: start, score: 0}]))

    d = Map.to_list(distances)

    cheats =
      for {{ix, iy}, di} <- d, {{jx, jy}, dj} <- d do
        mh = abs(jx - ix) + abs(jy - iy)

        if mh <= cheat_amount && dj >= di + mh do
          dj - (di + mh)
        else
          0
        end
      end

    Enum.filter(cheats, &(&1 >= target)) |> Enum.count()
  end

  defp search(track, finish, visited, queue) do
    if(:queue.is_empty(queue)) do
      visited
    else
      {{_, current}, queue} = :queue.out(queue)

      possible_moves =
        for dir <- [Grid.up(), Grid.down(), Grid.left(), Grid.right()] do
          %Move{pos: Grid.add(current.pos, dir), score: current.score + 1}
        end

      possible_moves =
        Enum.reject(possible_moves, fn move ->
          visited[move.pos] || !track[move.pos] || outside?(track, move.pos)
        end)

      queue =
        for move <- possible_moves, reduce: queue do
          acc -> :queue.in(move, acc)
        end

      visited =
        for move <- possible_moves, into: visited do
          {move.pos, move.score}
        end

      search(track, finish, visited, queue)
    end
  end

  defp outside?(track, {x, y}), do: x < 0 || y < 0 || x > track.max_x || y > track.max_y
end

[file, cheat_amount, target] = System.argv()


Day20.solve(file, String.to_integer(cheat_amount), String.to_integer(target)) |> IO.inspect()
