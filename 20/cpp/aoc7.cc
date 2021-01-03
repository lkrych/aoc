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
        if ( entry.find(parent) == entry.end() ) {
            // not found
                entry[parent] = {std::pair<std::string,int> (stripped, n)};
        } else {
        // found
            entry[parent].push_back(std::pair<std::string,int> (stripped, n));
        }

    }
    return entry;
}

int childrenBags( std::map<std::string, std::vector<std::pair<std::string, int>>> adjacency_list, std::string target, int bagN) {
    int count = 0;
    std::vector<std::pair<std::string, int>> children = adjacency_list[target];
    std::cout << target << " count is " << bagN << std::endl;
    count += bagN;
    for (int i = 0; i < children.size(); i++) {
        std::string child = children[i].first;
        std::cout << "There are " << children[i].second << " " << child << " bags in a " << target << std::endl;
        int nestedBags = childrenBags(adjacency_list, child, children[i].second);
        std::cout << target << " contains " <<  nestedBags << std::endl;
        count +=  bagN * nestedBags;
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
            adjacency_list[it->first] = it->second;
        }
    }    
    std::map<std::string, bool> parent_check;
    int bags = childrenBags(adjacency_list, target, 1);
    return bags;
}


int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc7.txt");

    int parentBags = baggageClaim(data, "shiny gold");
    //need to subtract one to not count the target
    std::cout << "the number of parent bags " << parentBags - 1 << std::endl;
}