from pathlib import Path

file_dir = Path(__file__).resolve().parent

def get_lines(filename):
    with open(f"{file_dir}/../input/{filename}") as f:
        return list(f)
    

def process_input(str):
    return int(str.strip())

# take list of ints and return how many times the subsequent value is larger than the previous
def num_increases(input):
    count = 0
    prev_depth = input[0]
    for depth in input[1:]:
        curr_depth = depth
        if curr_depth > prev_depth:
            count += 1
        prev_depth = curr_depth
    return count

def get_sum_in_range(input, x, y):
    sum = 0
    # range is not inclusive
    for idx in range(x, y+1):
        sum += input[idx]
    return sum

def num_increases_window(input, w_size):
    count = 0
    w_start = 1
    w_end = w_size
    prev_sum = get_sum_in_range(input, 0, w_size-1)
    while w_end < len(input):
        curr_sum = get_sum_in_range(input, w_start, w_end)
        if curr_sum > prev_sum:
            count += 1
        prev_sum = curr_sum
        w_start += 1
        w_end += 1
    return count

if __name__ == "__main__":
    lines = get_lines("day1.txt")
    input = [process_input(line) for line in lines]
    print(num_increases_window(input, 3))