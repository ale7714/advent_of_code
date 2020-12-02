expenses = File.read("input.txt").split.map(&:to_i)
# 2020 - (current_outer_loop + current_inner_loop) = result

result = nil

expenses.each_with_index do |current_outer, index_outer|
  seen = {}
  expenses[index_outer..-1].each_with_index do |current_inner, index_inner|
    seen[current_inner] = index_inner + index_outer
    tentative_result = 2020 - (current_outer + current_inner)
    position = seen[tentative_result]
    result = current_outer * current_inner * expenses[position] unless position.nil?
    break unless result.nil?
  end
end

puts result