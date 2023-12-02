def part1():
    with open('input.txt', 'r') as reader:
        res = 0
        for l in reader:
            line = l.replace('\n', '')
            segments = line.split(": ")
            game_id = segments[0].split(" ")[1]
            is_possible = True
            for draw in segments[1].split("; "):
                if not is_possible:
                    break
                cubes = draw.split(", ")
                for cube in cubes:
                    q, c = int(cube.split(" ")[0]), cube.split(" ")[1]
                    if c == "red" and q > 12:
                        is_possible = False
                        break
                    if c == "green" and q > 13:
                        is_possible = False
                        break
                    if c == "blue" and q > 14:
                        is_possible = False
                        break

            if is_possible:
                res += int(game_id)
        print(res)


def part2():
    with open('input.txt', 'r') as reader:
        res = 0
        for l in reader:
            line = l.replace('\n', '')
            segments = line.split(": ")
            dic = {
                "green": 0,
                "blue": 0,
                "red": 0,
            }
            for draw in segments[1].split("; "):
                cubes = draw.split(", ")
                for cube in cubes:
                    q, c = int(cube.split(" ")[0]), cube.split(" ")[1]
                    if dic[c] < q:
                        dic[c] = q

            res += dic["green"] * dic["blue"] * dic["red"]
        print(res)


part1()
part2()
