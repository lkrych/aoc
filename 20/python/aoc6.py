def customs(data):
    current_customs = {}
    q_count = 0
    for line in data:
        line = line.strip()
        # time to check if the passport is valid
        if len(line) < 1:
            q_count += len(current_customs.items())
            current_customs = {}
        else:
            for c in line:
                current_customs[c] = True
        
    return q_count


if __name__ == "__main__":
    text_file = open("../input/aoc6.txt", "r")
    string_data = text_file.readlines()
    ans = customs(string_data)
    print("the number of customs questions is {}".format(ans))
