#include <iostream>
#include <fstream>
using namespace std;

//Fuel required to launch a given module is based on its mass. 
//Specifically, to find the fuel required for a module, take its mass, 
//divide by three, round down, and subtract 2.

//part 1
int calculate_fuel_p1(int total, int module) {
    total += ((module / 3) - 2);
    return total;
}

//part 2
int calculate_fuel_p2(int total, int module) {
    while (module > 0) {
        module = ((module / 3) - 2);
        if (module > 0) {
            total += module;
        }
    }
    return total;
}

int main() {
    //create an object newfile with fstream
    fstream newfile; 
    newfile.open("aoc_day1.txt",ios::in);
    if (newfile.is_open()) {
        string line;
        int module;
        int total_fuel;
        while (getline(newfile, line)) {
            // printf("%s\n", line.c_str());
            module = stoi(line);
            total_fuel = calculate_fuel_p2(total_fuel, module);
        }
        newfile.close();
        printf("Total fuel needed is %d\n", total_fuel);
    }

}