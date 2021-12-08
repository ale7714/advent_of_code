UNIQ_DIGITS = {
    2 => 1,
    3 => 7,
    4 => 4,
    7 =>  8
}.freeze
signals = []
outputs = []
File.open("input.txt").read.lines.each do |l| 
    s, o = l.split(" | ").map(&:split)
    signals.push(*s)    
    outputs << o 
end

p outputs.flatten.select { |o| [2,3,4,7].include? o.size  }.count 