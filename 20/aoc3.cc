#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include "aoc_helper.h"

// // Part 1
// int treesEncountered(std::vector<std::string>data, std::vector<int> path) {
//     int horizontal = path[0];
//     int down = path[1];
//     int tree_count = 0;
//     std::vector<int> coordinates = {0,0};
//     for(int i = 0; i < data.size(); i++) {
//         int current_pos = coordinates[0];
//         std::string row = data[i];
//         if (row[current_pos] == '#') {
//             tree_count++;
//         }
//         //update coordinates
//         coordinates[0] = (current_pos + horizontal) % row.size();
//         coordinates[1] += 1;
//     }
//     return tree_count;
// }

// part 2 better/more general tree encountering algorithm
int treesEncountered(std::vector<std::string>data, std::vector<int> path) {
    int horizontal = path[0];
    int down = path[1];
    int tree_count = 0;
    std::vector<int> coordinates = {0,0};
    int i = 0;
    while(i < data.size()) {
        int current_pos = coordinates[0];
        std::string row = data[i];
        if (row[current_pos] == '#') {
            tree_count++;
        }
        //update coordinates
        coordinates[0] = (current_pos + horizontal) % row.size();
        i += down;
    }
    return tree_count;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc3.txt");

    // traverse the map and count the number of trees
    // this path vector means right 3, down 1
    std::vector<std::vector<int>> paths = {
        {1, 1},
        {3, 1},
        {5, 1},
        {7, 1},
        {1, 2}
    };

    int product = 1;
    for(int i = 0; i < paths.size(); i++) {
        std::vector<int> path = paths[i];
        int trees = treesEncountered(data, path);
        std::cout << "Encountered " << trees << " trees along the toboggan path " << path[0] << ", " << path[1] << std::endl;
        product *= trees;
    }
    std::cout << "Product of trees is " << product << std::endl;
}