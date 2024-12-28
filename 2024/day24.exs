defmodule Day24 do
  defmodule Gate do
    defstruct [:l, :r, :op, :dest]
  end

  defp parse(file) do
    {wires, instructions} =
      File.stream!(file) |> Stream.map(&String.trim/1) |> Enum.split_while(&(&1 != ""))

    circuit =
      for wire <- wires, into: %{} do
        [name, value] = String.split(wire, ":")
        {name, String.to_integer(String.trim(value))}
      end

    gates =
      for i <- Enum.drop(instructions, 1) do
        [left, right] = String.split(i, "->")
        [l, op, r] = String.split(left)
        %Gate{l: l, op: String.to_atom(op), r: r, dest: String.trim(right)}
      end

    {circuit, gates}
  end

  def part1(file) do
    {circuit, gates} = parse(file)
    circuit = sim(circuit, gates)

    Map.keys(circuit)
    |> Enum.filter(&String.starts_with?(&1, "z"))
    |> Enum.sort()
    |> Enum.map(&to_string(circuit[&1]))
    |> Enum.join()
    |> String.reverse()
    |> String.to_integer(2)
  end

  defp sim(circuit, []), do: circuit

  defp sim(circuit, gates) do
    to_run = Enum.find(gates, &(Map.has_key?(circuit, &1.l) && Map.has_key?(circuit, &1.r)))

    circuit =
      case(to_run.op) do
        :AND -> do_op(circuit, to_run, &Bitwise.band/2)
        :OR -> do_op(circuit, to_run, &Bitwise.bor/2)
        :XOR -> do_op(circuit, to_run, &Bitwise.bxor/2)
      end

    rest = List.delete(gates, to_run)
    sim(circuit, rest)
  end

  defp do_op(circuit, gate, op) do
    l = circuit[gate.l]
    r = circuit[gate.r]
    res = op.(l, r)
    Map.put(circuit, gate.dest, res)
  end

  def part2(file) do
    {_, gates} = parse(file)
    z_gates = Enum.filter(gates, &String.starts_with?(&1.dest, "z"))

    connected_gates = find_connected(gates, z_gates, [])

    # All gates that output z should be an XOR (other than the last one)
    bad =
      Enum.filter(z_gates, fn gate ->
        gate.op != :XOR
      end)
      |> Enum.map(& &1.dest)
      |> Enum.drop(-1)

    # all XOR that come from OR or XOR should go to z gates
    badXOR =
      Enum.filter(connected_gates, &(&1.op == :XOR))
      |> Enum.filter(fn gate ->
        prev_l = prev_op(gates, gate.l)
        prev_r = prev_op(gates, gate.r)
        (prev_l == :XOR && prev_r == :OR) || (prev_l == :OR && prev_r == :XOR)
      end)
      |> Enum.map(& &1.dest)

    bad = bad ++ badXOR

    # Or gates after and gates
    bad_OR_l =
      Enum.filter(gates, &(&1.op == :OR))
      |> Enum.filter(fn gate ->
        prev_op(gates, gate.l) != :AND
      end)
      |> Enum.map(& &1.l)

    bad = bad ++ bad_OR_l

    bad_OR_r =
      Enum.filter(gates, &(&1.op == :OR))
      |> Enum.filter(fn gate ->
        prev_op(gates, gate.r) != :AND
      end)
      |> Enum.map(& &1.r)

    bad = bad ++ bad_OR_r

    bad_AND_l =
      Enum.filter(gates, &(&1.op == :AND))
      |> Enum.filter(fn gate ->
        prev = previous_gate(gates, gate.l)
        prev && prev.op == :AND && prev.l != "x00" && prev.r != "x00"
      end)
      |> Enum.map(& &1.l)

    bad = bad ++ bad_AND_l

    bad_AND_r =
      Enum.filter(gates, &(&1.op == :AND))
      |> Enum.filter(fn gate ->
        prev = previous_gate(gates, gate.r)
        prev && prev.op == :AND && prev.l != "x00" && prev.r != "x00"
      end)
      |> Enum.map(& &1.r)

    bad = bad ++ bad_AND_r

    # each OR should be connected to AND
    badOrs =
      Enum.filter(gates, fn gate ->
        gate.op == :OR && previous_gate(gates, gate.l).op != :AND
      end)
      |> Enum.map(& &1.l)

    bad = bad ++ badOrs

    badOrs =
      Enum.filter(gates, fn gate ->
        gate.op == :OR && previous_gate(gates, gate.r).op != :AND
      end)
      |> Enum.map(& &1.r)

    bad = bad ++ badOrs

    Enum.sort(bad)
    |> Enum.uniq()
    |> Enum.join(",")
  end

  defp find_connected(_, [], acc),
    do: acc |> Enum.uniq() |> Enum.reject(&String.starts_with?(&1.dest, "z"))

  defp find_connected(gates, [current | to_check], acc) do
    next_l = previous_gate(gates, current.l)
    next_r = previous_gate(gates, current.l)

    to_check =
      if next_l do
        [next_l | to_check]
      else
        to_check
      end

    to_check =
      if next_r do
        [next_r | to_check]
      else
        to_check
      end

    find_connected(gates, to_check, [current | acc])
  end

  defp prev_op(gates, code) do
    found = Enum.find(gates, &(&1.dest == code))

    if found do
      found.op
    else
      nil
    end
  end

  defp previous_gate(gates, code) do
    Enum.find(gates, &(&1.dest == code))
  end
end

[file] = System.argv()
IO.puts("Part 1")
Day24.part1(file) |> IO.inspect()
IO.puts("Part 2")
Day24.part2(file) |> IO.puts()
