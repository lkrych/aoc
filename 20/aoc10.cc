#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

int joltageProduct(std::vector<long> data) {
    int diffOne = 0;
    int diffThree = 0;

    std::sort(data.begin(), data.end());
    
    // initial diff
    int initialDiff = data[0] - 0;
    if ( initialDiff == 1) {
        diffOne++;
    } else if ( initialDiff == 3) {
        diffThree++;
    }

    for (int i = 0; i < data.size() - 1; i++) {
        int curr = data[i];
        int next = data[i + 1];
        int diff = next - curr;
        if (diff == 1) {
            diffOne++;
        } else if (diff == 3) {
            diffThree++;
        }
    }
    // device's built-in adapter 
    diffThree++;
    return diffThree * diffOne;
}


int main() {
    //read input from file
    std::vector<long> data;
    data = getInputLong("./input/aoc10.txt");

    int joltage = joltageProduct(data);
    std::cout << "the joltage product is " << joltage << std::endl;
}