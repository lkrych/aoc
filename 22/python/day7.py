from collections import deque
import pprint
# read from file input
file_name = input()
f = open(f"/Users/lkrych/aoc/22/input/{file_name}.txt", "r")

#https://stackoverflow.com/a/23049823/4458404
def get_item(db, keys):
    #iterate through each key and find subitem
    for key in keys:
        if key in db:
            db = db[key]
        else:
            db[key] = {}
            db = db[key]
    return db

# every time we need to set an item, we need to use the elements in the stack to drill down and fetch the item
def set_item(db, keys, file_name, file_size):
    # iterate through each key and make sure 
    db = get_item(db, list(keys))
    if file_size != "dir":
        db[file_name] = file_size

# recursively collect dir sizes
def get_dict_size(source_dict, sum_dict):
    for key in source_dict.keys():
        if isinstance(key, dict):
            sum_dict[key] = get_dict_size(key, sum_dict)
        else:
            sum_dict[key] += source_dict[key]
    return sum_dict


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
            set_item(files, directories, file_name, file_size)


dir_size = {}
dir_size = get_dict_size(files, dir_size)
pprint.pprint(dir_size)
