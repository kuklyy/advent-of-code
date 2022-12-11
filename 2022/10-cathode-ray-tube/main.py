def check_c(c, x):
    if c % 40 == 20:
        return c * x
    return 0


def draw(c, x):
    m = c % 40
    if m == 0:
        print()
    if m in [x-1, x, x+1]:
        print("#", end="")
    else:
        print(".", end="")


def part1():
    c = 0
    s = 0
    x = 1
    with open('input.txt', 'r') as reader:
        for l in reader:
            cmd = l.replace('\n', '').split(' ')
            if cmd[0] == 'noop':
                s += check_c(c, x)
                c += 1
            else:
                s += check_c(c, x)
                c += 1
                s += check_c(c, x)
                x += int(cmd[1])
                c += 1

    print(s)


def part2():
    c = 0
    x = 1
    with open('input.txt', 'r') as reader:
        for l in reader:
            cmd = l.replace('\n', '').split(' ')
            if cmd[0] == 'noop':
                draw(c, x)
                c += 1
            else:
                draw(c, x)
                c += 1
                draw(c, x)
                c += 1
                x += int(cmd[1])


part1()
part2()
