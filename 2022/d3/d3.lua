local seq = require("pl.seq")
local function split_compartments(table)
    local c1, c2 = {}, {}
    for _, items in pairs(table) do
        local len = string.len(items)
        c1[#c1+1] = string.sub(items, 1, len/2)
        c2[#c2+1] = string.sub(items, len/2 + 1, len)
    end
    return c1, c2
end

local function get_data()
    local lines = {}
    for line in io.lines('input') do
        lines[#lines+1] = line
    end
    return lines
end

local function get_value(c)
    local val = string.byte(c)
    if 97 <= val and val <= 122 then
        -- lower case
        return val - 96
    elseif 65 <= val and val <= 90 then
        -- upper case
        return val - 38
    end
end

local function str_to_table(s)
    local chars = {}
    for i = 1, #s do
        chars[s:sub(i, i)] = i
    end
    return chars
end

local function find_in_table(chars, table)
    for i = 1, #chars do
        local char = chars:sub(i, i)
        if table[char] then
            return char
        end
    end
end

local function find_duplicated(c1, c2)
    local duplicated = {}
    for s1, s2 in seq.zip(c1, c2) do
        duplicated[#duplicated+1] = find_in_table(s2, str_to_table(s1))
    end
    return duplicated
end

local function get_sum_value(table)
    local sum = 0
    for _, v in ipairs(table) do
        sum = sum + get_value(v)
    end
    return sum
end

-- Make groups from the list of backpack items
local function get_groups(items)
    local groups = {}
    for i = 1, #items, 3 do
        groups[#groups+1] = {
            items[i],
            items[i+1],
            items[i+2],
        }
    end
    return groups
end

local function get_shared_char(group)
    local chars = str_to_table(group[1])

    -- Iterate through both lists, keeping only chars that appear twice
    for i = 2, 3  do
        local buf = str_to_table(group[i])
        local keep = {}
        for char in pairs(chars) do
            if buf[char] then
                keep[char] = buf[char]
            end
        end
        chars = keep
    end

    -- There's no way to get table values by numerical index if you've got a key/value table...
    for char in pairs(chars) do
        return char
    end
end

local function sum_shared_chars(groups)
    local sum = 0
    for _, group in pairs(groups) do
        sum = sum + get_value(get_shared_char(group))
    end
    return sum
end

-- p1
local data = get_data()
local c1, c2 = split_compartments(data)
print(get_sum_value(find_duplicated(c1, c2)))

-- p2
local groups = get_groups(data)
print(sum_shared_chars(groups))
