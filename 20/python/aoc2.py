
def is_valid(char_range, char, passwd):
    range_split = char_range.split("-")
    char_count = 0
    for c in passwd:
        if c == char:
            char_count += 1
    return char_count >= int(range_split[0]) and char_count <= int(range_split[1])

def xor(x, y):
    return bool(x)+bool(y) == 1

def is_valid2(char_range, char, passwd):
    range_split = char_range.split("-")
    return xor(passwd[int(range_split[0]) - 1] == char,  passwd[int(range_split[1]) - 1] == char)

def valid_passwords(arr):
    valid = 0
    for line in arr:
        split = line.split(" ")
        char_range = split[0]
        char = split[1][0]
        passwd = split[2]
        if is_valid2(char_range, char, passwd):
            valid += 1
    return valid


if __name__ == "__main__":
    text_file = open("../input/aoc2.txt", "r")
    string_data = text_file.readlines()
    ans = valid_passwords(string_data)
    print("the number of valid passwords is {}".format(ans))
   