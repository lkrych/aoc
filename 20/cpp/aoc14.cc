#include <algorithm>
#include <cmath>
#include <iostream>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include "aoc_helper.h"

// https://stackoverflow.com/a/58493878/4458404
std::string toBinary(int n) {
    if (n==0) return "0";
    else if (n==1) return "1";
    else if (n%2 == 0) return toBinary(n/2) + "0";
    else if (n%2 != 0) return toBinary(n/2) + "1";
}

//s needs to be reverse binary
long fromBinary (std::string s) {
    // std::cout << "translating " << s << " back to binary" << std::endl;
    long val = 0;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] == '1') { 
            val +=  pow(2, i);
        }
    }
    return val;
}

// long maskit(long val, std::string mask) {
//     // std::cout << "masking " << val << " with " << mask << std::endl;
//     std::vector<char> masked = {};
//     std::string binval = toBinary(val);
//     //pad binval with zeros
//     while(binval.size() < 36) {
//         binval = std::string("0").append( binval);
//     }
//     // std::cout << "binary value of " << val << " is " << binval << std::endl;
//     for (int i = 35; i >= 0; i--) {
//         if (mask[i] == 'X') {
//             masked.push_back(binval[i]);
//         } else {
//             masked.push_back(mask[i]);
//         }
//     }
//     std::string s(masked.begin(), masked.end());
//     long valbin = fromBinary(s);
//     return valbin;
// }

// part1
// long mask1(std::vector<std::string> data) {
//     long running_sum = 0;
//     std::string current_mask = "";
//     std::map<int,long> memory = {}; //memory_address -> value
//     // std::cout << "input is size " << data.size() << std::endl;
//     for(int i = 0; i < data.size(); i++) {
//         std::string instr = data[i];
//         // std::cout << "parsing " << instr << std::endl;
//         std::vector<std::string> split = splitString(instr, " ");
//         // std::cout << "split size " << split.size() << std::endl;
//         // there are two kinds of instructions
//         if (instr.substr(0,3) == "mem") {
//             //ex: mem[40278] = 36774405
//             long val = std::stol(split[2]);
//             std::string mem_address= splitString(split[0], "mem")[1];
//             mem_address = mem_address.substr(1, mem_address.size() - 2);
//             long masked_val = maskit(val, current_mask);
//             // std::cout << "masked_val = " << masked_val << std::endl;
//             memory[std::stoi(mem_address)] = masked_val;

//         } else if (instr.substr(0,4) == "mask") {
//             //ex: mask = 1X11X010X000X0X101X00100011X10100111
//             current_mask = split[2];
//         }
//     }

//     // sum all the values in memory
//     std::map<int, long>::iterator it;
//     for (it = memory.begin(); it != memory.end(); it++) {
//         running_sum += it->second;
//     }
//     return running_sum;
// }

std::vector<long> maskit2(long val, std::string mask) {
    std::vector<std::vector<char>> addresses = {{}};
    std::vector<char> masked = {};
    std::string binval = toBinary(val);
    //pad binval with zeros
    while(binval.size() < 36) {
        binval = std::string("0").append( binval);
    }
    // std::cout << "masking " << binval << std::endl;
    // std::cout << "with " << mask << std::endl;
    
    for (int i = 35; i >= 0; i--) {
        if (mask[i] == 'X') {
           // make a copy of the previous addresses and add 0, 1
            std::vector<std::vector<char>> new_addresses;
            for (int j = 0; j < addresses.size(); j++) {
                masked = addresses[j];
                std::vector<char> copy_of_masked(masked); 
                masked.push_back('1');
                copy_of_masked.push_back('0');
                new_addresses.push_back(copy_of_masked);
                addresses[j] = masked;
            }
            for(int j = 0; j < new_addresses.size(); j++) {
                masked = new_addresses[j];
                addresses.push_back(masked);
            }

        } else if (mask[i] == '1') {
            for (int j = 0; j < addresses.size(); j++) {
                masked = addresses[j];
                masked.push_back('1');
                addresses[j] = masked;
            }
        } else {
            for (int j = 0; j < addresses.size(); j++) {
                masked = addresses[j];
                masked.push_back(binval[i]);
                addresses[j] = masked;
            }
        }
    }

    std::vector<long> masked_addresses;
    for (int i = 0; i < addresses.size(); i++) {
        std::vector<char> address = addresses[i];
        std::string s(address.begin(), address.end());
        // std::cout << "address is " << s << std::endl;
        long valbin = fromBinary(s);
        // std::cout << "adding " << valbin << " to masked addresses" << std::endl;
        masked_addresses.push_back(valbin);
    }
    return masked_addresses;
}

long mask2(std::vector<std::string> data) {
    long running_sum = 0;
    std::string current_mask = "";
    std::map<long,long> memory = {}; //memory_address -> value
    for(int i = 0; i < data.size(); i++) {
        std::string instr = data[i];
        std::vector<std::string> split = splitString(instr, " ");
        // there are two kinds of instructions
        if (instr.substr(0,3) == "mem") {
            //ex: mem[40278] = 36774405
            long val = std::stol(split[2]);
            std::string mem_address= splitString(split[0], "mem")[1];
            mem_address = mem_address.substr(1, mem_address.size() - 2);
            std::vector<long> masked_vals = maskit2(std::stol(mem_address), current_mask);
            for (int j = 0; j < masked_vals.size(); j++) {
                long masked_val = masked_vals[j];
                memory[masked_val] = val;
            }

        } else if (instr.substr(0,4) == "mask") {
            //ex: mask = 1X11X010X000X0X101X00100011X10100111
            current_mask = split[2];
        }
    }

    // sum all the values in memory
    std::map<long, long>::iterator it;
    for (it = memory.begin(); it != memory.end(); it++) {
        running_sum += it->second;
    }
    return running_sum;
}

int main() {
    //read input from file
    std::vector<std::string> data;
    data = getInputString("./input/aoc14.txt");

    long maskSum = mask2(data);
    std::cout << "the mask sum is " << maskSum << std::endl;
}