def valid_password?(min, max, letter, password)
  ocurrences = (password.chars.group_by(&:itself)[letter] || []).size

  min <= ocurrences && ocurrences <= max ? true : false
end

valid_passwords = 0

File.read("input.txt").split(/\n/).each do |line|
  policy, password = line.split(": ")
  range, letter = policy.split
  min, max = range.split("-").map(&:to_i)

  valid_passwords+=1 if valid_password?(min, max, letter, password)
end

puts valid_passwords