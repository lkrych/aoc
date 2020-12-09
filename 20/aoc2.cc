#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include "aoc_helper.h"

bool invalidPassword(std::string rules, std::string target, std::string pass) {
    int lowerBound, upperBound;
    std::vector<std::string> split;
    split = splitString(rules, "-");
    lowerBound = std::stoi(split[0]);
    upperBound = std::stoi(split[1]);
    int count = 0;
    for (int i = 0; i < pass.size(); i++) {
        char c = pass[i];
        if (c == target[0]) {
            count++;
        }
    }
    return count >= lowerBound && count <= upperBound
}

std::vector<std::string> invalidPasswords(std::vector<std::string> arr) {
    std::string entry, password, rules, target;
    std::vector<std::string> inv_passwords;
    for(int i = 0; i < arr.size(); i++) {
        std::vector<std::string> split;
        entry = arr[i];
        // 1-3 a: abcde
        split = splitString(entry, " ");
        rules = split[0];
        target = split[1][0];
        password= split[2];
        if (invalidPassword(rules, target, password)) {
            inv_passwords.push_back(password);
        }

    }
    return inv_passwords;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc2.txt");

    // use three sum to find the three matching nums
    std::vector<std::string> invalid_passwords = invalidPasswords(data);
    std::cout << "There are " << invalid_passwords.size() << " invalid passwords" << std::endl;
}