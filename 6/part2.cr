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

dists = [] of Int32

max_x = input.max_by {|ex| ex.x }.x
max_y = input.max_by {|ey| ey.y }.y

0.upto max_y do |y|
  0.upto max_x do |x|
    c = Coord.new(x, y)
    ds = input.map{|ec| ec.manhattan(c) }
    dists << ds.sum
  end
end

puts dists.count {|d| d < 10000 }
