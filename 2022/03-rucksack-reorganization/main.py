def asci_lower_to_prior(asci):
    return asci - 96


def asci_upper_to_prior(asci):
    return asci - 64 + 26


def find_common(first, second):
    for i in first:
        for j in second:
            if i == j:
                return i


def part1():
    score = 0
    with open('input.txt', 'r') as reader:
        for line in reader:
            line = line.replace('\n', '')
            first = line[:int(len(line)/2)]
            second = line[int(len(line)/2):]
            common = find_common(first, second)
            if common.upper() == common:
                score += asci_upper_to_prior(ord(common))
            else:
                score += asci_lower_to_prior(ord(common))
    print(score)


def process_line(line, dict):
    local_dict = {}
    for c in line:
        if dict.get(c) == None:
            dict[c] = 1
            local_dict[c] = 1
            continue
        if local_dict.get(c) == None:
            local_dict[c] = 1
            dict[c] += 1


def part2():
    score = 0
    i = 0
    d = {}
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            process_line(line, d)
            i += 1
            if i == 3:
                common = ''
                for k, v in d.items():
                    if v == 3:
                        common = k
                        break
                if common.upper() == common:
                    score += asci_upper_to_prior(ord(common))
                else:
                    score += asci_lower_to_prior(ord(common))
                i = 0
                d = {}

    print(score)


part1()
part2()
