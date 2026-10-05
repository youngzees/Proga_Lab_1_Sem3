#include "queue.hpp"
#include <sstream>
using namespace std;

void MyQueue::push(int val) { data.push_back(val); }

int MyQueue::pop() {
    if (data.empty()) return -1;
    int v = data.front();
    data.pop_front();
    return v;
}

int MyQueue::front() const { return data.empty() ? -1 : data.front(); }
bool MyQueue::empty() const { return data.empty(); }
int MyQueue::size() const { return (int)data.size(); }

string MyQueue::serialize() const {
    ostringstream os;
    for (size_t i = 0; i < data.size(); ++i) {
        if (i) os << ' ';
        os << data[i];
    }
    return os.str();
}

void MyQueue::deserialize(const string& s) {
    data.clear();
    istringstream is(s);
    int v;
    while (is >> v) data.push_back(v);
}

string MyQueue::to_string() const {
    ostringstream os;
    os << "[ front -> ";
    for (int v : data) os << v << ' ';
    os << ']';
    return os.str();
}