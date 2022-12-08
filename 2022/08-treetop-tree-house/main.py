def check_sibling(si, c):
    for s in si:
        if s >= c:
            return False

    return True


def part1():
    d = []
    result = 0
    with open('input.txt', 'r') as reader:
        for li in reader:
            cd = []
            line = li.replace('\n', '')
            for c in line:
                cd.append(int(c))
            d.append(cd)

    for y in range(1, len(d)-1):
        for x in range(1, len(d)-1):
            c = d[y][x]

            l = d[y][0:x]
            r = d[y][x+1:len(d)]
            t = []
            b = []
            for i in range(0, y):
                t.append(d[i][x])
            for i in range(y+1, len(d)):
                b.append(d[i][x])

            if check_sibling(l, c):
                result += 1
                continue
            if check_sibling(r, c):
                result += 1
                continue
            if check_sibling(t, c):
                result += 1
                continue
            if check_sibling(b, c):
                result += 1
                continue

    result += len(d) * 2 + (len(d)-2)*2
    print(result)


def sibling_score(si, c):
    score = 0
    for s in si:
        score += 1
        if s >= c:
            break

    return score


def part2():
    d = []
    result = 0
    with open('input.txt', 'r') as reader:
        for li in reader:
            cd = []
            line = li.replace('\n', '')
            for c in line:
                cd.append(int(c))
            d.append(cd)

    for y in range(1, len(d)-1):
        for x in range(1, len(d)-1):
            c = d[y][x]
            local_res = 1

            l = d[y][0:x][::-1]
            r = d[y][x+1:len(d)]
            t = []
            b = []
            for i in range(0, y):
                t.insert(0, d[i][x])
            for i in range(y+1, len(d)):
                b.append(d[i][x])

            local_res *= sibling_score(l, c)
            local_res *= sibling_score(t, c)
            local_res *= sibling_score(r, c)
            local_res *= sibling_score(b, c)
            if local_res > result:
                result = local_res

    print(result)


part1()
part2()
