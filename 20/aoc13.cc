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

int busSchedule(std::vector<std::string> data) {
    int departure = std::stoi(data[0]);
    int leastDiffVal = 1000000;
    int leastBusId = 1;
    std::vector<std::string> buses = splitString(data[1], ",");
    for (int i = 0; i < buses.size(); i++) {
        std::string bus = buses[i];
        if (bus == "x") {
            continue;
        }
        int busi = std::stoi(bus);
        int diff = differenceFromTarget(busi, departure);
        std::cout << "The diff of bus " << busi << " is " << diff << std::endl;
        if (diff < leastDiffVal) {
            leastDiffVal = diff;
            leastBusId = busi;
        }
    }
    return leastBusId * leastDiffVal;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc13.txt");

    int busMult = busSchedule(data);
    std::cout << "the bus riddle is " << busMult << std::endl;
}