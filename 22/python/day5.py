from collections import deque
import re

# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

stacks = {}
for line in f:
    l = line.rstrip()
    # create stacks
    if "[" in l:
        # replace an empty space with a placeholder text, we can then skip over this placeholder and still enumerate based off index
        l = l.replace("    ", "placeholder ")
        split = l.split(" ")
        print(split)
        for idx, el in enumerate(split, start=1):
            #skip over the placeholders
            if el == "placeholder" or not el:
                continue
            e = el.replace("[", "")
            e = e.replace("]", "")
            e = e.replace("placeholder", "")
            if idx not in stacks:
                d = deque()
                # we are using appendleft because we are reading from the top of the stack
                d.appendleft(e)
                print(f"adding {e} to stack {idx}")
                stacks[idx] = d
            else:
                print(f"adding {e} to stack {idx}")
                stacks[idx].appendleft(e)
    # once stacks are created start processing the moves
    elif l.startswith("move"):
        # example input: move 1 from 2 to 1
        moves = re.findall("[0-9]+", l)
        moves = [int(el) for el in moves]
        num_to_move, from_col, to_col = moves
        for _ in range(num_to_move):
            item = stacks[from_col].pop()
            stacks[to_col].append(item)

for idx in stacks.keys():
    print(f"Idx: {idx}, item: {stacks[idx][-1]}") 
