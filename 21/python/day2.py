from pathlib import Path

file_dir = Path(__file__).resolve().parent

def get_lines(filename):
    with open(f"{file_dir}/../input/{filename}") as f:
        return list(f)

# input: forward 2
# output: ['forward', 2]
def process_input(str):
    stripped = str.strip()
    split = str.split(" ")
    split[1] = int(split[1])
    return split

# takes in array of position movements and calculates product of final horizontal and vertical position
def product_position(input):
    h_pos = 0
    v_pos = 0
    for line in input:
        dir = line[0]
        mag = line[1]
        if dir == "forward":
            h_pos += mag
        elif dir == "up":
            v_pos -= mag
        elif dir == "down":
            v_pos += mag
    
    return h_pos * v_pos


if __name__ == "__main__":
    lines = get_lines("day2.txt")
    input = [process_input(line) for line in lines]
    print(product_position(input))