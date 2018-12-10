struct Coord
    getter x : Int32
    getter y : Int32

    def initialize(@x, @y)
    end

    def manhattan(to : Coord)
      (x - to.x).abs + (y - to.y).abs
    end
end

input = File.read_lines("input6").map do |l| 
    xy = l.split(", ")
    Coord.new(xy[0].to_i, xy[1].to_i)
end

closests = Hash(Coord, Set(Coord)).new {|a, k| a[k] = Set(Coord).new }

max_x = input.max_by {|ex| ex.x }.x
max_y = input.max_by {|ey| ey.y }.y
puts max_x, max_y
0.upto max_y do |y|
  0.upto max_x do |x|
    c = Coord.new(x, y)
    dists = input.map {|ce| {ce, c.manhattan(ce)}}
    closest, dist = dists.min_by {|e| e[1]}
    closests[closest] << c unless dists.count {|(_, d)| d == dist} > 1
  end
end

totals = closests.reject do |_, cs|
  cs.any? {|c| c.x == 0 || c.y == 0 || c.x == max_x || c.y == max_y }
end.map do |(_, ds)|
  ds.size
end

puts totals.max
