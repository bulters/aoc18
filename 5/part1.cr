chain = File.read("input5").strip

i = 0
while (i+1) < chain.size
    cur = chain[i]
    nxt = chain[i+1]

    if (cur.uppercase? && nxt == cur.downcase) || (cur.lowercase? && nxt == cur.upcase)
        chain = chain[0...i] + chain[i + 2..-1]

        # chain is shortened, so go back to previous and continue
        if i > 0
            i -= 1
        end
    else
        i += 1
    end
end

puts chain.size
