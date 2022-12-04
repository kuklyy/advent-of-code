def part1():
    score = 0
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            first, second = line.split(',')
            first_s, first_e = first.split('-')
            second_s, second_e = second.split('-')
            first_s, first_e, second_s, second_e = int(
                first_s), int(first_e), int(second_s), int(second_e)

            if ((second_s >= first_s and second_e <= first_e) or
                    (second_s <= first_s and second_e >= first_e)):
                score += 1

    print(score)


def part2():
    score = 0
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            first, second = line.split(',')
            first_s, first_e = first.split('-')
            second_s, second_e = second.split('-')
            first_s, first_e, second_s, second_e = int(
                first_s), int(first_e), int(second_s), int(second_e)

            if ((first_s >= second_s and first_s <= second_e) or
                (first_e >= second_s and first_e <= second_e) or
                (second_s >= first_s and second_e <= first_e) or
                    (second_e >= first_s and second_s <= first_e)):
                score += 1

    print(score)


part1()
part2()
