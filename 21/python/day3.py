from pathlib import Path

file_dir = Path(__file__).resolve().parent

def get_lines(filename):
    with open(f"{file_dir}/../input/{filename}") as f:
        return list(f)

def power_consumption(input):
    # nested dict pos -> int -> count 
    pos_dict = {}
    for line in input:
        line = line.strip()
        for idx, digit in enumerate(list(line)):
            dig_i = int(digit)
            if idx in pos_dict:
                if dig_i in pos_dict.get(idx):
                    pos_dict.get(idx)[dig_i] = pos_dict.get(idx)[dig_i] + 1
                else:
                    pos_dict.get(idx)[dig_i] = 1
            else:
                pos_dict[idx] = {0: 0, 1: 0}
                pos_dict.get(idx)[dig_i] = 1
    
    # (pos_dict)
    # gamma - most common bits
    # epsilon - least common bits
    gamma = ""
    epsilon = ""

    for idx in pos_dict:
        if pos_dict.get(idx).get(1) > pos_dict.get(idx).get(0):
            gamma += "1"
            epsilon += "0"
        else:
            gamma += "0"
            epsilon += "1"

    print(f"gamma: {gamma}, epsilon: {epsilon}")
    return int(gamma, 2) * int(epsilon, 2)

    

if __name__ == "__main__":
    lines = get_lines("day3.txt")
    print(power_consumption(lines))