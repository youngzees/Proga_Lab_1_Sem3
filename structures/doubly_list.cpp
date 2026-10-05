#include "doubly_list.hpp"
#include <sstream>
using namespace std;

DoublyList::~DoublyList() { clear(); }

void DoublyList::clear() {
    while (head) {
        NodeL* t = head;
        head = head->next;
        delete t;
    }
    tail = nullptr;
}

//      добавление

void DoublyList::push_head(int val) {
    NodeL* n = new NodeL(val);
    if (!head) { head = tail = n; return; }
    n->next = head;
    head->prev = n;
    head = n;
}

void DoublyList::push_tail(int val) {
    NodeL* n = new NodeL(val);
    if (!tail) { head = tail = n; return; }
    tail->next = n;
    n->prev = tail;
    tail = n;
}

bool DoublyList::push_before(int target, int val) {
    NodeL* t = head;
    while (t && t->val != target) t = t->next;
    if (!t) return false;
    NodeL* n = new NodeL(val);
    n->prev = t->prev;
    n->next = t;
    if (t->prev) t->prev->next = n;
    else         head = n;
    t->prev = n;
    return true;
}

bool DoublyList::push_after(int target, int val) {
    NodeL* t = head;
    while (t && t->val != target) t = t->next;
    if (!t) return false;
    NodeL* n = new NodeL(val);
    n->next = t->next;
    n->prev = t;
    if (t->next) t->next->prev = n;
    else         tail = n;
    t->next = n;
    return true;
}

//      удаление

bool DoublyList::pop_head() {
    if (!head) return false;
    NodeL* t = head;
    head = head->next;
    if (head) head->prev = nullptr;
    else      tail = nullptr;
    delete t;
    return true;
}

bool DoublyList::pop_tail() {
    if (!tail) return false;
    NodeL* t = tail;
    tail = tail->prev;
    if (tail) tail->next = nullptr;
    else      head = nullptr;
    delete t;
    return true;
}

bool DoublyList::del(int val) {
    for (NodeL* t = head; t; t = t->next) {
        if (t->val != val) continue;
        if (t->prev) t->prev->next = t->next;
        else         head = t->next;
        if (t->next) t->next->prev = t->prev;
        else         tail = t->prev;
        delete t;
        return true;
    }
    return false;
}

bool DoublyList::del_index(int idx) {
    if (idx < 0) return false;
    NodeL* t = head;
    for (int i = 0; i < idx && t; ++i) t = t->next;
    if (!t) return false;
    if (t->prev) t->prev->next = t->next;
    else         head = t->next;
    if (t->next) t->next->prev = t->prev;
    else         tail = t->prev;
    delete t;
    return true;
}

bool DoublyList::find(int val) const {
    for (NodeL* t = head; t; t = t->next)
        if (t->val == val) return true;
    return false;
}

//      сериализация

string DoublyList::serialize() const {
    ostringstream os;
    bool first = true;
    for (NodeL* t = head; t; t = t->next) {
        if (!first) os << ' ';
        os << t->val;
        first = false;
    }
    return os.str();
}

void DoublyList::deserialize(const string& s) {
    clear();
    istringstream is(s);
    int v;
    while (is >> v) push_tail(v);
}

string DoublyList::to_string() const {
    ostringstream os;
    os << "[ ";
    for (NodeL* t = head; t; t = t->next) os << t->val << ' ';
    os << ']';
    return os.str();
}