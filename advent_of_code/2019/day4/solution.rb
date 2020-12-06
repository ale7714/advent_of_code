def valid?(password, matching2 = false)
  adjecent = {}
  no_decrease = true
  digits = password.digits.reverse

  digits.each_with_index do |digit, index|
    break if index + 1 == digits.size

    if digits[index + 1] == digit
      adjecent[digit] ||= 1
      adjecent[digit] += 1
    end

    if digit > digits[index + 1]
      no_decrease = false
      break
    end
  end

  if matching2
    adjecent.any? { |_, v| v == 2 } && no_decrease
  else
    adjecent.any? && no_decrease
  end
end

valid_passwords = (124075..580769).select do |password|
  valid? password
end.size

puts "Answer part 1: #{valid_passwords}"

valid_passwords = (124075..580769).select do |password|
  valid? password, true
end.size

puts "Answer part 2: #{valid_passwords}"