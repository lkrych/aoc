#include <fstream>
#include <iostream>
#include <map>
#include <sstream>
#include <vector>

using namespace std;
//day 3 crossed wires
//What is the Manhattan distance from the central port to the closest intersection?
// ...........
// .+-----+...
// .|.....|...
// .|..+--X-+.
// .|..|..|.|.
// .|.-X--+.|.
// .|..|....|.
// .|.......|.
// .o-------+.
// ...........

void print_array (vector<string> arr) {
    int i;
    printf("arr size: %lu \n", arr.size());
    printf("[ ");
    for (i = 0; i < arr.size(); i++) {
        printf("%s ", arr[i].c_str());
    }
    printf("]\n");
}

vector<string> get_wire(int idx) {
    fstream newfile;
    vector<string> arr;

    newfile.open("aoc_day3.txt", ios::in);
    if (newfile.is_open()) {
        string line;
        //skip to correct line
        for ( int i = 1; i <= idx; i++) {
            getline(newfile, line);
        }
        //split string
        stringstream ss(line);
        string token;
        while (getline(ss, token, ',')) {
            arr.push_back(token);
        }
    }
    newfile.close();
    return arr;
}

map<int, vector<int> > add_to_map( map<int, vector<int> > m, int x, int y) {
    vector<int> v;
    if (m.find(x) == m.end() ) {
        v.push_back(y);
        m[x] = v;
    } else {
        v = m[x];
        v.push_back(y);
        m[x] = v;

    }
    return m;
}


map<int, vector<int> > map_wire_to_coordinates(vector<string> w) {
    // variables to map current position
    int x = 0;
    int y = 0;
    int length;
    string direction;
    // {x: [y1,y2,y3]}
    map<int, vector<int> > m;

    for (int i = 0; i < w.size(); i++) {
        //iterate through directions
        direction = w[i][0];
        length = stoi(w[i].substr(1));
        printf("d: %s, l: %d\n", direction.c_str(), length);
        //add wires to map
        if (direction == "U") {
            for (int i = y + 1; i < y + length; i++) {
                m = add_to_map(m, x, i);
            }
            y += length;
        } else if ( direction == "L") {
            for (int j = x - 1; j > x - length; j--) {
                m = add_to_map(m, j, y);
            }
            x -= length;
        } else if ( direction == "R") {
            for (int j = x + 1; j < x + length; j++) {
                m = add_to_map(m, j, y);
            }
            x += length;
        } else if ( direction == "D") {
            for (int i = y - 1; i > y - length; i--) {
                m = add_to_map(m, x, i);
            }
            y -= length;
        } 
    }
    return m;
}

void check_if_intersection(map<int, vector<int> > m1, map<int, vector<int> > m2, int x, int y) {
    vector<int> v1;
    vector<int> v2;

    if (m1.find(x) == m1.end()) {
        return;
    } else {
        if (m2.find(x) == m2.end()) {
            return;
        } else {
            //check if y exists in both
            v1 = m1[x];
            v2 = m2[x];
            if(find(v1.begin(), v1.end(), y) != v1.end()){
                if(find(v2.begin(), v2.end(), y) != v2.end()){
                    printf("Found intersection at x: %d, y: %d", x, y);
                }
            } 
        }
    }
}

void find_closest_intersection(map<int, vector<int> > m1, map<int, vector<int> > m2) {
    // the idea of this function is to start at the origin
    // and to slowly walk outwards from it
    // the first intersection found will be the closest
    // start at top
    // [0,1],[1,1],[1,0],[1,-1],[0,-1],[-1,-1], [-1,0], [-1,1]
    int x = 0;
    int y = 0;
    int d = 1;
    int m = 1;
    
    while(x < 100) {
        while (2 * x * d < m) {
            check_if_intersection(m1, m2, x,y);
            x = x + d;
        }
        while (2 * y * d < m) {
            check_if_intersection(m1, m2, x,y);
            y = y + d;
        }
        d = -1 * d;
        m = m + 1;
    }
}


int main() {
    vector<string> w1;
    vector<string> w2;
    map<int, vector<int> > m1;
    map<int, vector<int> > m2;

    w1 = get_wire(1);
    w2 = get_wire(2);

    m1 = map_wire_to_coordinates(w1);
    m2 = map_wire_to_coordinates(w2);

    find_closest_intersection(m1, m2);

}