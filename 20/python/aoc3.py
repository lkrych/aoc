def toboggan(data):
    trees = 0
    curr_pos = 0
    for line in data:
        line = line.strip()
        if line[curr_pos] == '#':
            trees += 1
        curr_pos = (curr_pos + 3) % len(line)
    return trees


if __name__ == "__main__":
    text_file = open("../input/aoc3.txt", "r")
    string_data = text_file.readlines()
    ans = toboggan(string_data)
    print("the number of trees encountered is {}".format(ans))
   