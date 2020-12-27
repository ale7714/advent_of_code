class FerryGrid
  attr_accessor :rows

  def initialize
    @rows = []
  end

  def add_row(row)
    @rows << row
  end

  def self.parse(input)
    grid = FerryGrid.new
    input.split(/\n/).map(&:chars).each do |row|
      row.unshift('.')
      row << '.'
      grid.add_row(row)
    end
    grid.rows.unshift(['.'] * grid.rows[0].length)
    grid.add_row(['.'] * grid.rows[0].length)
    grid
  end

  def play(mode: 'visible')
    copy = FerryGrid.new
    width = rows.first.size - 1
    height = rows.size - 1
    rows.each_with_index do |row, y|
      new_row = row.dup
      row.each_with_index do |seat, x|
        case seat
        when 'L'
          new_row[x] = '#' if (mode == 'visible' ? visible_seats(y, x, width, height) == 0 : occupied_seats(y, x) == 0)
        when '#'
          new_row[x] = 'L' if (mode == 'visible' ? visible_seats(y, x, width, height) >= 5 : occupied_seats(y, x) >= 4)
        end
      end
      copy.add_row(new_row)
    end
    copy
  end

  def count_seats
    rows.sum { |row| row.sum{ |s| ( s == '#' ? 1 : 0 ) } }
  end

  def == grid
    self.rows == grid.rows
  end

  def to_s
    rows.map { |r| r.inspect }
  end

  private

  def occupied_seats(pos_y, pos_x)
    adjacent = 0
    neighbours = [[1, 0], [-1, 0], [0, 1], [0, -1], [1, 1], [1, -1], [-1, 1], [-1, -1]]
    neighbours.each do |dx, dy|
      x, y = pos_y, pos_x
      x += dx
      y += dy
      adjacent += 1 if rows[x][y] == "#"
    end
    adjacent
  end


  def visible_seats(pos_y, pos_x, x_max, y_max)
    occupied = 0
    directions = [[1, 0], [-1, 0], [0, 1], [0, -1], [1, 1], [1, -1], [-1, 1], [-1, -1]]
    directions.each do |dx, dy|
      x = pos_x + dx
      y = pos_y + dy
      while x >= 0 && x <= x_max && y >= 0 && y <= y_max do
        if rows[y][x] == "."
          x += dx
          y += dy
          next
        end
        occupied += rows[y][x] == "#" ? 1 : 0
        break
      end
    end
    occupied
  end
end

grid =  FerryGrid.parse(File.read("input.txt"))

loop do
  copy = grid.play(mode: 'default')
  break if grid == copy
  grid = copy
end

puts "Day 11 part 1: #{grid.count_seats}"

grid =  FerryGrid.parse(File.read("input.txt"))

loop do
  copy = grid.play(mode: 'visible')
  break if grid == copy
  grid = copy
end

puts "Day 11 part 2: #{grid.count_seats}"