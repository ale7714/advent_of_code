require 'set'
program = File.read("input.txt").split(/\n/)

executed = Set.new
accumulator = 0
current = 0

while true
  match = Regexp.new(/(\w{3})\s([+-]\d+)/).match(program[current])
  operation = match[1].strip
  argument = match[2].to_f

  break if executed.include? current

  case operation
  when "acc"
    accumulator += argument
    executed << current
    current += 1
  when "jmp"
    executed << current
    current += argument
  when "nop"
    current += 1
  end
end

puts "Answer part 1: #{accumulator}"

