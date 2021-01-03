#ifndef AOC_HELPER_H 
#define AOC_HELPER_H

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

std::vector<std::string> getInputString(std::string filename) {
    std::ifstream myfile(filename);
    std::vector<std::string> input;
    std::string line;
    if (myfile.is_open()) {
        while ( getline (myfile, line) ){
            input.push_back(line);
        }
        myfile.close();
    } else {
        std::cout << "There was a problem reading from " << filename;
    }
    return input;
}

std::vector<long> getInputLong(std::string filename) {
    std::vector<long> input;
    std::vector<std::string> stringInput;
    
    stringInput = getInputString(filename);
    
    for(int i = 0; i < stringInput.size(); i++) {
        int el = std::stol(stringInput[i]);
        input.push_back(el);
    }
    return input;
}

std::vector<std::string> splitString(std::string s, std::string delim) {
    std::string token;
    size_t pos = 0;
    std::vector<std::string> split;
    while ((pos = s.find(delim)) != std::string::npos) {
        token = s.substr(0, pos);
        split.push_back(token);
        s.erase(0, pos + delim.length());
    }
    split.push_back(s);
    return split;
}

// https://thispointer.com/how-to-remove-substrings-from-a-string-in-c/
void eraseAllSubStr(std::string & mainStr, const std::string & toErase)
{
    size_t pos = std::string::npos;
    // Search for the substring in string in a loop untill nothing is found
    while ((pos  = mainStr.find(toErase) )!= std::string::npos)
    {
        // If found then erase it from string
        mainStr.erase(pos, toErase.length());
    }
}

std::string removeFromString(std::string s, std::vector<std::string> targets) {
    for(int i = 0; i < targets.size(); i++) {
        std::string target = targets[i];
        eraseAllSubStr(s, target);
    }
    return s;
}

std::string lstrip(std::string s) {
    int idx = 0;
    while (!isalnum(s[idx])) {
        idx++;
    }
    return s.substr(idx, s.size() - idx);
}

std::string rstrip(std::string s) {
    int last_idx = s.size() - 1;
    while (!isalpha(s[last_idx])) {
        last_idx--;
    }
    return s.substr(0,last_idx + 1);
}

std::string strip(std::string s) {
    std::string strip = lstrip(s);
    strip = rstrip(strip);
    return strip;
}

std::vector<std::vector<char>> copyMatrix(std::vector<std::vector<char>> arr) {
    std::vector<std::vector<char>> matrix = {};
    for(int i = 0; i < arr.size(); i++) {
        std::vector<char> row = arr[i];
        std::vector<char> copy = {row.begin(), row.end()};
        matrix.push_back(copy);
    }
    return matrix;
}

// https://stackoverflow.com/a/38251268/4458404
int mod(int a, int b) { return (a % b + b) % b; }

#endif 