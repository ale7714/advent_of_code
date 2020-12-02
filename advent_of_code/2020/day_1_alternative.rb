expenses = File.read("input.txt").split.map(&:to_i).sort!

def get_sum(values, start, target, pivot: nil)
  sum = nil
  final = -1

  while sum.nil? do
    break if values[start].nil? || values[final].nil?

    tentative_result = (pivot || 0) + values[start] + values[final]

    if tentative_result == target
      sum = (pivot || 1) * values[start] * values[final]
    elsif tentative_result > target
      final = final - 1
    elsif tentative_result < target
      start = start + 1
    end
  end

  return sum
end

puts "Part 1 answer: #{get_sum(expenses, 0, 2020)}"

answer_part_2 = expenses.each_with_index do |current_outer, index|
  result = get_sum(expenses, index, 2020, pivot: current_outer)
  break result unless result.nil?
end

puts "Part 2 answer: #{answer_part_2}"
