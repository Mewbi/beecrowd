#include <stdio.h>

int main() {
    int n;
    while (scanf("%d", &n) != EOF) {
        const char* out = "vai ter copa!";
        if (n > 0) {
            out = "vai ter duas!";
        }
        printf("%s\n", out);
    }
    return 0;
}

