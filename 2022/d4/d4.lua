local function get_data()
    local lines = {}
    for line in io.lines('input') do
        lines[#lines+1] = line
    end
    return lines
end

local function get_range(rangestr)
    -- String range: 1-3
    local dash = rangestr:find('-')
    return {tonumber(rangestr:sub(1, dash)), tonumber(rangestr:sub(dash+1))}
end

local function parse(data)
    local elves = {}
    for _, elfpair in pairs(data) do
        local comma = elfpair:find(',')
        if comma then
            elves[#elves+1] = {
                {
                    get_range(elfpair:sub(1, comma-1)),
                    get_range(elfpair:sub(comma+1))
                }
            }
        end
    end
    return elves
end

local function does_pair_overlap(elfpair)
    return (
        elfpair[1][1] >= elfpair[2][1] and elfpair[1][1] <= elfpair[2][2]
    ) or (
        elfpair[1][2] >= elfpair[2][1] and elfpair[1][2] <= elfpair[2][2]
    )
end

local function is_pair_contained(elfpair)
    return (
        elfpair[1][1] >= elfpair[2][1] and elfpair[1][2] <= elfpair[2][2]
    ) or (
        elfpair[1][1] <= elfpair[2][1] and elfpair[1][2] >= elfpair[2][2]
    )
end

local function eleves_overlap(data)
    local n_overlaps = 0
    local elves = parse(data)
    for _, elfpair in pairs(elves) do
        print(elfpair[1], elfpair[2])
        -- print('[' .. elfpair[1][1] .. ' ' .. elfpair[1][2] .. '] [' .. elfpair[2][1] .. ' ' .. elfpair[2][2] .. ']')
        -- if is_pair_contained(elfpair) then
        --     n_overlaps = n_overlaps+1
        -- end
    end
    return n_overlaps
end


-- p1
local data = get_data()
print(eleves_overlap(data))
