
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


def is_ticket_valid2(val, rules):
    valid_rules = []
    for rule in rules:
        split_by_or = rule.split("or")
        for i in range(len(split_by_or)):
            split_by_or[i] = split_by_or[i].strip()
        
        for s in split_by_or:
            split_by_dash = s.split("-")
            if val >= int(split_by_dash[0]) and val <= int(split_by_dash[1]):
                valid_rules.append(rule)

    return valid_rules


def parse_tickets2(data):
    idx = 0
    rules = {}
    my_ticket = []
    valid_tickets = []

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
        rules[rule] = split[0].strip()
        idx += 1
    print(f"rules: {rules}")
    # save my ticket
    while True:
        if len(data[idx].strip()) <= 1:
            idx += 1
            continue
        # print(f"saving my ticket: {data[idx].strip()}")
        if "nearby tickets:" in data[idx]:
            idx += 1
            break
        my_ticket = data[idx].strip().split(",")
        idx += 1
    

    # check the validity of the rest of the tickets
    while idx < len(data):
        if len(data[idx].strip()) <= 1:
            idx += 1
            continue
        current_ticket = data[idx].strip().split(",")
        valid_ticket = True
        for j in range(len(current_ticket)):
            to_check = int(current_ticket[j])
            valid_ticket = is_ticket_valid(to_check, rules.keys())
            if not valid_ticket:
                # we don't need to check the rest of the ticket
                break
        if valid_ticket:
            valid_tickets.append(current_ticket)
            
        idx += 1
    
    # now we have an array of valid tickets, we need to determine which fields belong to which rules
    fields = {}
    for rule, field in rules.items():
        fields[field] = {}
    
    #create a distribution of valid fields
    print(f"fields: {fields}")
    for ticket in valid_tickets:
        print(f"checking ticket {ticket}")
        for j in range(len(ticket)):
            to_check = int(ticket[j])
            valid_rules = is_ticket_valid2(to_check, rules.keys())
            for r in valid_rules:
                rule = rules[r]
                if j in fields[rule]:
                    fields[rule][j] += 1
                else:
                    fields[rule][j] = 1
    
    print(f"distribution: {fields}")

    # iterate through the distribution until you've determined unique designation
    designated_fields = {}
    num_items = len(fields.keys())
    num_tickets = len(valid_tickets)
    while len(designated_fields.keys()) < num_items:
        for rule, distribution in fields.items():
            # print(f"checking rule: {rule}")
            potential_fields = []
            for idx, count in distribution.items():
                #only iterate through non-designated fields
                if idx in designated_fields:
                    continue
                # print(f"checking idx: {idx}:{count}")
                # print(f"num_items: {num_items}")
                if count == num_tickets:
                    potential_fields.append(idx)
            # only designate a field if it is unique

            # print(f"potential_fields after iterating for {rule}: {potential_fields}")
            if len(potential_fields) == 1:
                # add it to the answer
                print(f"adding {potential_fields[0]} to designated fields for {rule}")
                designated_fields[potential_fields[0]] = rule 
    
    # calculate product 
    print(designated_fields)
    product = 1
    for idx, val in designated_fields.items():
        if "departure" in val:
            product *= int(my_ticket[idx])
    return product


        


if __name__ == "__main__":
    text_file = open("../input/aoc16.txt", "r")
    string_data = text_file.readlines()
    ans = parse_tickets2(string_data)
    print(f"the error rate is {ans}")