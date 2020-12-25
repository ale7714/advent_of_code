adapters = File.read("input.txt").split(/\n/).map(&:to_i).sort

start = 0
differences = {}
differences[1] = 0
differences[3] = 1

adapters.each do |adapter|
  differences[adapter - start]+=1
  start =  adapter
end

puts "Day 10 part 1: #{differences[1] * differences[3]}"

def find_arrangement(list, arrangements, pos: 0)
  current = list[pos]

  return arrangements[current] if arrangements[current] > 0

  if pos == list.size - 1
    arrangements[current] += 1
    return 1
  else
    possible = 0
    list[pos+1..-1].each_with_index do |other, index|
      break if (other - current) > 3

      possible += find_arrangement(list,arrangements, pos: (pos + 1 + index))
    end
    arrangements[current] = possible
    return possible
  end
end

arrangements = Hash.new(0)
adapters.unshift(0)
adapters << adapters.last + 3
find_arrangement(adapters, arrangements)
puts "Day 10 part 2: #{arrangements[0]}"