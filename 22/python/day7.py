from functools import reduce
import operator
from collections import deque
import pprint
# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

# https://stackoverflow.com/a/61493389/4458404
def get_item(db, keys):
    try: return reduce(operator.getitem, keys, db)
    except Exception as e:
        return None

def set_item(db, keys, file_name, file_size):
    for index in range(1,len(keys)):
            subitem = get_item(db, keys[:index])
            if not isinstance(subitem, dict):
                get_item(db, keys[:index][:-1])[keys[:index][-1]] = {}
    get_item(db, keys[:-1])[file_name] = file_size
    return db

files = {}
current_cmd = None
directories = deque()
for line in f:
    l = line.strip()
    if l.startswith("$"):
        #evaluate which command has been issued
        split = l.split(" ")
        cmd = split[1]
        if cmd == "cd":
            current_cmd = cmd
            target = split[2]
 
            if target == "..":
                directories.pop()
            else:
                directories.append(target)
            
        elif cmd == "ls":
            current_cmd = cmd
    else:
        #continue parsing command
        if current_cmd == "ls":
            split = l.split(" ")
            file_size = split[0]
            file_name = split[1]
            directories = set_item(files, directories, file_name, file_size)

pprint.pprint(files)