#include "array.hpp"
#include <sstream>
using namespace std;

void MyArray::push(int val) { data.push_back(val); }

bool MyArray::del(int index) {
    if (index < 0 || index >= (int)data.size()) return false;
    data.erase(data.begin() + index);
    return true;
}

int MyArray::get(int index) const {
    if (index < 0 || index >= (int)data.size()) return -1;
    return data[index];
}

void MyArray::set(int index, int val) {
    if (index >= 0 && index < (int)data.size()) data[index] = val;
}

int MyArray::length() const { return (int)data.size(); }
void MyArray::clear() { data.clear(); }

string MyArray::serialize() const {
    ostringstream os;
    for (size_t i = 0; i < data.size(); ++i) {
        if (i) os << ' ';
        os << data[i];
    }
    return os.str();
}

void MyArray::deserialize(const string& s) {
    data.clear();
    istringstream is(s);
    int v;
    while (is >> v) data.push_back(v);
}