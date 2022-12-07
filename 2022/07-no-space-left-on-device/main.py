def part1():
    path = []
    dir_dict = {}
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            if line.startswith('$'):
                cmd = line.split(' ')
                if cmd[1] == 'ls':
                    continue
                if cmd[2] == '/':
                    path = []
                elif cmd[2] == '..':
                    path.pop()
                else:
                    path.append(cmd[2])
            elif line.startswith('dir'):
                continue
            else:
                for i in range(0, len(path)+1):
                    p = '/'.join(path[0:len(path)-i])
                    s = int(line.split(' ')[0])
                    if dir_dict.get(p) == None:
                        dir_dict[p] = s
                    else:
                        dir_dict[p] += s

        s = 0
        for i in dir_dict.values():
            if i > 100000:
                continue
            s += i

        print(s)


def part2():
    path = []
    dir_dict = {}
    with open('input.txt', 'r') as reader:
        for l in reader:
            line = l.replace('\n', '')
            if line.startswith('$'):
                cmd = line.split(' ')
                if cmd[1] == 'ls':
                    continue
                if cmd[2] == '/':
                    path = []
                elif cmd[2] == '..':
                    path.pop()
                else:
                    path.append(cmd[2])
            elif line.startswith('dir'):
                continue
            else:
                for i in range(0, len(path)+1):
                    p = '/'.join(path[0:len(path)-i])
                    s = int(line.split(' ')[0])
                    if dir_dict.get(p) == None:
                        dir_dict[p] = s
                    else:
                        dir_dict[p] += s

        m = 70000000
        r = m - 30000000 - dir_dict['']

        for i in dir_dict.values():
            if r + i > 0 and i < m:
                m = i

        print(m)


part1()
part2()
