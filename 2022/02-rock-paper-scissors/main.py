score_dict_p1 = {
    'A': {
        'X': 4,
        'Y': 8,
        'Z': 3,
    },
    'B': {
        'X': 1,
        'Y': 5,
        'Z': 9,
    },
    'C': {
        'X': 7,
        'Y': 2,
        'Z': 6,
    },
}

score_dict_p2 = {
    'A': {
        'X': 3,
        'Y': 4,
        'Z': 8,
    },
    'B': {
        'X': 1,
        'Y': 5,
        'Z': 9,
    },
    'C': {
        'X': 2,
        'Y': 6,
        'Z': 7,
    },
}


def part1():
    score = 0
    with open('input.txt', 'r') as reader:
        for line in reader:
            score += score_dict_p1[line[0]][line[2]]

    print(score)


def part2():
    score = 0
    with open('input.txt', 'r') as reader:
        for line in reader:
            score += score_dict_p2[line[0]][line[2]]

    print(score)


part1()
part2()
