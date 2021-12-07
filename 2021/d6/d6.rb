def count_laternfish day_count, days
    days.times { |_| day_count.rotate!; day_count[6] += day_count.last }
    day_count.sum
end


LFS = File.open("input.txt").read.split(",").map(&:to_i)
lf_day_count = (0..8).map { |d| LFS.count(d) || 0 }

p count_laternfish lf_day_count.dup, 80
p count_laternfish lf_day_count, 256
