local function print_table(t)
    for _, v in ipairs(t) do
        print('Elf=' .. v.number)
        for j, entry in ipairs(v.meals) do
            print('  ' .. entry)
        end
        print('  total=' .. v.total)
    end
end

local function get_top_3_total(t)
    print("Top 3 elves:")
    table.sort(t, function(a, b) return a.total < b.total end)

    local total = 0
    for j = 0,2,1 do
        print('  ' .. t[#t - j].number, t[#t - j].total)
        total = total + t[#t - j].total
    end
    print('Total of the top 3 elves=' .. total)
end


local function get_elf_table()
    local elves = {}
    local i = 1
    table.insert(elves, {number =  1, meals = {}, total = 0})
    for line in io.lines('input') do
        local num = tonumber(line)

        if num == nil then
            table.insert(elves, {number = i, meals = {}, total = 0})
        else
            -- Insert the number into the meals table for the latest elf
            table.insert(elves[#elves].meals, num)
            elves[#elves].total = elves[#elves].total + num
        end
        i = i + 1
    end
    return elves
end

local elves = get_elf_table()

-- Part 1
print_table(elves)

-- Part 2
get_top_3_total(elves)
