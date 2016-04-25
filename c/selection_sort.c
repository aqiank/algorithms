#include <stdio.h>

static void swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

void selection_sort(int *items, int length) {
    if (length < 2) {
        return;
    }

    for (int j = length - 1; j >= 0; j--) {
        int maxPos = j;

        for (int i = 0; i < j; i++) {
            if (items[i] > items[maxPos]) {
                maxPos = i;
            }
        }

        if (maxPos != j) {
            swap(&items[maxPos], &items[j]);
        }
    }
}

void main() {
    int items[] = { 10, 9, 3, 3, 4, 4, 3, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1 };

    selection_sort(items, 17);

    for (int i = 0; i < 17; i++) {
        printf("%d ", items[i]);
    }
    printf("\n");
}
