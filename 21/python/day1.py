from pathlib import Path

file_dir = Path(__file__).resolve().parent

def process_input(str):
    return int(str.strip())

# take list of ints and return how many times the subsequent value is larger than the previous
def num_increases(input):
    count = 0
    prev_depth = process_input(input[0])
    for depth in input[1:]:
        if depth: # check for empty string
            curr_depth = process_input(depth)
            if curr_depth > prev_depth:
                count += 1
            prev_depth = curr_depth
    return count

if __name__ == "__main__":
    with open(f"{file_dir}/../input/day1.txt") as f:
        input = list(f)
        print(num_increases(input))