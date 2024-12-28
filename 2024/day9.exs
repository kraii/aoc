defmodule Day9 do
  defp read_file() do
    File.read!("test/day9.txt")
    |> String.trim()
    |> String.graphemes()
    |> Enum.with_index()
  end

  def part1() do
    data = read_file()

    fs =
      for {block, index} <- data do
        size = String.to_integer(block)

        if(rem(index, 2) == 0) do
          List.duplicate({:file, Integer.floor_div(index, 2)}, size)
        else
          List.duplicate({:free}, size)
        end
      end
      |> List.flatten()

    frag(fs, [])
    |> Enum.reverse()
    |> Enum.with_index()
    |> Enum.reduce(0, fn {{:file, id}, idx}, acc -> acc + id * idx end)
  end

  defp frag([], acc), do: acc

  defp frag([{:file, id} | fs], acc), do: frag(fs, [{:file, id} | acc])

  defp frag([{:free} | fs], acc) do
    last = List.last(fs)

    case last do
      {:file, id} -> frag([{:file, id} | Enum.drop(fs, -1)], acc)
      {:free} -> frag([{:free} | Enum.drop(fs, -1)], acc)
      nil -> acc
    end
  end

  def part2() do
    data = read_file()

    {_, fs} =
      for {block, index} <- data, reduce: {0, []} do
        {block_start, acc} ->
          size = String.to_integer(block)

          if(rem(index, 2) == 0) do
            {block_start + size,
             [
               %{type: :file, size: size, index: Integer.floor_div(index, 2), start: block_start}
               | acc
             ]}
          else
            {block_start + size, [%{type: :free, size: size, start: block_start} | acc]}
          end
      end

    compressed =
      for chunk <- fs, reduce: Enum.reverse(fs) do
        acc ->
          cond do
            chunk.type == :file ->
              try_move(chunk, acc)

            true ->
              acc
          end
      end

    Enum.sort_by(compressed, fn chunk -> chunk.start end)
    |> Enum.reject(fn chunk -> chunk.size == 0 end)
    |> Enum.reduce(0, fn chunk, acc ->
      acc + total_score(chunk, chunk.size + chunk.start - 1, 0, chunk.size)
    end)
  end

  defp total_score(%{type: :free}, _, _, _), do: 0
  defp total_score(_, _, acc, 0), do: acc

  defp total_score(chunk, block_index, acc, remaining) do
    total_score(
      chunk,
      block_index - 1,
      acc + chunk.index * block_index,
      remaining - 1
    )
  end

  defp try_move(chunk, fs) do
    available_space =
      Enum.find_index(fs, fn c ->
        c.type == :free && c.size >= chunk.size && c.start < chunk.start
      end)

    if(available_space) do
      with_chunk_removed = List.delete(fs, chunk)
      space = Enum.at(fs, available_space)

      List.update_at(with_chunk_removed, available_space, fn space ->
        %{type: :file, size: chunk.size, index: chunk.index, start: space.start}
      end)
      |> List.insert_at(available_space + 1, %{
        type: :free,
        size: space.size - chunk.size,
        start: space.start + chunk.size
      })
    else
      fs
    end
  end
end

IO.puts("part 1")
Day9.part1() |> IO.inspect()
IO.puts("part 2")
Day9.part2() |> IO.inspect()
