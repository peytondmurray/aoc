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
    local low = tonumber(rangestr:sub(1, dash-1))
    local high = tonumber(rangestr:sub(dash+1))
    return {low, high}
end

local function parse(data)
    local elves = {}
    for _, elfpair in pairs(data) do
        local comma = elfpair:find(',')
        if comma then
            local left = get_range(elfpair:sub(1, comma-1))
            local right = get_range(elfpair:sub(comma+1))
            elves[#elves+1] = {left, right}
        end
    end
    return elves
end

local function does_pair_overlap(elfpair)
    return (
        elfpair[1][1] <= elfpair[2][1] and elfpair[1][2] >= elfpair[2][1]
    ) or (
        elfpair[1][1] >= elfpair[2][1] and elfpair[1][1] <= elfpair[2][2]
    )
end

local function is_pair_contained(elfpair)
    return (
        elfpair[1][1] >= elfpair[2][1] and elfpair[1][2] <= elfpair[2][2]
    ) or (
        elfpair[1][1] <= elfpair[2][1] and elfpair[1][2] >= elfpair[2][2]
    )
end

local function eleves_overlap(data, overlap_func)
    local n_overlaps = 0
    local elves = parse(data)
    for _, elfpair in pairs(elves) do
        if overlap_func(elfpair) then
            n_overlaps = n_overlaps+1
        end
    end
    return n_overlaps
end


-- p1
local data = get_data()
print(eleves_overlap(data, is_pair_contained))

-- p2
print(eleves_overlap(data, does_pair_overlap))
