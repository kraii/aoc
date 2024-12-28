defmodule Day23 do
  defp parse(file) do
    for line <- File.stream!(file), reduce: %{} do
      acc ->
        [a, b] =
          String.trim(line)
          |> String.split("-")

        acc = Map.update(acc, a, MapSet.new([b]), &MapSet.put(&1, b))
        Map.update(acc, b, MapSet.new([a]), &MapSet.put(&1, a))
    end
  end

  def part1(file) do
    connections = parse(file)

    with_t =
      Map.filter(connections, fn {k, v} ->
        String.starts_with?(k, "t") || any_starts_with_t?(v)
      end)

    for {k, v} <- Map.to_list(with_t),
        [a, b] <- cartesian_prod(v) do
      if Enum.member?(Map.get(with_t, a, []), b) do
        MapSet.new([k, a, b])
      end
    end
    |> Enum.filter(&Function.identity/1)
    |> Enum.filter(&any_starts_with_t?/1)
    |> Enum.uniq()
    |> Enum.count()
  end

  def cartesian_prod(vals) do
    for a <- vals, b <- vals do
      [a, b]
    end
  end

  defp any_starts_with_t?(nodes), do: Enum.any?(nodes, &String.starts_with?(&1, "t"))

  def part2(file) do
    connections = parse_and_dedup(file)

    Stream.flat_map(Map.keys(connections), &search(connections, [&1], Enum.count(connections)))
    |> Enum.max_by(&Enum.count(&1))
    |> Enum.reverse()
    |> Enum.join(",")
  end

  def search(_, nodes, 0), do: [nodes]

  def search(connections, [node | nodes], n) do
    neighbors =
      Map.get(connections, node, MapSet.new())
      |> Enum.filter(fn candidate ->
        Enum.all?(nodes, &MapSet.member?(connections[&1], candidate))
      end)

    case neighbors do
      [] -> [[node | nodes]]
      candidates -> Enum.flat_map(candidates, &search(connections, [&1, node | nodes], n - 1))
    end
  end

  def parse_and_dedup(file) do
    File.stream!(file)
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.split(&1, "-"))
    |> Enum.map(fn [a, b] -> if a < b, do: {a, b}, else: {b, a} end)
    |> Enum.reduce(%{}, fn {a, b}, connections ->
      Map.update(connections, a, MapSet.new([b]), &MapSet.put(&1, b))
    end)
  end
end

IO.puts("Part 1")
[file] = System.argv()
Day23.part1(file) |> IO.inspect()
IO.puts("Part 2")
Day23.part2(file) |> IO.inspect()
