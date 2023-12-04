def part1():
    ans = 0
    with open("input.txt", "r") as reader:
        for l in reader:
            winning = {}
            score = 0
            line = l.replace("\n", "")
            segments = line.split(" | ")
            for wn in segments[0].split(": ")[1].split(" "):
                if wn == "":
                    continue
                winning[wn] = 1

            for n in segments[1].split(" "):
                if n == "":
                    continue

                if n in winning:
                    if score == 0:
                        score = 1
                    else:
                        score *= 2

            ans += score

        print(ans)


def part2():
    with open("input.txt", "r") as reader:
        card_count = {}
        for l in reader:
            winning = {}
            score = 0
            line = l.replace("\n", "")
            segments = line.split(" | ")
            card_id = int(segments[0].split(": ")[0].split(" ")[-1])
            if card_id in card_count:
                card_count[card_id] += 1
            else:
                card_count[card_id] = 1
            for wn in segments[0].split(": ")[1].split(" "):
                if wn == "":
                    continue
                winning[wn] = 1

            for n in segments[1].split(" "):
                if n == "":
                    continue

                if n in winning:
                    score += 1

            for c in range(score):
                card_idx = card_id+c+1
                if card_idx in card_count:
                    card_count[card_idx] += card_count[card_id]
                else:
                    card_count[card_idx] = card_count[card_id]

        sum = 0
        for v in card_count.values():
            sum += v
        print(sum)


part1()
part2()
