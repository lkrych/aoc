
def process(line, acc, idx):
    split = line.split(" ")
    action = split[0]
    sign = split[1][0]
    amount = int(split[1][1:])
    print("action: {}, sign: {}, amt: {}".format(action, sign, amount))
    if action == "acc":
        if sign == "+":
            acc += amount
        elif sign == "-":
            acc -= amount
        idx += 1
    elif action == "jmp":
        if sign == "+":
            idx += amount
        elif sign == "-":
            idx -= amount
    elif action == "nop":
        idx += 1
    return (acc, idx)

def halting(data):
    acc = 0
    idx_map = {}
    idx = 0
    while True:
        line = data[idx]
        line = line.strip()
        (new_acc, new_idx) = process(line, acc, idx)
        if new_idx in idx_map:
            return acc
        else:
            idx_map[new_idx] = True
            acc = new_acc
            idx = new_idx

def process2(line, current_state, already_tested):
    #process the input
    acc, idx = current_state.acc, current_state.idx
    started_test = current_state.testing
    split = line.split(" ")
    action = split[0]
    sign = split[1][0]
    amount = int(split[1][1:])
    #debug
    print("action: {}, sign: {}, amt: {}, idx: {}".format(action, sign, amount, idx))
    #actions
    if action == "acc":
        if sign == "+":
            acc += amount
        elif sign == "-":
            acc -= amount
        idx += 1
    elif action == "jmp":
        if not started_test and current_state.idx not in already_tested:
            #nop behavior
            idx += 1
            started_test = True
        else:
            #normal behavior
            if sign == "+":
                idx += amount
            elif sign == "-":
                idx -= amount
    elif action == "nop":
        if not started_test and current_state.idx not in already_tested:
            #jmp behavior
            if sign == "+":
                idx += amount
            elif sign == "-":
                idx -= amount
            started_test = True
        else:
            #normal behavior
            idx += 1
    return (acc, idx, started_test)

class HaltingState:
    def __init__(self, acc, idx, idx_map):
        self.acc = acc
        self.idx = idx
        self.idx_map = idx_map
        self.testing = False

def halting2(data):
    # to solve this problem we need to go through possible paths of execution
    # until we reach the end of the data array (idx = len(data))
    # this means that we need to keep track of 
    # 1. the paths that we have tested
    # 2. the state of the path before we begin testing
    # if we detect an infinite loop in the current path we reset the state and try an alternative path
    current_state = HaltingState(0,0,{})
    saved_state = HaltingState(0,0,{})
    tested_idxs = {}
    while current_state.idx < len(data):
        line = data[current_state.idx]
        line = line.strip()
        (new_acc, new_idx, testing) = process2(line, current_state, tested_idxs)

        if testing != current_state.testing:
            # we've started testing a path
            # save the state
            print("we've started testing")
            saved_state.idx_map = current_state.idx_map
            saved_state.acc = current_state.acc
            saved_state.idx = current_state.idx
            # keep track of the testing
            current_state.testing = True
            tested_idxs[current_state.idx] = True
        
        if new_idx in current_state.idx_map:
            # we've encountered an infinite loop
            # reset the state
            print("we've encountered an infinite loop")
            current_state.idx_map = saved_state.idx_map
            current_state.acc = saved_state.acc
            current_state.idx = saved_state.idx
            current_state.testing = False
        else:
            # everything is good, we are moving through the current path
            current_state.idx_map[new_idx] = True
            current_state.acc = new_acc
            current_state.idx = new_idx



if __name__ == "__main__":
    text_file = open("../input/aoc8.txt", "r")
    string_data = text_file.readlines()
    ans = halting2(string_data)
    print("the acc once the program terminates is {}".format(ans))
