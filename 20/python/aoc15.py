
def memory_game(data, target_idx):
    starting_vals = data[0].strip().split(",")
    current_idx = len(starting_vals) + 1
    current_val = 0
    memory = {}

    # populate the memory
    for i in range(len(starting_vals)):
        val = starting_vals[i]
        # turns start at 1
        memory[int(val)] = i + 1

    while current_idx < target_idx:
        if current_val not in memory:
            memory[current_val] = current_idx
            current_val = 0
        else:
            old_current_val = current_val
            current_val = current_idx - memory[old_current_val]
            memory[old_current_val] = current_idx
        current_idx += 1
    
    return current_val


if __name__ == "__main__":
    text_file = open("../input/aoc15.txt", "r")
    string_data = text_file.readlines()
    ans = memory_game(string_data, 2020)
    print(f"the 2020th value is {ans}")