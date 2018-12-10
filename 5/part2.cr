chain = File.read("input5").strip

units = Set(Char).new
chain.each_char do |c|
    units.add(c.downcase)
end

shortest = chain.size
units.each do |c|
    this_chain = chain.delete(c).delete(c.upcase)
    i = 0
    while (i+1) < this_chain.size
        cur = this_chain[i]
        nxt = this_chain[i+1]

        if (cur.uppercase? && nxt == cur.downcase) || (cur.lowercase? && nxt == cur.upcase)
            this_chain = this_chain[0...i] + this_chain[i + 2..-1]
            # this_chain is shortened, so go back to previous and continue
            if i > 0
                i -= 1
            end
        else
            i += 1
        end
    end
    if this_chain.size < shortest
        shortest = this_chain.size
    end
end

puts shortest
