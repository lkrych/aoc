def is_passport_valid(passport):
    valid_fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
    for field in valid_fields:
        if field not in passport:
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
