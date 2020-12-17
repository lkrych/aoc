#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// part1
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

// part1
std::vector<int> modifiedTwoSum(std::vector<long> arr, int idx, int window) {
    int target = arr[idx + window];
    std::vector<int> subArr = {arr.begin() + idx, arr.begin() + idx + window};
    
    std::vector<std::pair<int,int>> vals = twoSum(subArr, target);
    if (vals.size() == 0) {
        return {};
    }
    return { vals[0].first, vals[0].second };
}

// part1
int crackXmas(std::vector<long> data, int window_size) {
    int curr_idx = 0;
    while (true) {

        std::vector<int> indices = modifiedTwoSum(data, curr_idx, window_size);
        if (indices.size() == 0) {
            return data[curr_idx + window_size];
        }
        
        curr_idx++;
    }
}

long sumRange(std::vector<long> data, int low, int high) {
    long sum = 0;
    for (int i = low; i <= high; i++) {
        long el = data[i];
        sum += el;
    }
    return sum;
}

int crackXmas2(std::vector<long> data, long target) {
    int low_idx = 0;
    int high_idx = 0;
    while(true) {
        long sum = sumRange(data, low_idx, high_idx);
        if (sum == target) {
            long min = *std::min_element(data.begin() + low_idx, data.begin() + high_idx);
            long max = *std::max_element(data.begin() + low_idx, data.begin() + high_idx);
            return min + max;
        }
        
        if (sum < target) {
            high_idx++;
        } else if ( sum > target) {
            low_idx++;
        }
    }
}

int main() {
    //read input from file
    std::vector<long> data;
    data = getInputLong("./input/aoc9.txt");

    long xmasFault = crackXmas2(data, 22477624);
    std::cout << "the value is " << xmasFault << std::endl;
}