#include <algorithm>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// part 1
// std::map<std::string, std::vector<std::string>> parseBags(std::string row) {
//     std::map<std::string, std::vector<std::string>> entry;
//     // vibrant aqua bags contain 1 shiny magenta bag, 2 muted teal bags, 1 dim magenta bag, 1 muted chartreuse bag.
//     std::vector<std::string> splitByContain = splitString(row, "bags contain");
//     if (splitByContain.size() != 2) {
//         std::cout << "There was a problem with your input data: " << row << std::endl;
//         exit(1);
//     }
//     std::string parent = strip(splitByContain[0]);
//     // std::cout << "container: " << key << " size: " << key.size() << std::endl;
//     entry[parent] = {};
//     //parse the containing bags
//     std::vector<std::string> splitByCommas = splitString(splitByContain[1], ",");
//     for (int i = 0; i < splitByCommas.size(); i++) {
//         std::string unparsed = splitByCommas[i];
//         std::string noBags = removeFromString(unparsed, {"bags", "bag", "."});
//         std::string stripped = strip(noBags);
//         // std::cout << stripped << " size: " << stripped.size() << std::endl;
//         if (noBags.find("no other") != std::string::npos) {
//             continue;
//         }
//         if ( entry.find(stripped) == entry.end() ) {
//             // not found
//                 entry[stripped] = {parent};
//         } else {
//         // found
//             entry[stripped].push_back(parent);
//         }
//     }
//     return entry;
// }

// part 2
std::map<std::string, std::vector<std::pair<std::string, int>>> parseBags(std::string row) {
    std::map<std::string, std::vector<std::pair<std::string, int>>> entry;
    // vibrant aqua bags contain 1 shiny magenta bag, 2 muted teal bags, 1 dim magenta bag, 1 muted chartreuse bag.
    std::vector<std::string> splitByContain = splitString(row, "bags contain");
    if (splitByContain.size() != 2) {
        std::cout << "There was a problem with your input data: " << row << std::endl;
        exit(1);
    }
    std::string parent = strip(splitByContain[0]);
    // std::cout << "container: " << key << " size: " << key.size() << std::endl;
    entry[parent] = {};
    //parse the containing bags
    std::vector<std::string> splitByCommas = splitString(splitByContain[1], ",");
    for (int i = 0; i < splitByCommas.size(); i++) {
        std::string unparsed = splitByCommas[i];
        std::string noBags = removeFromString(unparsed, {"bags", "bag", "."});
        std::string stripped = strip(noBags);
        // https://stackoverflow.com/a/30727561/4458404
        int n = stripped[0] - '0'; // convert to int
        stripped = strip(stripped.substr(1, stripped.size()- 1));
        
        if (noBags.find("no other") != std::string::npos) {
            continue;
        }
        if ( entry.find(stripped) == entry.end() ) {
            // not found
                entry[stripped] = {std::pair<std::string,int> (parent, n)};
        } else {
        // found
            entry[stripped].push_back(std::pair<std::string,int> (parent, n));
        }
    }
    return entry;
}

int parentBags( std::map<std::string, std::vector<std::pair<std::string, int>>> adjacency_list, std::string target, std::map<std::string, bool> *parent_check) {
    int count = 0;
    std::vector<std::pair<std::string, int>> parents = adjacency_list[target];

    for (int i = 0; i < parents.size(); i++) {
        std::string parent = parents[i].first;
        if (parent_check->find(parent) == parent_check->end()) {
            count++;
            (*parent_check)[parent] = true;
            count += parentBags(adjacency_list, parent, parent_check);
        }
    }
    
    return count;
}

int baggageClaim(std::vector<std::string> arr, std::string target) {
    std::string row;
    std::map<std::string, std::vector<std::pair<std::string, int>>> adjacency_list;
    for(int i = 0; i < arr.size(); i++) {
        row = arr[i];
        std::map<std::string, std::vector<std::pair<std::string, int>>> entry = parseBags(row);
        std::map<std::string, std::vector<std::pair<std::string, int>>>::iterator it;
        for ( it = entry.begin(); it != entry.end(); it++ )
        {
            if ( adjacency_list.find(it->first) == adjacency_list.end() ) {
            // not found
                adjacency_list[it->first] = it->second;
            } else {
            // found. add elements from entry to the adjacency list
                adjacency_list[it->first].insert(adjacency_list[it->first].end(), it->second.begin(), it->second.end());
            }
        }
    }    
    std::map<std::string, bool> parent_check;
    int bags = parentBags(adjacency_list, target, &parent_check);
    return bags;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc7.txt");

    int parentBags = baggageClaim(data, "shiny gold");
    std::cout << "the number of parent bags " << parentBags << std::endl;
}