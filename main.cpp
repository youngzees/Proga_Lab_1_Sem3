#include <iostream>         //g++ -Wall -Wextra main.cpp dbms.cpp structures/*.cpp -o dbms
#include <string>           // или так: g++ -Wall -Wextra main.cpp dbms.cpp structures/*.cpp -o dbms
using namespace std;

#include "dbms.hpp"

int main(int argc, char* argv[]) {
    string file = "file.data";
    string query;

    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) {
            file = argv[++i];
        } else if (arg == "--query" && i + 1 < argc) {
            query = argv[++i];
        }
    }

    DBMS db(file);

    if (!query.empty()) {
        // одна команда и выход (сценарий из задания)
        db.execute(query);
    } else {
        // интерактивный режим — удобно для тестов
        cout << "DBMS interactive mode. Type 'exit' to quit.\n";
        string line;
        while (getline(cin, line)) {
            if (line == "exit" || line == "quit") break;
            if (line.empty()) continue;
            db.execute(line);
        }
    }
    return 0;
}