#include <stdio.h>

static int binary_search(int *items, int length, int value) {
    if (length < 2) {
        return 0;
    }


    int index = length / 2;
    while (value != items[index]) {
        if (value > items[index]) {
            index += (length - index - 1) / 2;
        } else {
            index /= 2;
        }
    }

    return index;
}

void main() {
    int items[] = { 1, 1, 1, 3, 4, 5, 5, 7, 7, 10, 11, 14, 17 };

    int index = binary_search(items, 13, 10);
    printf("%d\n", index);
}
