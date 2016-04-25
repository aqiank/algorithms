#include <stdio.h>
#include <stdlib.h>

typedef struct {
    void *left, *right;
    int value;
} bst;

static bst * bst_new(bst *left, bst *right, int value) {
    bst *tree = malloc(sizeof(bst));
    tree->left = left;
    tree->right = right;
    tree->value = value;
}

static bst *bst_insert(bst *tree, int value) {
    if (!tree) {
        return NULL;
    }

    void **next = NULL;

    if (tree->value > value) {
        next = &tree->right;
    } else {
        next = &tree->left;
    }

    if (*next) {
        bst_insert(*next, value);
    } else {
        *next = bst_new(NULL, NULL, value);
    }

    return *next;
}

static bst *bst_find(bst *tree, int value) {
    if (!tree) {
        return NULL;
    }

    if (tree->value == value) {
        return tree;
    } else if (tree->value > value) {
        return bst_find(tree->right, value);
    } else {
        return bst_find(tree->left, value);
    }
}

void main() {
    int numbers[] = { 5, 1, 9, 2, 5, 3, 7, 8, 10, 4 };
    bst *tree = bst_new(NULL, NULL, numbers[0]);

    // Insert numbers
    for (int i = 1; i < 10; i++) {
        bst_insert(tree, numbers[i]);
    }

    // Find number
    tree = bst_find(tree, 8);
    printf("%d\n", tree->value);
}
