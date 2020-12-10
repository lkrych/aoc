#include <algorithm>
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include "aoc_helper.h"

bool validPassport(std::map<std::string, bool> passport, std::vector<std::string> required) {
    // check if all the required fields are in the map
    std::cout << "checking if passport with " << passport.size() <<" keys is valid" << std::endl;
    for (int i = 0; i < required.size(); i++) {
        std::string req = required[i];
        if ( passport.find(req) == passport.end()) {
            return false;
        } 
    }
    std::cout << "passport is valid!" << std::endl;
    return true;
}

int validatePassports(std::vector<std::string> passports, std::vector<std::string> req_fields) {
    int valid = 0;
    std::cout << "there are at most " << passports.size() << " passports to check" << std::endl;
    std::map<std::string, bool> current_passport;
    
    std::cout << "initialized map of size " << current_passport.size() << std::endl;
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
            std::string field = splitColon[0];
            // std::cout << "j is " << j <<  " field is " << field << " split size is " << splitColon.size() << std::endl;
            if (std::find(req_fields.begin(), req_fields.end(), field) != req_fields.end())
            {
                current_passport[field] = true;
            }
        }
    };
    if (validPassport(current_passport, req_fields)) {
        valid++;
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