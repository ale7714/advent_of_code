modules = File.read("input.txt").split.map(&:to_i)

def fuel_requirement(mass)
  (mass/3).floor - 2
end

def recursive_fuel(mass)
  total = fuel_requirement(mass)
  return 0 if total <= 0
  total += recursive_fuel(total)
end

fuel = modules.sum { |module_mass| fuel_requirement(module_mass) }

puts "Answer part 1: #{fuel}"
puts "Answer part 2: #{modules.sum { |module_mass| recursive_fuel(module_mass) }}"

