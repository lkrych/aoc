#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <regex>
#include <string>
#include <vector>
#include "aoc_helper.h"

// part 1
// bool validPassport(std::map<std::string, std::string> passport, std::vector<std::string> required) {
//     // check if all the required fields are in the map
//     passport["checked"] = true;
//     std::cout << "checking if passport with " << passport.size() <<" keys is valid" << std::endl;
//     for (int i = 0; i < required.size(); i++) {
//         std::string req = required[i];
//         if ( passport.find(req) == passport.end()) {
//             return false;
//         } 
//     }
//     std::cout << "passport is valid!" << std::endl;
//     return true;
// }

// part 2
bool validPassport(std::map<std::string, std::string> passport, std::vector<std::string> required) {
    // check if all the required fields are in the map
    if (passport.size() != 7) {
        return false;
    }
    passport["checked"] = true;
    std::vector<std::string> valid_ecl = {"amb","blu","brn","gry","grn","hzl","oth"};
    for (int i = 0; i < required.size(); i++) {
        std::string req = required[i];
        if ( passport.find(req) == passport.end()) {
            return false;
        } else {
            std::string item = passport[req];
            if ("byr" == req) {
                int byr = std::stoi(item);
                if (byr < 1920 || byr > 2002) {
                    return false;
                }
            } else if("iyr" == req) {
                int iyr = std::stoi(item);
                if (iyr < 2010 || iyr > 2030) {
                    return false;
                }
            } else if("eyr" == req) {
                int eyr = std::stoi(item);
                if (eyr < 2020 || eyr > 2030) {
                    return false;
                }
            } else if("hgt" == req) {
                std::string hgt = item;
                if (hgt.find("in") != std::string::npos) {
                    int hgtin = std::stoi(item.substr(0, 2));
                    if (hgtin < 59 || hgtin > 76) {

                        return false;
                    }
                } else if (hgt.find("cm") != std::string::npos) {
                    int hgtcm = std::stoi(item.substr(0, 3));
                    if (hgtcm < 150 || hgtcm > 193) {

                        return false;
                    }
                } else {
                    return false;
                }
            } else if("hcl" == req) {
                std::string hcl = item;
                if (hcl.size() != 7) {
                    return false;
                }
                if (hcl[0] != '#') {
                    return false;
                }
                if (!std::regex_match (hcl.substr(1,6), std::regex("([a-f0-9])+"))) {
                    return false;
                }
            } else if("ecl" == req) {
                std::string ecl = item;
                if (std::find(valid_ecl.begin(), valid_ecl.end(), ecl) == valid_ecl.end()) {
                    return false;
                }
            } else if("pid" == req) {
                std::string pid = item;
                if (pid.size() != 9) {
                    return false;
                }
            }
        }
    }
    return true;
}



int validatePassports(std::vector<std::string> passports, std::vector<std::string> req_fields) {
    int valid = 0;
    std::map<std::string, std::string> current_passport;
    
    for (int i = 0; i < passports.size(); i++) {
        std::string row = passports[i];
        if (row.size() == 0) {
            // check to see if the passport is valid
            if (validPassport(current_passport, req_fields)) {
                valid++;
            }
            // it's time to reset the passport
            current_passport.clear();
            continue;
        }
        // otherwise, it's time to fill the map
        // eyr:2021 hgt:168cm hcl:#fffffd pid:180778832 byr:1923 ecl:amb iyr:2019 cid:241
        std::vector<std::string> splitSpace = splitString(row, " ");
        for(int i = 0; i < splitSpace.size(); i++) {
            std::string kvpair = splitSpace[i];
            if (kvpair.size() == 0) {
                break;
            }
            // std::cout << "i is " << i <<  " kvpair is " << kvpair << " split size is " << splitSpace.size() << std::endl;
            std::vector<std::string> splitColon = splitString(kvpair, ":");
            std::string key = splitColon[0];
            std::string val = splitColon[1];
            // std::cout << "j is " << j <<  " field is " << field << " split size is " << splitColon.size() << std::endl;
            if (std::find(req_fields.begin(), req_fields.end(), key) != req_fields.end())
            {
                current_passport[key] = val;
            }
        }
    };
    if ( current_passport.find("checked") == current_passport.end()) {
        // 
        if (validPassport(current_passport, req_fields)) {
            valid++;
        }
    } 
    return valid;
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc4.txt");

    int validCount = 0;
    std::vector<std::string> requiredFields = {
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid"
    };
    validCount = validatePassports(data, requiredFields);
    std::cout << "The number of validPassports are " << validCount << std::endl;
}