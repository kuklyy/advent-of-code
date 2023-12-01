def part1():
    with open('input.txt', 'r') as reader:
        sum = 0
        for l in reader:
            line = l.replace('\n', '')
            first, second = None, None
            for c in line:
                if c.isdigit():
                    if first == None:
                        first = c
                    else:
                        second = c

            if second == None:
                second = first

            d = int(str(first)+str(second))
            sum += d
        print(sum)


def part2():
    dic = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    with open('input.txt', 'r') as reader:
        sum = 0
        for l in reader:
            line = l.replace('\n', '')
            first, second = None, None
            for i, c in enumerate(line):
                if c.isdigit():
                    if first == None:
                        first = c
                    else:
                        second = c
                    continue

                for k, v in dic.items():
                    if line[i:].startswith(k):
                        if first == None:
                            first = v
                        else:
                            second = v

            if second == None:
                second = first

            d = int(str(first)+str(second))
            sum += d
        print(sum)


part1()
part2()
