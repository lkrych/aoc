# part 1
# def is_passport_valid(passport):
#     valid_fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
#     for field in valid_fields:
#         if field not in passport:
#             return False
#     return True


# def passport(data):
#     current_passport = {}
#     valid = 0
#     for line in data:
#         line = line.strip()
#         # time to check if the passport is valid
#         if len(line) < 1:
#             if is_passport_valid(current_passport):
#                 valid += 1
#             current_passport = {}
#         else:
#             split = line.split(" ")
#             for fields in split:
#                 field = fields.split(":")
#                 current_passport[field[0]] = field[1]
#     if is_passport_valid(current_passport):
#         valid += 1

#     return valid



# if __name__ == "__main__":
#     text_file = open("../input/aoc4.txt", "r")
#     string_data = text_file.readlines()
#     ans = passport(string_data)
#     print("the number of valid passports is {}".format(ans))

import re

def is_passport_valid(passport):
    items = passport.items()
    valid_fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
    for field in valid_fields:
        if field not in passport:
            return False
    for item in items:
        field = item[0]
        val = item[1]
        if field == "byr":
            if int(val) > 2002 or int(val) < 1920:
                return False
        elif field == "iyr":
            if int(val) > 2020 or int(val) < 2010:
                return False
        elif field == "eyr":
            if int(val) > 2030 or int(val) < 2020:
                return False
        elif field == "hgt":
            if val[-2:] == "cm":
                if int(val[:-2]) > 193 or int(val[:-2]) < 150:
                    return False
            elif val[-2:] == "in":
                if int(val[:-2]) > 76 or int(val[:-2]) < 59:
                        return False
        elif field == "hcl":
            if len(val) == 7:
                if not re.search("#([a-f0-9])", val):
                    return False
            else:
                return False
        elif field == "ecl":
            valid_colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
            if val in valid_colors:
                continue
            else:
                return False
        elif field == "pid":
            if len(val) != 9:
                return False
    return True


def passport(data):
    current_passport = {}
    valid = 0
    for line in data:
        line = line.strip()
        # time to check if the passport is valid
        if len(line) < 1:
            if is_passport_valid(current_passport):
                valid += 1
            current_passport = {}
        else:
            split = line.split(" ")
            for fields in split:
                field = fields.split(":")
                current_passport[field[0]] = field[1]
    if is_passport_valid(current_passport):
        valid += 1
        
    return valid



if __name__ == "__main__":
    text_file = open("../input/aoc4.txt", "r")
    string_data = text_file.readlines()
    ans = passport(string_data)
    print("the number of valid passports is {}".format(ans))
