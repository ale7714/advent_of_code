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
  element = list[pos]
  if arrangements[element] > 0
    return arrangements[element]
  elsif pos == list.size - 1
    arrangements[element] += 1
    return 1
  else
    possible = 0
    list[pos+1..-1].each_with_index do |other, index|
      if (other - element) <= 3
        possible += find_arrangement(list,arrangements, pos: (pos + 1 + index))
      else
        break
      end
    end
    arrangements[element] = possible
    return possible
  end
end

arrangements = Hash.new(0)
adapters.unshift(0)
adapters << adapters.last + 3
find_arrangement(adapters, arrangements)
puts puts "Day 10 part 2: #{arrangements[0]}"