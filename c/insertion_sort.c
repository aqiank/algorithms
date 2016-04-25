#include <stdio.h>

static void insertion_sort(int *items, int length) {
    if (length < 2) {
        return;
    }

    for (int j = 1; j < length; j++) {
        int value = items[j];
        int i = j - 1;
        while (i >= 0 && items[i] > value) {
            items[i + 1] = items[i];
            i--;
        }
        items[i + 1] = value;
    }
}

void main() {
    int items[] = { 10, 9, 3, 3, 4, 4, 3, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1 };

    insertion_sort(items, 17);

    for (int i = 0; i < 17; i++) {
        printf("%d ", items[i]);
    }
    printf("\n");
}
