import std;
using namespace std;
int main() {
    bool x = true;
    bool y = false;
    if (x and y) {
        std::println("Both true");
    }
    if (x or y) {
        std::println("At least one true");
    }
    if (not x) {
        std::println("x is false");
    }
    int a = 0b1010;
    int b = 0b1100;
    int c = a bitand b;  
    int d = a bitor b;   
    int e = compl a;     
    
    return 0;
}