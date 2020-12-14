input = File.read("input.txt").split(/\n/).map(&:to_i)

def valid_number?(preamble, number)
  seen = {}
  valid = false

  preamble.each_with_index do |current, index|
    seen[current] = index
    tentative_result = number - current
    position = seen[tentative_result]
    valid = true unless position.nil?
    break if valid
  end

  valid
end

index = 25
invalid_number = nil
while index < input.size do
  preamble = input[index - 25..(index - 1)]
  number = input[index]
  unless valid_number?(preamble, number)
    invalid_number = number
    break
  end
  index += 1
end
puts "Day 9 part 1: #{invalid_number}"

list = input[0..index - 1]
numbers = nil
list.each_with_index do |outer, outer_index|
  acc = [outer]
  list[(outer_index + 1)..-1].each_with_index do |inner, innex_index|
    acc << inner
    break if acc.reduce(:+) >= invalid_number
  end
  if acc.reduce(:+) == invalid_number
    numbers = acc
    break
  end
end

puts "Day 9 part 2: #{numbers.min + numbers.max}"
