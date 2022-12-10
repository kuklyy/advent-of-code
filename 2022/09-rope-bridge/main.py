def dir_key(x, y):
    return f"{x},{y}"


mod = {
    "R": [1, 0],
    "D": [0, -1],
    "L": [-1, 0],
    "U": [0, 1],
    "LU": [-1, 1],
    "RU": [1, 1],
    "LD": [-1, -1],
    "RD": [1, -1],
    "C": [0, 0],
}


def should_move_tail(hx, hy, tx, ty):
    for i in mod.values():
        if tx == hx + i[0] and ty == hy + i[1]:
            return False

    return True


def part1():
    t_pos = {
        "0,0": 1,
    }
    hx, hy, tx, ty = 0, 0, 0, 0
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            d, v = line.split(' ')

            for _ in range(0, int(v)):
                hx += mod[d][0]
                hy += mod[d][1]

                if should_move_tail(hx, hy, tx, ty):
                    tx = hx + mod[d][0] * -1
                    ty = hy + mod[d][1] * -1
                    t_pos[dir_key(tx, ty)] = 1

    s = 0
    for k in t_pos.keys():
        s += 1
    print(s)


def part2():
    t_pos = {
        "0,0": 1,
    }
    h = [0, 0]
    knots_len = 9
    knots = []
    for _ in range(0, knots_len):
        knots.append([0, 0])

    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            d, v = line.split(' ')

            for _ in range(0, int(v)):
                h[0] += mod[d][0]
                h[1] += mod[d][1]
                p = h
                for i in range(0, knots_len):
                    if should_move_tail(p[0], p[1], knots[i][0], knots[i][1]):
                        nx = p[0] - knots[i][0]
                        ny = p[1] - knots[i][1]
                        if nx < -1:
                            nx = -1
                        elif nx > 1:
                            nx = 1
                        if ny < -1:
                            ny = -1
                        elif ny > 1:
                            ny = 1
                        knots[i][0] += nx
                        knots[i][1] += ny
                        if i == knots_len-1:
                            t_pos[dir_key(knots[i][0], knots[i][1])] = 1
                    p = knots[i]

    s = 0
    for k in t_pos.keys():
        s += 1
    print(s)


part1()
part2()
