defmodule AoC do
  def get_reports(input_text) do
    reports =
      Enum.reduce(input_text, [], fn line, acc ->
        parts = String.split(line, " ")
        [parts | acc]
      end)
      |> Enum.reverse()

    reports
  end

  def solve do
    lines =
      case File.read("input.txt") do
        {:ok, content} -> content
        {:error, reason} -> "Error: #{reason}"
      end
      |> String.split("\n", trim: true)

    reports = get_reports(lines)

    IO.inspect(reports)
    # diff_sum =
    #   Enum.zip(lArr, rArr)
    #   |> Enum.map(fn sum -> abs(elem(sum, 0) - elem(sum, 1)) end)
    #   |> Enum.sum()

    # sim_sum =
    #   Enum.map(lArr, fn e -> e * Enum.count(rArr, fn x -> x == e end) end)
    #   |> Enum.sum()
  end
end

AoC.solve()
