elves = []

File.readlines(ARGV[0], "\n\n").each do |block|
    elves << block.rstrip.split("\n").map(&:to_i).sum
end

elves.sort!.reverse! 

puts elves[0]
puts elves.take(3).sum