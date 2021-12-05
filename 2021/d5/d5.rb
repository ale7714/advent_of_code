def points s
  x1, y1, x2, y2 = s

  dx = x2 <=> x1
  dy = y2 <=> y1

  if dx == 0
    y1.step(y2, dy).map { |i| [x1, i] }
  elsif dy == 0
    x1.step(x2, dx).map { |i| [i, y1] }
  else
    x1.step(x2, dx).zip(y1.step(y2, dy))
  end
end

def count_intersections segments
  segments.flat_map { |s| points(s) }.tally.values.count { |i| i > 1 }
end


SEGMENTS = File.open("input.txt").read.lines.map { |l| l.scan(/\d+/).map &:to_i }

p count_intersections SEGMENTS.select { |x1, y1, x2, y2| x1 == x2 or y1 == y2 }
p count_intersections SEGMENTS