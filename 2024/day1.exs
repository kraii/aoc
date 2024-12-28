getLeftAndRightLists = fn ->
  stream =
    File.stream!("test/day1.txt")
    |> Stream.map(&String.trim/1)
    |> Stream.map(fn line ->
      [l, r] = String.split(line)
      {String.to_integer(l), String.to_integer(r)}
    end)
    |> Enum.to_list()

  Enum.unzip(stream)
end

part1 = fn ->
  {ls, rs} = getLeftAndRightLists.()
  sortedLs = Enum.sort(ls)
  sortedRs = Enum.sort(rs)

  result =
    Enum.zip(sortedLs, sortedRs)
    |> Enum.map(fn {l, r} -> abs(l - r) end)
    |> Enum.sum()

  result
end

part2 = fn ->
  {ls, rs} = getLeftAndRightLists.()

  Enum.reduce(
    ls,
    0,
    fn l, acc ->
      c = l * Enum.count(rs, fn r -> r == l end)
      acc + c
    end
  )
end

IO.puts("part 1 - " <> Integer.to_string(part1.()))

IO.puts("part 2 - " <> Integer.to_string(part2.()))
