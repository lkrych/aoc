#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// part 1
// int joltageProduct(std::vector<long> data) {
//     int diffOne = 0;
//     int diffThree = 0;

//     std::sort(data.begin(), data.end());
    
//     // initial diff
//     int initialDiff = data[0] - 0;
//     if ( initialDiff == 1) {
//         diffOne++;
//     } else if ( initialDiff == 3) {
//         diffThree++;
//     }

//     for (int i = 0; i < data.size() - 1; i++) {
//         int curr = data[i];
//         int next = data[i + 1];
//         int diff = next - curr;
//         if (diff == 1) {
//             diffOne++;
//         } else if (diff == 3) {
//             diffThree++;
//         }
//     }
//     // device's built-in adapter 
//     diffThree++;
//     return diffThree * diffOne;
// }

std::map<int,long> calculateCombos(int el, std::map<int,long> adapter_combos) {
    long combos = 0;
    int diffOne = el - 1;
    int diffTwo = el - 2;
    int diffThree = el - 3;

    // std::cout << "el: " << el << " search for " << diffOne << " and " << diffTwo << " and " << diffThree << std::endl;
    
    if (adapter_combos.find(diffOne) != adapter_combos.end()) {
        combos += adapter_combos[diffOne];
    }
    if (adapter_combos.find(diffTwo) != adapter_combos.end()) {
        combos += adapter_combos[diffTwo];
    }
    if (adapter_combos.find(diffThree) != adapter_combos.end()) {
        combos +=  adapter_combos[diffThree];;
    }

    adapter_combos[el] = combos;
    return adapter_combos;
}

long joltageCombination(std::vector<long> data) {
    long combos;
    std::map<int, long> adapter_combos;
    adapter_combos[0] = 1;

    std::sort(data.begin(), data.end());

    for (int i = 0; i < data.size(); i++) {
        int el = data[i];
        adapter_combos = calculateCombos(el, adapter_combos);
    }

    int max = data[data.size() - 1] + 3;
    adapter_combos = calculateCombos(max, adapter_combos);

    return adapter_combos[max];
}


int main() {
    //read input from file
    std::vector<long> data;
    data = getInputLong("./input/aoc10.txt");

    long joltageComb = joltageCombination(data);
    std::cout << "the joltage product is " << joltageComb << std::endl;
}