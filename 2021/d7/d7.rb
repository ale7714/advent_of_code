def get_cost(positions)
    fuel = 0
    positions.each do |position|
        fuel += yield position
    end
    fuel
end

POS = File.open("input.txt").read.split(",").map(&:to_i)

fuels = []
fuels_mean = []

(POS.min..POS.max).each do |target|
    fuels << get_cost(POS) { |p| (p - target).abs }
    fuels_mean << get_cost(POS) { |p| ((p - target).abs * ((p - target).abs + 1))/2 }
end

p fuels.min
p fuels_mean.min


