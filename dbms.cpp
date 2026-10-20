#include "dbms.hpp"
#include <iostream>
#include <fstream>
#include <sstream>
using namespace std;

DBMS::DBMS(const string& fname) : filename(fname) {
    load();
}

DBMS::~DBMS() {
    save();
}

//      Файл
// Формат строки: <ТЕГ> <имя> <данные>
void DBMS::load() {
    ifstream in(filename);
    if (!in.is_open()) return;

    string line;
    while (getline(in, line)) {
        if (line.empty()) continue;
        istringstream is(line);
        string tag, name;
        is >> tag >> name;

        string rest;
        getline(is, rest);
        size_t p = 0;
        while (p < rest.size() && rest[p] == ' ') ++p;
        rest = rest.substr(p);

        if (tag == "M") arrays[name].deserialize(rest);
        else if (tag == "F") flists[name].deserialize(rest);
        else if (tag == "L") dlists[name].deserialize(rest);
        else if (tag == "S") stacks[name].deserialize(rest);
        else if (tag == "Q") queues[name].deserialize(rest);
        else if (tag == "T") trees[name].deserialize(rest);
    }
}

void DBMS::save() const {
    ofstream out(filename);
    if (!out.is_open()) return;

    for (const auto& p : arrays)  out << "M " << p.first << " " << p.second.serialize() << "\n";
    for (const auto& p : flists)  out << "F " << p.first << " " << p.second.serialize() << "\n";
    for (const auto& p : dlists)  out << "L " << p.first << " " << p.second.serialize() << "\n";
    for (const auto& p : stacks)  out << "S " << p.first << " " << p.second.serialize() << "\n";
    for (const auto& p : queues)  out << "Q " << p.first << " " << p.second.serialize() << "\n";
    for (const auto& p : trees)   out << "T " << p.first << " " << p.second.serialize() << "\n";
}

//      PRINT
void DBMS::printAll() const {
    bool any = false;

    for (const auto& p : arrays) {
        cout << "ARRAY    " << p.first << " : " << p.second.serialize() << "\n";
        any = true;
    }
    for (const auto& p : flists) {
        cout << "FLIST    " << p.first << " : " << p.second.to_string() << "\n";
        any = true;
    }
    for (const auto& p : dlists) {
        cout << "DLIST    " << p.first << " : " << p.second.to_string() << "\n";
        any = true;
    }
    for (const auto& p : stacks) {
        cout << "STACK    " << p.first << " : " << p.second.to_string() << "\n";
        any = true;
    }
    for (const auto& p : queues) {
        cout << "QUEUE    " << p.first << " : " << p.second.to_string() << "\n";
        any = true;
    }
    for (const auto& p : trees) {
        cout << "TREE     " << p.first << " : " << p.second.to_string() << "\n";
        any = true;
    }

    if (!any) cout << "(empty)\n";
}

//      Разбор команд
void DBMS::execute(const string& query) {
    istringstream is(query);
    string cmd;
    is >> cmd;

    if (cmd == "PRINT") {
        printAll();
        return;
    }

    string name;
    is >> name;

    if (name.empty()) {
        cout << "Error: no structure name\n";
        return;
    }

    //      M: массив
    if (cmd == "MPUSH") { int v; is >> v; arrays[name].push(v); }
    else if (cmd == "MDEL") { int i; is >> i; if (!arrays[name].del(i)) cout << "Error: bad index\n"; }
    else if (cmd == "MGET") { int i; is >> i; cout << arrays[name].get(i) << "\n"; }
    else if (cmd == "MSET") { int i, v; is >> i >> v; arrays[name].set(i, v); }
    else if (cmd == "MLEN") { cout << arrays[name].length() << "\n"; }
    else if (cmd == "MPRINT") { cout << arrays[name].serialize() << "\n"; }

    //      F: односвязный список
    else if (cmd == "FADDH") { int v; is >> v; flists[name].push_head(v); }
    else if (cmd == "FADDT" || cmd == "FPUSH") { int v; is >> v; flists[name].push_tail(v); }
    else if (cmd == "FADDB") { int t, v; is >> t >> v; if (!flists[name].push_before(t, v)) cout << "Error: target not found\n"; }
    else if (cmd == "FADDA") { int t, v; is >> t >> v; if (!flists[name].push_after (t, v)) cout << "Error: target not found\n"; }
    else if (cmd == "FDELH") { if (!flists[name].pop_head()) cout << "Error: empty\n"; }
    else if (cmd == "FDELT") { if (!flists[name].pop_tail()) cout << "Error: empty\n"; }
    else if (cmd == "FDEL")  { int v; is >> v; if (!flists[name].del(v)) cout << "Error: not found\n"; }
    else if (cmd == "FDELI") { int i; is >> i; if (!flists[name].del_index(i)) cout << "Error: bad index\n"; }
    else if (cmd == "FGET")  { int v; is >> v; cout << (flists[name].find(v) ? "TRUE" : "FALSE") << "\n"; }
    else if (cmd == "FPRINT"){ cout << flists[name].to_string() << "\n"; }

    //      L: двусвязный список
    else if (cmd == "LADDH") { int v; is >> v; dlists[name].push_head(v); }
    else if (cmd == "LADDT" || cmd == "LPUSH") { int v; is >> v; dlists[name].push_tail(v); }
    else if (cmd == "LADDB") { int t, v; is >> t >> v; if (!dlists[name].push_before(t, v)) cout << "Error: target not found\n"; }
    else if (cmd == "LADDA") { int t, v; is >> t >> v; if (!dlists[name].push_after (t, v)) cout << "Error: target not found\n"; }
    else if (cmd == "LDELH") { if (!dlists[name].pop_head()) cout << "Error: empty\n"; }
    else if (cmd == "LDELT") { if (!dlists[name].pop_tail()) cout << "Error: empty\n"; }
    else if (cmd == "LDEL")  { int v; is >> v; if (!dlists[name].del(v)) cout << "Error: not found\n"; }
    else if (cmd == "LDELI") { int i; is >> i; if (!dlists[name].del_index(i)) cout << "Error: bad index\n"; }
    else if (cmd == "LGET")  { int v; is >> v; cout << (dlists[name].find(v) ? "TRUE" : "FALSE") << "\n"; }
    else if (cmd == "LPRINT"){ cout << dlists[name].to_string() << "\n"; }

    //      S: стек
    else if (cmd == "SPUSH") { int v; is >> v; stacks[name].push(v); }
    else if (cmd == "SPOP") { cout << stacks[name].pop() << "\n"; }
    else if (cmd == "SPRINT") { cout << stacks[name].to_string() << "\n"; }

    //      Q: очередь
    else if (cmd == "QPUSH") { int v; is >> v; queues[name].push(v); }
    else if (cmd == "QPOP") { cout << queues[name].pop() << "\n"; }
    else if (cmd == "QPRINT") { cout << queues[name].to_string() << "\n"; }

    //      T: Full Binary Tree
    else if (cmd == "TINSERT") { int v; is >> v; trees[name].insert(v); }
    else if (cmd == "TGET") { int v; is >> v; cout << (trees[name].find(v) ? "TRUE" : "FALSE") << "\n"; }
    else if (cmd == "TISFULL") { cout << (trees[name].isFull() ? "TRUE" : "FALSE") << "\n"; }
    else if (cmd == "TPRINT") { cout << trees[name].to_string() << "\n"; }

    else {
        cout << "Unknown command: " << cmd << "\n";
    }
}