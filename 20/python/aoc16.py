
def is_ticket_valid(val, rules):
    # print(f"checking the validity of {val}")
    found_between_range = False
    for rule in rules:
        split_by_or = rule.split("or")
        for i in range(len(split_by_or)):
            split_by_or[i] = split_by_or[i].strip()
        
        for s in split_by_or:
            split_by_dash = s.split("-")
            if val >= int(split_by_dash[0]) and val <= int(split_by_dash[1]):
                found_between_range = True

    # if found_between_range == False:
    #     print(f"{val} doesn't validate with {rule}")
    return found_between_range

def parse_tickets(data):
    idx = 0
    rules = []
    my_ticket = ""
    invalid_fields_count = 0

    # parse rules
    while True:
        # print(f"parsing rules: {data[idx].strip()}")
        if len(data[idx].strip()) <= 1:
            idx += 1
            continue
        if "your ticket:" in data[idx]:
            idx += 1
            break
        line = data[idx].strip()
        split = line.split(":")
        rule = split[1].strip()
        rules.append(rule)
        idx += 1

    # save my ticket
    while True:
        if len(data[idx].strip()) <= 1:
            idx += 1
            continue
        # print(f"saving my ticket: {data[idx].strip()}")
        if "nearby tickets:" in data[idx]:
            idx += 1
            break
        my_ticket = data[idx].strip()
        idx += 1
    

    # check the validity of the rest of the tickets
    while idx < len(data):
        if len(data[idx].strip()) <= 1:
            idx += 1
            continue
        current_ticket = data[idx].strip().split(",")
        for j in range(len(current_ticket)):
            to_check = int(current_ticket[j])
            if not is_ticket_valid(to_check, rules):
                invalid_fields_count += to_check
                # we don't need to check the rest of the ticket
                break
        idx += 1
    
    return invalid_fields_count

    
        

if __name__ == "__main__":
    text_file = open("../input/aoc16.txt", "r")
    string_data = text_file.readlines()
    ans = parse_tickets(string_data)
    print(f"the error rate is {ans}")