def bin_search(input, lower_key, startpos, endpos)
  if input.size == 1
    return input == lower_key ? startpos : endpos
  end

  middle = ((endpos + startpos)/2).to_i

  if input[0] == lower_key
    bin_search(input[1..-1], lower_key, startpos, middle)
  else
    bin_search(input[1..-1], lower_key, middle + 1, endpos)
  end
end

passes = File.read("input.txt").split(/\n/)

seat_ids = passes.map do |pass|
  row = bin_search(pass[0..6], "F", 0, 127)
  column = bin_search(pass[7..9], "L", 0, 7)
  row * 8 + column
end.sort

puts "Answer part 1: #{seat_ids.max}"

seat_id = seat_ids.each_with_index do |current, index|
  missing = current + 1
  break missing unless missing == seat_ids[index + 1]
end

puts "Answer part 2: #{seat_id}"

