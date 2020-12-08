class Instruction
  attr_reader :operation, :argument

  def initialize(operation, argument)
    @operation = operation
    @argument = argument
  end

  def swapable?
    @swapable ||= ["nop", "jmp"].include? operation
  end

  def swap
    return self unless swapable?
    new_operation = (operation == "nop" ? "jmp" : "nop")
    Instruction.new(new_operation, argument)
  end

  def self.parse(input)
    match = Regexp.new(/(\w{3})\s([+-]\d+)/).match(input)
    self.new(match[1].strip, match[2].to_f)
  end
end

class Program
  class EarlyTerminationError < StandardError; end

  attr_reader :instructions

  def initialize
    @instructions = []
  end

  def add(instruction)
    instructions << instruction
  end

  def get_acc(complete = false)
    executed = []
    accumulator = 0
    current = 0

    while true
      break if instructions[current].nil? || executed.include?(current)
      operation = instructions[current].operation
      argument = instructions[current].argument

      case operation
      when "acc"
        accumulator += argument
        executed << current
        current += 1
      when "jmp"
        executed << current
        current += argument
      when "nop"
        executed << current
        current += 1
      end
    end

    raise EarlyTerminationError if complete && !executed.include?(instructions.size - 1)

    accumulator
  end

  def swap(index)
    new_instructions = instructions.dup
    new_instructions[index] = new_instructions[index].swap
    program = Program.new
    new_instructions.each { |i| program.add i }
    program
  end

  def swapable_index
    swapable = []
    instructions.each_with_index { |inst, index| swapable << index if inst.swapable?}
    swapable
  end

  def self.parse(input)
    program = self.new
    input.each do |inst|
      program.add(Instruction.parse(inst))
    end
    program
  end
end


program = Program.parse(File.read("input.txt").split(/\n/))
puts "Answer part 1: #{program.get_acc}"

accumulator = nil
program.swapable_index.each do |index|
  begin
    new_program = program.swap(index)
    accumulator = new_program.get_acc(true)
    if accumulator
      break
    end
  rescue Program::EarlyTerminationError
      next
  end
end

puts "Answer part 2: #{accumulator}"