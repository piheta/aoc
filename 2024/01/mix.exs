defmodule FindDiff do
  def get_sorted_arrays(input_text) do
    {lArr, rArr} =
      Enum.reduce(input_text, {[], []}, fn line, {left_acc, right_acc} ->
        case String.split(line, "   ") do
          [left, right] ->
            {leftint, _} = Integer.parse(left)
            {rightint, _} = Integer.parse(right)
            {[leftint | left_acc], [rightint | right_acc]}

          _ ->
            {left_acc, right_acc}
        end
      end)

    lArr = Enum.sort(lArr)
    rArr = Enum.sort(rArr)
    {lArr, rArr}
  end

  def solve do
    lines =
      case File.read("input.txt") do
        {:ok, content} -> content
        {:error, reason} -> "Error: #{reason}"
      end
      |> String.split("\n", trim: true)

    {lArr, rArr} = get_sorted_arrays(lines)

    diff_sum =
      Enum.zip(lArr, rArr)
      |> Enum.map(fn sum -> abs(elem(sum, 0) - elem(sum, 1)) end)
      |> Enum.sum()

    sim_sum =
      Enum.map(lArr, fn e -> e * Enum.count(rArr, fn x -> x == e end) end)
      |> Enum.sum()

    IO.inspect(diff_sum, label: "diff_sum")
    IO.inspect(sim_sum, label: "sim_sum")
  end
end

FindDiff.solve()
