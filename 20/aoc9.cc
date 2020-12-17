#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

std::vector<std::pair<int, int>> twoSum(std::vector<int> arr, int target) {
    std::vector<std::pair<int, int>> vals = {};
    std::map<int, int> hash;
    //iterate through array and try to insert the element into the map, 
    //checking to see if it's inverse already exists
    for(int i = 0; i < arr.size(); i++) {
        int el = arr[i];
        int diff = target - el;
        // first check to see if the diff exists, if it does, then these two values are our answer
        if ( hash.find(diff) == hash.end() ) {
        // we add the el to the map
            hash[el] = 1;
        } else {
        // found
            vals.push_back(std::pair<int,int>(el, diff));
        }
    }
    return vals;
}

std::vector<int> modifiedTwoSum(std::vector<int> arr, int idx, int window, std::vector<int> exclude) {
    int target = arr[idx + window];
    std::vector<int> subArr = {arr.begin() + idx, arr.begin() + idx + (window - 1)};
    std::vector<std::pair<int,int>> vals = twoSum(subArr, target);
    if (vals.size() == 0) {
        return {};
    }
    // else check to make sure none of the values are excluded
    for(int i = 0; i < vals.size(); i++) {
        std::pair<int,int> pair = vals[i];
        int first = pair.first;
        int second = pair.second;
        for (int j = 0; j < exclude.size(); j++) {
            int excl = exclude[j];
            if (first == excl || second == excl) {
                // remove the pair from vals
                vals.erase(vals.begin() + i);
            }
        }
    }
    return { vals[0].first, vals[0].second };
}

int crackXmas(std::vector<int> data, int window_size) {
    int curr_idx = 0;
    std::vector<int> exclude = {};
    while (true) {
        std::vector<int> indices = modifiedTwoSum(data, curr_idx, window_size, exclude);
        if (indices.size() == 0) {
            return data[curr_idx + window_size];
        }
        for(int i = 0; i < indices.size(); i++) {
            exclude.push_back(indices[i]);
        }
        curr_idx++;
    }
}

int main() {
    //read input from file
    std::vector<int> data;
    data = getInputInt("./input/aoc9.txt");

    int xmasFault = crackXmas(data, 25);
    std::cout << "the value is " << xmasFault << std::endl;
}