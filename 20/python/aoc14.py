import re

def apply_mask(val, mask):
    #expand val into 36-bit binary value
    bin_val = bin(int(val))[2:].zfill(36)
    masked_val = ["0"] * 36
    #iterate from the b
    for i in range(len(mask)):
        mask_val = mask[i]
        if mask_val == "X":
            masked_val[i] = bin_val[i]
        elif mask_val in ["1","0"]:
            masked_val[i] = mask_val
        
    return int("".join(masked_val), 2)


def bitmask1(data):
    memory = {}
    current_mask = ""
    # loop through instructions and fill up 
    for line in data:
        if "mask" in line:
            current_mask = line.split(" ")[2].strip()
            # print(f"new_mask: {mask}")
        else:
            instructions = line.split("=")
            target_addr = re.findall(r"\[.*?\]", instructions[0])[0]
            target_addr.replace("[","")
            target_addr.replace("]","")
            value_to_save = instructions[1].strip()
            # print(f"target: {target_addr}, val: {value_to_save}")
            # apply bitmask to value
            masked_val = apply_mask(value_to_save, current_mask)
            memory[target_addr] = masked_val
    
    # iterate through memory and return sum
    mem_sum = 0
    for _,val in memory.items():
        mem_sum += val
    
    return mem_sum

def apply_mask2(val, mask):
    # print(f"apply_mask2: val -> {val}, mask -> {mask}")
    #expand val into 36-bit binary value
    bin_val = bin(int(val))[2:].zfill(36)
    masked_val = ["0"] * 36
    addresses = [masked_val]
    #iterate from the b
    for i in range(len(mask)):
        mask_val = mask[i]
        if mask_val == "X":
            # a floating bit
            # store the copies that are made and add them to the addresses after processing
            copies = []
            for address in addresses:
                copy = address.copy()
                address[i] = "0"
                copy[i] = "1"
                copies.append(copy)
            addresses.extend(copies)
        elif mask_val == "0":
            for address in addresses:
                address[i] = bin_val[i]
        elif mask_val == "1":
            for address in addresses:
                address[i] = "1"
    
    for i in range(len(addresses)):
        addresses[i] = int("".join(addresses[i]), 2)
    
    return addresses

def bitmask2(data):
    memory = {}
    current_mask = ""
    # loop through instructions and fill up 
    for line in data:
        if "mask" in line:
            current_mask = line.split(" ")[2].strip()
            # print(f"new_mask: {mask}")
        else:
            instructions = line.split("=")
            target_addr = re.findall(r"\[.*?\]", instructions[0])[0]
            target_addr = target_addr.replace("[","").replace("]","")
            value_to_save = int(instructions[1].strip())
            # print(f"target: {target_addr}, val: {value_to_save}")
            # apply bitmask to value
            masked_addrs = apply_mask2(target_addr, current_mask)
            for addr in masked_addrs:
                memory[addr] = value_to_save
    
    # iterate through memory and return sum
    mem_sum = 0
    for _,val in memory.items():
        mem_sum += val
    
    return mem_sum
    

if __name__ == "__main__":
    text_file = open("../input/aoc14.txt", "r")
    string_data = text_file.readlines()
    ans = bitmask2(string_data)
    print(f"the sum of memory addresses is {ans}")
