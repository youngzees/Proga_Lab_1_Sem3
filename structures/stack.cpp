#include "stack.hpp"
#include <sstream>
using namespace std;

void MyStack::push(int val) { data.push_back(val); }

int MyStack::pop() {
    if (data.empty()) return -1;
    int v = data.back();
    data.pop_back();
    return v;
}

int MyStack::top() const { return data.empty() ? -1 : data.back(); }
bool MyStack::empty() const { return data.empty(); }
int MyStack::size() const { return (int)data.size(); }

string MyStack::serialize() const {
    ostringstream os;
    for (size_t i = 0; i < data.size(); ++i) {
        if (i) os << ' ';
        os << data[i];
    }
    return os.str();
}

void MyStack::deserialize(const string& s) {
    data.clear();
    istringstream is(s);
    int v;
    while (is >> v) data.push_back(v);
}

string MyStack::to_string() const {
    ostringstream os;
    os << "[ top: ";
    for (auto it = data.rbegin(); it != data.rend(); ++it) os << *it << ' ';
    os << ']';
    return os.str();
}