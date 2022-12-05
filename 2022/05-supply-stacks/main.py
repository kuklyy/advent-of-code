def add_to_stacks(stacks, line):
    space_c = 0
    stack_index = 0
    for c in line:
        if space_c == 3:
            stack_index += 1
            space_c = 0
            continue
        if c == ' ':
            space_c += 1
            continue

        stacks[stack_index].append(c)
        stack_index += 1
        space_c = 0


def process_instructions_p1(stacks, line):
    sl = line.split(' ')
    intructions = []
    i = 0
    for s in sl:
        if s == '':
            continue
        if i > 0:
            intructions.append(int(s)-1)
            continue
        intructions.append(int(s))
        i += 1

    for i in range(intructions[0]):
        m = stacks[intructions[1]].pop(0)
        stacks[intructions[2]].insert(0, m)

    return


def part1():
    processing_intructions = False
    stacks_n = 9
    stacks = []
    for _ in range(stacks_n):
        stacks.append([])
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '').replace('[', '').replace(']', '')
            if line == '':
                processing_intructions = True
                continue
            if line.replace(' ', '').isdigit():
                continue

            if processing_intructions:
                line = line.replace('move', '').replace(
                    'from', '').replace('to', '')
                process_instructions_p1(stacks, line)
            else:
                add_to_stacks(stacks, line)

    message = ''
    for s in stacks:
        message += s[0]

    print(message)


def process_instructions_p2(stacks, line):
    sl = line.split(' ')
    intructions = []
    i = 0
    for s in sl:
        if s == '':
            continue
        if i > 0:
            intructions.append(int(s)-1)
            continue
        intructions.append(int(s))
        i += 1

    m = stacks[intructions[1]][0:intructions[0]]
    stacks[intructions[2]] = m+stacks[intructions[2]]
    stacks[intructions[1]] = stacks[intructions[1]][intructions[0]:]

    return


def part2():
    processing_intructions = False
    stacks_n = 9
    stacks = []
    for _ in range(stacks_n):
        stacks.append([])
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '').replace('[', '').replace(']', '')
            if line == '':
                processing_intructions = True
                continue
            if line.replace(' ', '').isdigit():
                continue

            if processing_intructions:
                line = line.replace('move', '').replace(
                    'from', '').replace('to', '')
                process_instructions_p2(stacks, line)
            else:
                add_to_stacks(stacks, line)

    message = ''
    for s in stacks:
        message += s[0]

    print(message)


part1()
part2()
