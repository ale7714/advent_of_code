program = File.read("input.txt").strip.split(",").map(&:to_i)

def get_output(noun, verb, ops)
  ops = ops.dup
  ops[1] = noun
  ops[2] = verb

  position = 0

  while position < ops.size
    op = ops[position]

    if op == 1
      sum = ops[ops[position+1]] + ops[ops[position+2]]
      ops[ops[position+3]] = sum
    elsif op == 2
      mul = ops[ops[position+1]] * ops[ops[position+2]]
      ops[ops[position+3]] = mul
    else
      break
    end

    position += 4
  end
  ops[0]
end

puts "Part 1 answer: #{get_output(12, 2, program)}"

output_verb = nil
output_noun = nil

(0..99).each do |noun|
  (0..99).each do |verb|
    if get_output(noun, verb, program) == 19690720
      output_verb = verb
      break
    end
  end

  if output_verb
    output_noun = noun
    break
  end
end

puts "Part 2 answer: #{100 * output_noun + output_verb}"