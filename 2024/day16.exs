Code.require_file("./grid.ex")

defmodule Day16 do
  defmodule Move do
    defstruct [:pos, :dir, :score, :path]
  end

  def part1(file) do
    maze = Grid.parse(file, &String.to_atom/1)
    Grid.print(maze)
    start = Grid.find(maze, &(&1 == :S))
    maze = Grid.set(maze, start, :.)
    search(maze, MapSet.new(), [%Move{pos: start, dir: Grid.right(), score: 0}])
  end

  defp search(_, _, []), do: raise("lost")

  defp search(maze, visited, [current | move_queue]) do
    tile = Grid.at(maze, current.pos)
    new_visited = MapSet.put(visited, {current.pos, current.dir})

    case(tile) do
      :E ->
        current.score

      # This move was bad, just drop it
      :"#" ->
        search(maze, new_visited, move_queue)

      :. ->
        queue = [
          %Move{
            pos: Grid.add(current.pos, current.dir),
            dir: current.dir,
            score: current.score + 1
          },
          %Move{
            pos: current.pos,
            dir: turn_clock(current.dir),
            score: current.score + 1000
          },
          %Move{
            pos: current.pos,
            dir: turn_anti(current.dir),
            score: current.score + 1000
          }
          | move_queue
        ]

        queue =
          Enum.reject(queue, &({&1.pos, &1.dir} in visited))
          |> Enum.sort_by(& &1.score)

        search(maze, new_visited, queue)
    end
  end

  def part2(file, target) do
    maze = Grid.parse(file, &String.to_atom/1)
    start = Grid.find(maze, &(&1 == :S))
    maze = Grid.set(maze, start, :.)

    queue = :queue.new()
    queue = :queue.in(%Move{pos: start, dir: Grid.right(), score: 0, path: [start]}, queue)

    search_all(
      maze,
      %{start: Grid.right()},
      target,
      queue,
      []
    )
    |> List.flatten()
    |> Enum.uniq()
    |> Enum.count()
  end

  defp search_all(maze, visited, target, queue, acc) do
    if :queue.is_empty(queue) do
      acc
    else
      do_search_all(maze, visited, target, queue, acc)
    end
  end

  defp do_search_all(maze, visited, target, queue, acc) do
    {{_, current}, queue} = :queue.out(queue)
    tile = Grid.at(maze, current.pos)
    new_visited = Map.put(visited, {current.pos, current.dir}, current.score)

    cond do
      current.score > Map.get(visited, {current.pos, current.dir}, 100_000_000) ->
        search_all(maze, visited, target, queue, acc)

      tile == :"#" ->
        search_all(maze, visited, target, queue, acc)

      current.score > target ->
        search_all(maze, visited, target, queue, acc)

      tile == :E ->
        search_all(maze, visited, target, queue, [current.path | acc])

      tile == :. ->
        new_pos = Grid.add(current.pos, current.dir)

        queue =
          :queue.in(
            %Move{
              pos: new_pos,
              dir: current.dir,
              score: current.score + 1,
              path: [new_pos | current.path]
            },
            queue
          )

        queue =
          :queue.in(
            %Move{
              pos: current.pos,
              dir: turn_clock(current.dir),
              score: current.score + 1000,
              path: current.path
            },
            queue
          )

        queue =
          :queue.in(
            %Move{
              pos: current.pos,
              dir: turn_anti(current.dir),
              score: current.score + 1000,
              path: current.path
            },
            queue
          )

        search_all(maze, new_visited, target, queue, acc)
    end
  end

  defp turn_clock({x, y}), do: {-y, x}

  defp turn_anti({x, y}), do: {y, -x}
end

file = "test/day16.txt"
IO.puts("part 1")
target = Day16.part1(file) |> IO.inspect()
IO.puts("part 2")
Day16.part2(file, target) |> IO.inspect()
