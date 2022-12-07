def process_line_p1(l):
    for k in range(0, len(l) - 3):
        b = set(l[k:k+4])
        if len(b) == 4:
            return k + 4


def process_line_p2(l):
    for k in range(0, len(l) - 13):
        b = set(l[k:k+14])
        if len(b) == 14:
            return k + 14


def part1():
    with open('input.txt', 'r') as reader:
        for l in reader:
            r = process_line_p1(l.replace('\n', ''))
            print(r)


def part2():
    with open('input.txt', 'r') as reader:
        for l in reader:
            r = process_line_p2(l.replace('\n', ''))
            print(r)


part1()
part2()
