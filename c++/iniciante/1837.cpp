#include <iostream>
#include <tuple>

std::tuple<int, int> euclideanDivision(int a, int b) {
    int q, r, limit;
    limit = b < 0 ? -b : b;
    for (r = 0; r < limit; r++) {
        if ((a - r) % b == 0) {
            q = (a - r) / b;
            return std::make_tuple(q, r);
        }
    }
    return std::make_tuple(0, 0);
}

int main() {
    int a, b, q, r;
    std::cin >> a >> b;

    std::tie(q, r) = euclideanDivision(a, b);

    std::cout << q << " " << r << std::endl;
    return 0;
}
