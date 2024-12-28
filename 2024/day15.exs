Code.require_file("./grid.ex")

defmodule Day15 do
  defp parse(file, room_translate) do
    {room, moves} =
      File.stream!(file)
      |> Enum.map(&String.trim/1)
      |> Enum.split_while(&(&1 != ""))

    moves =
      Enum.drop(moves, 1) |> Enum.join() |> String.graphemes() |> Enum.map(&String.to_atom/1)

    {Grid.parse_lines(room_translate.(room), &String.to_atom/1), moves}
  end

  def part1(file) do
    {room, moves} = parse(file, &Function.identity/1)
    start_at = Grid.find(room, &(&1 == :@))

    room = do_moves(moves, start_at, room)

    for {row, y} <- Tuple.to_list(room) |> Enum.with_index(),
        {el, x} <- Tuple.to_list(row) |> Enum.with_index(),
        reduce: 0 do
      acc ->
        if el == :O do
          acc + 100 * y + x
        else
          acc
        end
    end
  end

  defp do_moves([], _, room), do: room

  defp do_moves([move | moves], pos, room) do
    delta = to_delta(move)
    next_pos = Grid.add(pos, delta)
    tile = Grid.at(room, next_pos)

    {pos_after_move, room_after_move} =
      case(tile) do
        :. ->
          {next_pos, Grid.set(room, pos, :.) |> Grid.set(next_pos, :@)}

        :"#" ->
          {pos, room}

        :O ->
          move_box(pos, delta, room)

        :"[" ->
          move_big_box(pos, next_pos, next_pos, Grid.right(next_pos), delta, room)

        :"]" ->
          move_big_box(pos, next_pos, Grid.left(next_pos), next_pos, delta, room)
      end

    do_moves(moves, pos_after_move, room_after_move)
  end

  defp move_box(pos, delta, room) do
    next_pos = Grid.add(pos, delta)

    ahead = ahead(next_pos, delta, room)

    {_, non_boxes} = Enum.split_while(ahead, &(Grid.at(room, &1) == :O))
    non_box = hd(non_boxes)
    non_box_tile = Grid.at(room, non_box)

    case(non_box_tile) do
      :"#" ->
        {pos, room}

      :. ->
        room =
          Grid.set(room, pos, :.)
          |> Grid.set(next_pos, :@)
          |> Grid.set(non_box, :O)

        {next_pos, room}
    end
  end

  defp to_delta(move) do
    case(move) do
      :> -> Grid.right()
      :^ -> Grid.up()
      :< -> Grid.left()
      :v -> Grid.down()
    end
  end

  def part2(file) do
    {room, moves} = parse(file, &widen/1)
    start_at = Grid.find(room, &(&1 == :@))

    room = do_moves(moves, start_at, room)

    for {row, y} <- Tuple.to_list(room) |> Enum.with_index(),
        {el, x} <- Tuple.to_list(row) |> Enum.with_index(),
        reduce: 0 do
      acc ->
        if el == :"[" do
          acc + 100 * y + x
        else
          acc
        end
    end
  end

  defp widen(lines) do
    for line <- lines do
      for token <- String.graphemes(line), reduce: "" do
        acc ->
          case token do
            "#" -> acc <> "##"
            "O" -> acc <> "[]"
            "." -> acc <> ".."
            "@" -> acc <> "@."
          end
      end
    end
  end

  defp move_big_box(robot_current, robot_to, box_l, box_r, delta, room) do
    {room, changed} = do_box_move(box_l, box_r, delta, room)

    if(changed) do
      room = Grid.set(room, robot_current, :.) |> Grid.set(robot_to, :@)
      {robot_to, room}
    else
      {robot_current, room}
    end
  end

  defp do_box_move(box_l, box_r, delta, room) do
    next_l_pos = Grid.add(box_l, delta)
    next_r_pos = Grid.add(box_r, delta)
    next_l = Grid.at(room, next_l_pos)
    next_r = Grid.at(room, next_r_pos)

    cond do
      next_l == :"#" || next_r == :"#" ->
        {room, false}

      delta == Grid.left() && next_l == :. ->
        {Grid.set(room, next_l_pos, :"[") |> Grid.set(next_r_pos, :"]") |> Grid.set(box_r, :.),
         true}

      delta == Grid.right() && next_r == :. ->
        {Grid.set(room, next_l_pos, :"[") |> Grid.set(next_r_pos, :"]") |> Grid.set(box_l, :.),
         true}

      next_l == :. && next_r == :. ->
        {
          Grid.set(room, next_l_pos, :"[")
          |> Grid.set(next_r_pos, :"]")
          |> Grid.set(box_l, :.)
          |> Grid.set(box_r, :.),
          true
        }

      delta == Grid.right() && next_r == :"[" ->
        {room, moved} = do_box_move(next_r_pos, Grid.right(next_r_pos), delta, room)

        if moved do
          # Move again now we've made space
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end

      delta == Grid.left() && next_l == :"]" ->
        {room, moved} = do_box_move(Grid.left(next_l_pos), next_l_pos, delta, room)

        if moved do
          # Move again now we've made space
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end

      next_l == :"[" && next_r == :"]" ->
        {room, moved} = do_box_move(next_l_pos, next_r_pos, delta, room)

        if moved do
          # Move again now we've made space
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end

      next_l == :"]" && next_r == :"[" ->
        # two boxes above
        {l_room, l_moved} = do_box_move(Grid.left(next_l_pos), next_l_pos, delta, room)
        {_, r_moved} = do_box_move(next_r_pos, Grid.right(next_r_pos), delta, room)

        if l_moved && r_moved do
          {room, _} = do_box_move(next_r_pos, Grid.right(next_r_pos), delta, l_room)
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end

      (Grid.up() == delta || Grid.down() == delta) && next_l == :"]" ->
        {room, moved} = do_box_move(Grid.left(next_l_pos), next_l_pos, delta, room)

        if moved do
          # Move again now we've made space
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end

      (Grid.up() == delta || Grid.down() == delta) && next_r == :"[" ->
        {room, moved} = do_box_move(next_r_pos, Grid.right(next_r_pos), delta, room)

        if moved do
          # Move again now we've made space
          do_box_move(box_l, box_r, delta, room)
        else
          {room, false}
        end
    end
  end

  defp ahead(pos, delta, room) do
    Stream.iterate(pos, &Grid.add(&1, delta))
    |> Enum.take_while(&Grid.contains?(room, &1))
  end
end

file = "test/day15.txt"
IO.puts("part 1")
Day15.part1(file) |> IO.inspect()
IO.puts("part 2")
Day15.part2(file) |> IO.inspect()
