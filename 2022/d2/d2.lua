local seq = require("pl.seq")
local tablex = require("pl.tablex")

local function score_choice(choice)
    local choices = {
        R = 1,
        P = 2,
        S = 3
    }
    return choices[choice]
end

local function translate(choice)
    local choices = {
        X = 'R',
        Y = 'P',
        Z = 'S',
        A = 'R',
        B = 'P',
        C = 'S'
    }
    return choices[choice]
end

local function score_outcome(opponent, me)
    local choices = {
        R = {
            R = 3,
            P = 0,
            S = 6
        },
        P = {
            P = 3,
            S = 0
        },
        S = {
            S = 3
        }
    }

    if choices[me][opponent] == nil then
        return 6 - choices[opponent][me]
    else
        return choices[me][opponent]
    end
end

local function get_score(opponent, me)
    local score = 0
    for o, m in seq.zip(opponent, me) do
        score = score + score_outcome(o, m) + score_choice(m)
    end
    return score
end

local function get_data()
    local opponent, me = {}, {}
    for line in io.lines('input') do
        opponent[#opponent+1] = string.sub(line, 1, 1)
        me[#me+1] = string.sub(line, 3, 3)
    end
    return opponent, me
end

local function translate_me_p2(opponent, me)
    local choices = {
        -- Lose
        X = {
            R = 'S',
            P = 'R',
            S = 'P'
        },
        -- Tie
        Y = {
            R = 'R',
            P = 'P',
            S = 'S'
        },
        -- Win
        Z = {
            R = 'P',
            P = 'S',
            S = 'R'
        }
    }
    return choices[me][opponent]
end

-- part 1
local opponent, me = get_data()
local rps_opponent = tablex.map(translate, opponent)
local rps_me = tablex.map(translate, me)
print(get_score(rps_opponent, rps_me))


-- part 2
local p2_me = {}
for o, m in seq.zip(rps_opponent, me) do
    p2_me[#p2_me+1] = translate_me_p2(o, m)
end
print(get_score(rps_opponent, p2_me))
