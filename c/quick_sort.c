#include <stdio.h>

static void swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

static void quick_sort(int *items, int start, int end) {
    if (end - start < 2) {
        return;
    }

    int medianIndex = start;

    for (int i = start + 1; i < end; i++) {
        if (items[i] <= items[medianIndex]) {
            if (i > medianIndex + 1) {
                swap(&items[medianIndex], &items[medianIndex + 1]);
            }

            swap(&items[i], &items[medianIndex]);
            medianIndex++;
        }
    }

    quick_sort(items, start, medianIndex);
    quick_sort(items, medianIndex + 1, end);
}

void main() {
    int items[] = { 10, 9, 3, 3, 4, 4, 3, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1 };

    quick_sort(items, 0, 17);

    for (int i = 0; i < 17; i++) {
        printf("%d ", items[i]);
    }
    printf("\n");
}
