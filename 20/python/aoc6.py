
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

def customs2(data):
    current_customs = {}
    group_count = 0
    q_count = 0
    for line in data:
        line = line.strip()
        # time to check if the customs is valid
        if len(line) < 1:
            for c in current_customs:
                print("checking if {} {} is {}".format(c, current_customs[c], group_count))
                if current_customs[c] == group_count:
                    q_count += 1
            group_count = 0
            current_customs = {}
        else:
            group_count += 1
            for c in line:
                if c in current_customs:
                    current_customs[c] += 1
                else:
                    current_customs[c] = 1
        
    return q_count

if __name__ == "__main__":
    text_file = open("../input/aoc6.txt", "r")
    string_data = text_file.readlines()
    ans = customs2(string_data)
    print("the number of customs questions is {}".format(ans))
