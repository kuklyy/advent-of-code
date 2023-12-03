def evaluate(x, y, lenn, maxx, symbols):
    pos = []

    for xx in range(lenn):
        pos.append(f"{x-xx}:{max(y-1,0)}")
        pos.append(f"{x-xx}:{min(y+1,maxx-1)}")

    pos.append(f"{max(x-lenn,0)}:{max(y-1,0)}")
    pos.append(f"{max(x-lenn,0)}:{y}")
    pos.append(f"{max(x-lenn,0)}:{min(y+1,maxx-1)}")

    pos.append(f"{min(x+1,maxx)}:{max(y-1,0)}")
    pos.append(f"{min(x+1,maxx)}:{y}")
    pos.append(f"{min(x+1,maxx)}:{min(y+1,maxx-1)}")

    for p in pos:
        if symbols.get(p) == 1:
            return True

    return False


def evaluate_p2(x, y, n, maxx, symbols):
    pos = []
    lenn = len(n)

    for xx in range(lenn):
        pos.append(f"{x-xx}:{max(y-1,0)}")
        pos.append(f"{x-xx}:{min(y+1,maxx-1)}")

    pos.append(f"{max(x-lenn,0)}:{max(y-1,0)}")
    pos.append(f"{max(x-lenn,0)}:{y}")
    pos.append(f"{max(x-lenn,0)}:{min(y+1,maxx-1)}")

    pos.append(f"{min(x+1,maxx)}:{max(y-1,0)}")
    pos.append(f"{min(x+1,maxx)}:{y}")
    pos.append(f"{min(x+1,maxx)}:{min(y+1,maxx-1)}")

    for p in pos:
        if p in symbols:
            symbols.get(p).append(int(n))


def part1():
    symbols = {}
    sum = 0
    with open("input.txt", "r") as reader:
        y = 0
        for l in reader:
            x = 0
            line = l.replace("\n", "")
            for c in line:
                if c in "!@#$%^&*()-_=+[\{\}];:'\"|<,>?/":
                    symbols[f"{x}:{y}"] = 1
                x += 1
            y += 1

        y = 0
        reader.seek(0)

        for l in reader:
            line = l.replace("\n", "")
            x = 0
            n = ""
            for c in line:
                if c.isdigit():
                    n += c
                    if x == len(line)-1:
                        is_ok = evaluate(x, y, len(n), len(line), symbols)
                        if is_ok:
                            sum += int(n)
                        n = ""
                else:
                    if len(n) > 0:
                        is_ok = evaluate(x-1, y, len(n), len(line), symbols)
                        if is_ok:
                            sum += int(n)
                        n = ""

                x += 1
            y += 1
    print(sum)


def part2():
    symbols = {}
    sum = 0
    with open("input.txt", "r") as reader:
        y = 0
        for l in reader:
            x = 0
            line = l.replace("\n", "")
            for c in line:
                if c == "*":
                    symbols[f"{x}:{y}"] = []
                x += 1
            y += 1

        y = 0
        reader.seek(0)
        for l in reader:
            line = l.replace("\n", "")
            x = 0
            n = ""
            for c in line:
                if c.isdigit():
                    n += c
                    if x == len(line)-1:
                        evaluate_p2(x, y, n, len(line), symbols)
                        n = ""
                else:
                    if len(n) > 0:
                        evaluate_p2(x-1, y, n, len(line), symbols)
                        n = ""

                x += 1
            y += 1

        for gears in symbols.values():
            if len(gears) == 2:
                sum += gears[0]*gears[1]

        print(sum)


part1()
part2()
