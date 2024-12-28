defmodule Memo do
  use Agent

  def start_link() do
    Agent.start_link(fn -> %{} end, name: __MODULE__)
  end

  def get(k) do
    Agent.get(__MODULE__, fn memo -> memo[k] end)
  end

  def put(k, v) do
    Agent.update(__MODULE__, fn memo -> Map.put(memo, k, v) end)
  end
end

defmodule Day19 do
  defp parse(file) do
    [patterns | designs] = File.stream!(file) |> Enum.map(&String.trim/1) |> Enum.to_list()
    ps = String.split(patterns, ",") |> Enum.map(&String.trim/1)
    {ps, Enum.drop(designs, 1)}
  end

  def part1(file) do
    {patterns, designs} = parse(file)

    for design <- designs, reduce: 0 do
      acc ->
        if(make_design(design, patterns) > 0) do
          acc + 1
        else
          acc
        end
    end
  end

  def part2(file) do
    {patterns, designs} = parse(file)

    for design <- designs, reduce: 0 do
      acc ->
        acc + make_design(design, patterns)
    end
  end

  defp make_design("", _), do: 1

  defp make_design(design, patterns) do
    cached = Memo.get(design)

    if(cached != nil) do
      cached
    else
      res =
        Enum.filter(patterns, &String.starts_with?(design, &1))
        |> Enum.map(&make_design(String.replace_prefix(design, &1, ""), patterns))
        |> Enum.sum()

      Memo.put(design, res)
      res
    end
  end
end

Memo.start_link()
[file] = System.argv()
IO.puts("Part 1")
Day19.part1(file) |> IO.inspect()
IO.puts("Part 2")
Day19.part2(file) |> IO.inspect()
