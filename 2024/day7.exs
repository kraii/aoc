defmodule Day7 do
  defp parse_terms() do
    File.stream!("test/day7.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.map(fn line ->
      [l, r] = String.split(line, ":")
      rs = String.split(r) |> Enum.map(&String.to_integer/1)
      {String.to_integer(l), rs}
    end)
  end

  def solve(available_ops) do
    terms = parse_terms()

    for {l, rs} <- terms, reduce: 0 do
      acc ->
        case possible_to_calc?(l, rs, available_ops) do
          true -> acc + l
          false -> acc
        end
    end
  end

  defp possible_to_calc?(l, rs, available_ops), do: possible?(l, Enum.reverse(rs), available_ops)

  defp possible?(l, [r], _available_ops), do: l == r

  defp possible?(l, [r | rs], available_ops) do
    Enum.reject(available_ops, fn op -> !possible_op?(op, l, r) end)
    |> Enum.any?(fn op -> possible?(reverse_apply(op, l, r), rs, available_ops) end)
  end

  defp possible_op?(op, l, r) do
    case(op) do
      :+ -> l > r
      :* -> l > 0 && rem(l, r) == 0
      :|| -> l > r && String.ends_with?(to_string(l), to_string(r))
    end
  end

  defp reverse_apply(op, l, r) do
    case(op) do
      :+ ->
        l - r

      :* ->
        Integer.floor_div(l, r)

      :|| ->
        String.replace_suffix(to_string(l), to_string(r), "") |> String.to_integer()
    end
  end
end

IO.puts("Part 1")
Day7.solve([:+, :*]) |> IO.inspect()

IO.puts("Part 2")
Day7.solve([:+, :*, :||]) |> IO.inspect()
