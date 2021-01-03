#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>

std::vector<int> getInput(std::string filename) {
    std::ifstream myfile(filename);
    std::vector<int> input;
    std::string line;
    if (myfile.is_open()) {
        while ( getline (myfile, line) ){
            int el = std::stoi(line);
            input.push_back(el);
        }
        myfile.close();
    } else {
        std::cout << "There was a problem reading from " << filename;
    }
    return input;
}

std::vector<int> twoSum(std::vector<int> arr, int target) {
    std::vector<int> vals = {};
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
            vals.push_back(el);
            vals.push_back(diff);
            return vals;
        }
    }
    return vals;
}

std::vector<int> threeSum(std::vector<int> arr, int target) {
    //init three pointers
    int l = 0;
    int m = 1;
    int r = arr.size() - 1;
    std::vector<int> vals = {};

    //first sort from smallest to largest
    std::sort (arr.begin(), arr.end());

    while (true) {
        while (m < r) {
            std::cout << "l: " << l << " m: " << m << " r: " << r << std::endl;
            std::cout << "l: " << arr[l] << " m: " << arr[m] << " r: " << arr[r] << std::endl;
            int current_sum = arr[l] + arr[m] + arr[r];
            std::cout << "sum: " << current_sum << std::endl;
            if (current_sum == target) {
                vals.push_back(arr[l]);
                vals.push_back(arr[m]);
                vals.push_back(arr[r]);
                return vals;
            }
            if (current_sum > target) {
                // we've exhausted the choices for the rest of this window because the current sum is too large, 
                // we need to reset it by bringing the upper-bound down
                // or if there is a difference between middle and lower
                // bring the lower-bound up
                if (m - l > 1) {
                    l += 1;
                    m = l + 1;
                } else {
                    m = l + 1;
                    r = r - 1;
                }
            } else {
                m += 1;
            }
        }
        // we've exhausted all the possibilities in the window
        // we need to reset it by bringing the lower-bound up
        l += 1;
        m = l + 1;
        if (l >= r) {
            //there were no matching cases found, return an empty array
            return vals;
        }
    }
    return vals;
}

// Part 1
// int main() {
//     //read input from file
//     std::vector<int> data;
//     data = getInput("./input/aoc1.txt");

//     // use two sum to find the two matching nums
//     std::vector<int> vals = twoSum(data, 2020);
//     if (vals.size() > 0) {
//         std::cout << "The two found vals were " << vals[0] << " and " << vals[1] << std::endl;
//         std::cout << "Their product is " << vals[0] * vals[1] << std::endl;
//     }
// }

// Part 2
int main() {
    //read input from file
    std::vector<int> data;
    data = getInput("./input/aoc1.txt");

    // use three sum to find the three matching nums
    std::vector<int> vals = threeSum(data, 2020);
    if (vals.size() > 0) {
        std::cout << "The three found vals were " << vals[0] << " and " << vals[1] << " and " << vals[2] << std::endl;
        std::cout << "Their product is " << vals[0] * vals[1] * vals[2] << std::endl;
    } else {
        std::cout << "There was an error in threeSum!" << std::endl;
    }
}