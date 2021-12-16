from os import getgrouplist
from pathlib import Path

file_dir = Path(__file__).resolve().parent

def get_lines(filename):
    with open(f"{file_dir}/../input/{filename}") as f:
        return list(f)


def get_pos_dict(input):
    pos_dict = {}
    for line in input:
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
    return pos_dict

def power_consumption(input):
    # nested dict pos -> int -> count 
    pos_dict = get_pos_dict(input)
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

def get_input_by_val_idx(input, val, idx):
    new_input = []
    for line in input:
        if int(line[idx]) == val:
            new_input.append(line)
    return new_input

def life_support(input):
    oxy_input = input.copy()
    co2_input = input.copy()
    
    oxy_idx = 0
    while len(oxy_input) > 1:
        pos_dict = get_pos_dict(oxy_input)
        if pos_dict.get(oxy_idx).get(1) >= pos_dict.get(oxy_idx).get(0):
            oxy_input = get_input_by_val_idx(oxy_input, 1, oxy_idx)
        else:
            oxy_input = get_input_by_val_idx(oxy_input, 0, oxy_idx)
        oxy_idx += 1

    co2_idx = 0
    while len(co2_input) > 1:
        pos_dict = get_pos_dict(co2_input)
        if pos_dict.get(co2_idx).get(1) >= pos_dict.get(co2_idx).get(0):
            co2_input = get_input_by_val_idx(co2_input, 0, co2_idx)
        else:
            co2_input = get_input_by_val_idx(co2_input, 1, co2_idx)
        co2_idx += 1
    
    print(f"oxy: {oxy_input}, co2:{co2_input}")
    return int(oxy_input[0], 2) * int(co2_input[0], 2)

if __name__ == "__main__":
    lines = get_lines("day3.txt")
    lines = [line.strip() for line in lines]
    print(life_support(lines))