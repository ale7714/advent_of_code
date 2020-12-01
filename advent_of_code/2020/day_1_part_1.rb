expenses = File.read("input.txt").split.map(&:to_i)
# 2020 - (current) = result

seen = {}
result = nil

expenses.each_with_index do |current, index|
  seen[current] = index
  tentative_result = 2020 - current
  position = seen[tentative_result]
  result = current * expenses[position] unless position.nil?
  break unless result.nil?
end

puts result