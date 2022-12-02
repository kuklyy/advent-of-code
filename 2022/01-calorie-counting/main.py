def part1():
    max_sum = 0
    curr_sum = 0

    with open('input.txt', 'r') as reader:
        for line in reader:
            if line == '\n':
                if curr_sum > max_sum:
                    max_sum = curr_sum
                curr_sum = 0
                continue
            curr_sum += int(line)

    if curr_sum > max_sum:
        max_sum = curr_sum

    print(max_sum)


def part2():
    top = []
    curr_sum = 0

    with open('input.txt', 'r') as reader:
        for line in reader:
            if line == '\n':
                top.append(curr_sum)
                top.sort(reverse=True)
                top = top[0:3]
                curr_sum = 0
                continue
            curr_sum += int(line)

    top.append(curr_sum)
    top.sort(reverse=True)
    top = top[0:3]
    print(sum(top))


part1()
part2()
