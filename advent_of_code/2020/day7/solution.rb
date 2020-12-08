require 'set'

class Bag
  attr_reader :color, :content, :parents

  def initialize(color)
    @color = color
    @content = {}
    @parents = Set.new
  end

  def empty?
    content.empty?
  end

  def total_bags
    content.sum { |_, quantity| quantity }
  end
end

class Bags
  attr_reader :nodes

  def add_from_rule(input)
    match = Regexp.new(/(.+)bags\scontain\s(.+)\./).match(input)
    color = match[1].strip
    content = match[2]

    bag = nodes[color] || Bag.new(color)

    unless content == "no other bags"
      content.split(", ").each do |content_bag|
        quantity = content_bag[0]
        content_color = Regexp.new(/\d{1}\s(.*)\sbags?/).match(content_bag)[1].strip
        kid =  nodes[content_color] || Bag.new(content_color)

        bag.content[content_color] = quantity.to_i
        kid.parents << color
        nodes[kid.color] = kid
      end
    end

    nodes[color] = bag
  end

  def initialize
    @nodes = {}
  end

  def get_parents(color, parents)
    return if nodes[color].parents.empty?
    nodes[color].parents.each { |node| parents.add node; get_parents(node, parents) }
    parents
  end

  def total_bags(color, count)
    return 0 if nodes[color].empty?
    count = nodes[color].total_bags
    nodes[color].content.each { |node, quantity| count += quantity * total_bags(node, count) }
    count
  end
end


rules = File.read("input.txt").split(/\n/)

bags = Bags.new

rules.each do |rule|
  bags.add_from_rule(rule)
end

puts "Answer part 1: #{bags.get_parents("shiny gold", Set.new).count}"
puts "Answer part 2: #{bags.total_bags("shiny gold", 0)}"
