defmodule Day17 do
  defmodule Computer do
    defstruct [:a, :b, :c, ip: 0, output: []]
  end

  defp parse(file) do
    lines = File.stream!(file) |> Enum.map(&String.trim/1) |> Enum.to_list()

    [a, b, c, _, p] = lines

    computer = %Computer{
      a: int(value(a)),
      b: int(value(b)),
      c: int(value(c))
    }

    codes = String.split(value(p), ",") |> Enum.map(&int/1)
    {computer, codes}
  end

  defp int(s), do: String.to_integer(s)

  defp value(line) do
    String.split(line, ":") |> Enum.at(1) |> String.trim()
  end

  def part1(file) do
    {computer, codes} = parse(file)
    computer = evaluate(computer, codes)
    Enum.reverse(computer.output) |> Enum.join(",")
  end

  # Halt
  defp evaluate(computer, codes) when computer.ip >= length(codes), do: computer

  defp evaluate(computer, codes) do
    [op, operand] = Enum.slice(codes, computer.ip..(computer.ip + 1))

    computer =
      case(op) do
        0 -> adv(computer, operand) |> advance_ip()
        1 -> bxl(computer, operand) |> advance_ip()
        2 -> bst(computer, operand) |> advance_ip()
        3 -> jnz(computer, operand)
        4 -> bxc(computer) |> advance_ip()
        5 -> out(computer, operand) |> advance_ip()
        6 -> bdv(computer, operand) |> advance_ip()
        7 -> cdv(computer, operand) |> advance_ip()
      end

    evaluate(computer, codes)
  end

  defp eval_combo(computer, combo) do
    case combo do
      4 -> computer.a
      5 -> computer.b
      6 -> computer.c
      7 -> raise("777")
      _ -> combo
    end
  end

  defp adv(computer, operand) do
    d = Integer.pow(2, eval_combo(computer, operand))
    %{computer | a: Integer.floor_div(computer.a, d)}
  end

  defp bxl(computer, operand) do
    result = Bitwise.bxor(computer.b, operand)
    %{computer | b: result}
  end

  defp bst(computer, operand) do
    result = Integer.mod(eval_combo(computer, operand), 8)
    %{computer | b: result}
  end

  defp jnz(computer, operand) do
    if computer.a == 0 do
      advance_ip(computer)
    else
      %{computer | ip: operand}
    end
  end

  defp bxc(computer) do
    result = Bitwise.bxor(computer.b, computer.c)
    %{computer | b: result}
  end

  defp out(computer, operand) do
    result = Integer.mod(eval_combo(computer, operand), 8)
    %{computer | output: [result | computer.output]}
  end

  defp bdv(computer, operand) do
    d = Integer.pow(2, eval_combo(computer, operand))
    %{computer | b: Integer.floor_div(computer.a, d)}
  end

  defp cdv(computer, operand) do
    d = Integer.pow(2, eval_combo(computer, operand))
    %{computer | c: Integer.floor_div(computer.a, d)}
  end

  defp advance_ip(computer), do: %{computer | ip: computer.ip + 2}

  def part2(file) do
    {computer, codes} = parse(file)

    # The program just loops dividing by at so we can solve subsets and multiply by
    find_magic_A(computer, codes, codes)
  end

  defp find_magic_A(computer, codes, target) do
    a =
      if(length(target) == 1) do
        0
      else
        8 * find_magic_A(computer, codes, tl(target))
      end



    Stream.iterate(a, &(&1 + 1))
    |> Enum.find(fn a ->
      result = Enum.reverse(evaluate(%{computer | a: a}, codes).output)
      result == target
    end)
  end
end

file = hd(System.argv())
IO.puts("part 1")
Day17.part1(file) |> IO.inspect()
IO.puts("part 2")
Day17.part2(file) |> IO.inspect()
