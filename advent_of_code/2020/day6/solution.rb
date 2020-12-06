def get_sum(groups, operator)
  groups.sum do |group|
    group.split(/\n/).map(&:chars).inject(operator.to_sym).size
  end
end

groups = File.read("input.txt").split(/\n{2,}/)

puts "Answer part 1: #{get_sum(groups, "|")}"
puts "Answer part 2: #{get_sum(groups, "&")}"
