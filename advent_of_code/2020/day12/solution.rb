instructions = File.read("input.txt").split(/\n/)

position = {
  n: 0,
  s: 0,
  e: 0,
  w: 0,
  direction: :e,
}

DIRECTIONS = ['n', 's', 'e', 'w'].freeze
DEGREES = {
  e: 0,
  s: 1,
  w: 2,
  n: 3,
}

instructions.each do |i|
  match = Regexp.new("([NSEWLRF])(\\d+)").match(i)
  action = match[1].downcase
  value = match[2].to_i

  if DIRECTIONS.include? action
    position[action.to_sym] += value
    next
  end

  if action == 'f'
    position[position[:direction]] += value
    next
  end

  if action == 'r'
    current = DEGREES[position[:direction]]
    new_direction = (current + value/90)%4
    position[:direction] = DEGREES.key(new_direction)
  else
    current = DEGREES[position[:direction]]
    new_direction = (current - value/90)%4
    position[:direction] = DEGREES.key(new_direction)
  end
end

m_distance = (position[:s] - position[:n]).abs + (position[:w] - position[:e]).abs
puts "Day 12 part 1: #{m_distance}"