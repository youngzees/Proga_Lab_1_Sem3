#include "forward_list.hpp"
#include <sstream>
using namespace std;

ForwardList::~ForwardList() { clear(); }

void ForwardList::clear() {
    while (head) {
        NodeF* t = head;
        head = head->next;
        delete t;
    }
}

//      добавление

void ForwardList::push_head(int val) {
    NodeF* n = new NodeF(val);
    n->next = head;
    head = n;
}

void ForwardList::push_tail(int val) {
    NodeF* n = new NodeF(val);
    if (!head) { head = n; return; }
    NodeF* t = head;
    while (t->next) t = t->next;
    t->next = n;
}

bool ForwardList::push_before(int target, int val) {
    if (!head) return false;
    if (head->val == target) { push_head(val); return true; }
    NodeF* t = head;
    while (t->next && t->next->val != target) t = t->next;
    if (!t->next) return false;
    NodeF* n = new NodeF(val);
    n->next = t->next;
    t->next = n;
    return true;
}

bool ForwardList::push_after(int target, int val) {
    NodeF* t = head;
    while (t && t->val != target) t = t->next;
    if (!t) return false;
    NodeF* n = new NodeF(val);
    n->next = t->next;
    t->next = n;
    return true;
}

//      удаление

bool ForwardList::pop_head() {
    if (!head) return false;
    NodeF* t = head;
    head = head->next;
    delete t;
    return true;
}

bool ForwardList::pop_tail() {
    if (!head) return false;
    if (!head->next) {
        delete head;
        head = nullptr;
        return true;
    }
    NodeF* t = head;
    while (t->next && t->next->next) t = t->next;
    delete t->next;
    t->next = nullptr;
    return true;
}

bool ForwardList::del(int val) {
    if (!head) return false;
    if (head->val == val) return pop_head();
    NodeF* t = head;
    while (t->next && t->next->val != val) t = t->next;
    if (!t->next) return false;
    NodeF* toDel = t->next;
    t->next = t->next->next;
    delete toDel;
    return true;
}

bool ForwardList::del_index(int idx) {
    if (idx < 0 || !head) return false;
    if (idx == 0) return pop_head();
    NodeF* t = head;
    for (int i = 0; i < idx - 1 && t->next; ++i) t = t->next;
    if (!t->next) return false;
    NodeF* toDel = t->next;
    t->next = t->next->next;
    delete toDel;
    return true;
}

bool ForwardList::find(int val) const {
    for (NodeF* t = head; t; t = t->next)
        if (t->val == val) return true;
    return false;
}

//      сериализация

string ForwardList::serialize() const {
    ostringstream os;
    bool first = true;
    for (NodeF* t = head; t; t = t->next) {
        if (!first) os << ' ';
        os << t->val;
        first = false;
    }
    return os.str();
}

void ForwardList::deserialize(const string& s) {
    clear();
    istringstream is(s);
    int v;
    while (is >> v) push_tail(v);
}

string ForwardList::to_string() const {
    ostringstream os;
    os << "[ ";
    for (NodeF* t = head; t; t = t->next) os << t->val << ' ';
    os << ']';
    return os.str();
}