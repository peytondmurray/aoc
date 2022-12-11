local seq = require("pl.seq")
local function get_data()
    local c1, c2 = {}, {}
    for line in io.lines('input') do
        local len = string.len(line)
        c1[#c1+1] = string.sub(line, 1, len/2)
        c2[#c2+1] = string.sub(line, len/2 + 1, len)
    end
    return c1, c2
end

local function pprint(t1, t2)
    for v1, v2 in seq.zip(t1, t2) do
        print(v1 .. '   ' .. v2)
    end
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
        chars[s:sub(i, i)] = true
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

-- p1
local c1, c2 = get_data()
local dups = find_duplicated(c1, c2)

for i, v in ipairs(dups) do
    print(i, v)
end
