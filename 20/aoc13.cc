#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

int differenceFromTarget(int x, int y) {
    int mult = 1;
    int product = 1;
    while (product < y) {
        product = x * mult;
        mult++;
    }
    return product - y;
}

long long foundTarget(int val, long long target, std::map<int,int> memo) {
    long long mult = 1;
    long long product = 1;
    if (memo.find(val) != memo.end()) {
        mult = memo[val] / val;
        product = memo[val];
    } 
    while (product < target) {
        product = val * mult;
        if (target == 5398493001) {
            std::cout << "checking product " << product << std::endl;
        }
        mult++;
    }
    return product;
}

long long findOverlap(long current_product, int curr_element, int curr_diff, long lcm) {
    long long lcm_counter = lcm;
    bool foundOverlap = false;
    std::map<int,int> memo;
    long long target;
    while (true) {
        if (current_product == 1) {
            
            target = lcm_counter + curr_diff;
            if (target == 1) {
                lcm_counter += lcm;
                continue;
            }
            std::cout << "Searching for target " << target << std::endl;
            long long product = foundTarget(curr_element, target, memo);
            std::cout << "Found " << product << std::endl;
            if (product == target) {
                return product - curr_diff;
            }
        } else {
            target = current_product + lcm_counter + curr_diff;
        }
        if (target % curr_element == 0) {
            return target - curr_diff;
        }
        lcm_counter += lcm;
    }
}

//part1
// int busSchedule(std::vector<std::string> data) {
//     int departure = std::stoi(data[0]);
//     int leastDiffVal = 1000000;
//     int leastBusId = 1;
//     std::vector<std::string> buses = splitString(data[1], ",");
//     for (int i = 0; i < buses.size(); i++) {
//         std::string bus = buses[i];
//         if (bus == "x") {
//             continue;
//         }
//         int busi = std::stoi(bus);
//         int diff = differenceFromTarget(busi, departure);
//         std::cout << "The diff of bus " << busi << " is " << diff << std::endl;
//         if (diff < leastDiffVal) {
//             leastDiffVal = diff;
//             leastBusId = busi;
//         }
//     }
//     return leastBusId * leastDiffVal;
// }

long busSchedule2(std::vector<std::string> data) {
    std::vector<std::pair<int, int>> busIdAndDiff = {};
    std::vector<std::string> buses = splitString(data[1], ",");
    for (int i = 0; i < buses.size(); i++) {
        std::string bus = buses[i];
        if (bus == "x") {
            continue;
        }
        int busi = std::stoi(bus);
        busIdAndDiff.push_back(std::pair<int,int>(busi, i));
    }

    // the inputs to this algorithm are all prime numbers.
    // this means that the LCM of the two numbers are going to be the two numbers
    // multiplied by themselves.
    // to solve this algorithm, we need to find the first place where the two numbers
    // overlap, once we do this, we know that this pattern will repeat itself LCM(times later)
    // we can check this increment until the next bus overlaps, then find the LCM for those three or x nums
    // repeat this process until done. 
    long long prod = 1;
    // find overlap with x, y, and LCM
    // calculate LCM
    int updateTime = 1;
    long long running_lcm = busIdAndDiff[0].first;
    for (int i = 1; i < busIdAndDiff.size(); i++) {
        std::pair<int,int> current_elem = busIdAndDiff[i];
        prod = findOverlap(prod, current_elem.first, current_elem.second, running_lcm);
        std::cout << "updating prod to " << prod << std::endl;
        std::cout << updateTime << " time to update" << std::endl;
        updateTime++;
        //update lcm
        running_lcm *= current_elem.first;
        std::cout << "updating lcm to " << running_lcm << std::endl;
    }

    return prod;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc13.txt");

    long long busMult = busSchedule2(data);
    std::cout << "the bus riddle is " << busMult << std::endl;
}