travel_map = File.read("input.txt").split(/\n/)

def traverse(tmap, rigth, down)
  trees = 0
  position =  0

  tmap.each_slice(down) do |slice|
    line = slice[0]
    trees += 1 if line[position] == "#"
    position = (position + rigth) % line.size
  end

  trees
end

s1_1 = traverse(travel_map, 1, 1)
s3_1 = traverse(travel_map, 3, 1)
s5_1 = traverse(travel_map, 5, 1)
s7_1 = traverse(travel_map, 7, 1)
s1_2 = traverse(travel_map, 1, 2)

puts "Part 1 answer: #{s3_1}"
puts "Part 2 answer: #{s1_1 * s3_1 * s5_1 * s7_1 * s1_2}"
