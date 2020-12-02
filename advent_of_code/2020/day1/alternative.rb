expenses = File.read("input.txt").split.map(&:to_i).sort!

def get_pair(values, start, target)
  pair = nil
  final = -1

  while pair.nil? do
    break if values[start].nil? || values[final].nil?

    tentative_result = values[start] + values[final]

    if tentative_result == target
      pair =  [values[start], values[final]]
    elsif tentative_result > target
      final = final - 1
    elsif tentative_result < target
      start = start + 1
    end
  end

  return pair
end

puts "Part 1 answer: #{get_pair(expenses, 0, 2020).inject(:*)}"

answer_part_2 = expenses.each_with_index do |current_outer, index|
  pair = get_pair(expenses, index + 1, 2020 - current_outer)

  break pair.inject(:*) * current_outer unless pair.nil?
end

puts "Part 2 answer: #{answer_part_2}"
