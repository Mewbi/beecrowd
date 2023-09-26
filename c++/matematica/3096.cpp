#include <iostream>
#include <cmath>
using namespace std;

int kamenetsky(double n) {
    if (n < 0) {
        return 0;
    }

    if (n <= 1) {
        return 1;
    }

    double x = ((n * log10(n/M_E)) +
                (log10(2*M_PI*n) / 2.0));

    return int(floor(x) + 1);
}

int main() {
    double n;
    cin >> n;
    cout << kamenetsky(n) << endl;
    return 0;
}
