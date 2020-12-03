def valid_password?(first_pos, second_pos, letter, password)
  if password[first_pos] == letter && password[second_pos] != letter
    true
  elsif password[first_pos] != letter && password[second_pos] == letter
    true
  else
    false
  end
end

valid_passwords = 0

File.read("input.txt").split(/\n/).each do |line|
  policy, password = line.split(": ")
  range, letter = policy.split
  first_pos, second_pos = range.split("-").map { |i| i.to_i - 1 }

  valid_passwords+=1 if valid_password?(first_pos, second_pos, letter, password)
end

puts valid_passwords