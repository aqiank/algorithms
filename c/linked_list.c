#include <stdio.h>
#include <stdlib.h>

typedef struct {
    void *prev, *next;
    int value;
} linked_list;

static linked_list * linked_list_new(linked_list *prev, linked_list *next, int value) {
    linked_list *list = malloc(sizeof(linked_list));
    list->prev = prev;
    list->next = next;
    list->value = value;
}

void main() {
    int numbers[] = { 5, 1, 9, 2, 5, 3, 7, 8, 10, 4 };
    linked_list *head = linked_list_new(NULL, NULL, numbers[0]);
    linked_list *tail = NULL;
    linked_list *entry = head;

    for (int i = 1; i < 10; i++) {
        linked_list *next = linked_list_new(entry, NULL, numbers[i]);
        entry->next = next;
        entry = next;
        if (i == 9) {
            tail = next;
        }
    }

    // Follow linked-list forward
    entry = head;
    for (int i = 0; i < 10; i++) {
        printf("%d ", entry->value);
        entry = entry->next;
    }
    printf("\n");

    // Follow linked-list backward
    entry = tail;
    for (int i = 0; i < 10; i++) {
        printf("%d ", entry->value);
        entry = entry->prev;
    }
    printf("\n");
}
