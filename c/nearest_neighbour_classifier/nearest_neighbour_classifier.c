// Nearest Neighbour Classifier
// Requires the data samples from CIFAR-10 website at
// http://www.cs.toronto.edu/~kriz/cifar.html
// Download the C binary data and put the data inside into this folder
//
// This program correctly determines each picture's label at ~39% accuracy

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#include <errno.h>

#define NUM_BATCHES (5)
#define LABEL_SIZE (1)
#define PIXELS_SIZE (3072) // 32 * 32 * 3
#define SAMPLE_SIZE (LABEL_SIZE + PIXELS_SIZE)
#define SAMPLES_PER_BATCH (10000)
#define DATA_LEN (SAMPLES_PER_BATCH * NUM_BATCHES)

static const char *LABELS[] = {
    "airplane",
    "automobile",
    "bird",
    "cat",
    "deer",
    "dog",
    "frog",
    "horse",
    "ship",
    "truck",
};

// Keep the loaded sample data
static char train_data[NUM_BATCHES * DATA_LEN],
            test_data[DATA_LEN];

// Keep track of bytes read from each samples file
static size_t train_byte_offset,
              test_byte_offset;

// The classifier just keeps the training samples data and number of samples
struct nn {
    char *data;
    int count;
};

// Train the classifier (actually, it just sets training samples)
void nn_train(struct nn *classifier, char *data, int count) {
    classifier->data = data;
    classifier->count = count;
}

// Predicts an image's label
int nn_predict(struct nn *classifier, char *test_pixels) {
    int min_sum = INT_MAX;
    int min_index = -1;

    // Loop through each training samples we have
    for (size_t i = 0; i < classifier->count; i++) {
        char *pixels = &classifier->data[i * SAMPLE_SIZE + 1];
        int sum = 0;

        // Calculate the absolute pixel value differences and sum them up
        for (size_t j = 0; j < PIXELS_SIZE; j++) {
            sum += abs(pixels[j] - test_pixels[j]);
        }

        // Look for the sample with the least difference
        if (sum < min_sum) {
            min_sum = sum;
            min_index = i;
        }
    }

    // Return the determined label
    return classifier->data[min_index * SAMPLE_SIZE];
}

int read_data(char *filename, char *data, size_t *index) {
    FILE *file = NULL;
    int ret = 0;

    file = fopen(filename, "rb+");
    if (!file) {
        goto out;
    }

    // Load data
    for (;;) {
        size_t nread = 0;

        // Read label
        nread = fread(&data[*index], 1, 1, file);
        if (nread < 1) {
            if (!feof(file)) {
                ret = -1;
            }
            goto out;
        }

        // Read pixels
        nread = fread(&data[*index + 1], 1, PIXELS_SIZE, file);
        if (nread < PIXELS_SIZE) {
            if (!feof(file)) {
                ret = -1;
            }
            goto out;
        }

        // Advance byte index
        *index += SAMPLE_SIZE;
    }

out:
    if (file) {
        fclose(file);
    }
    return ret;
}

int main() {
    int ret = 0;

    // Read training data
    for (size_t i = 1; i <= NUM_BATCHES; i++) {
        char filename[256];

        ret = snprintf(filename, 256, "data_batch_%zd.bin", i);
        if (ret < 0) {
            goto out;
        }

        ret = read_data(filename, train_data, &train_byte_offset);
        if (ret < 0) {
            goto out;
        }
    }

    // Read test data
    ret = read_data("test_batch.bin", test_data, &test_byte_offset);
    if (ret < 0) {
        goto out;
    }

    // Predict label
    const size_t num_test = 10000;
    struct nn classifier;
    size_t ncorrect = 0;

    // "Train" the classifier
    nn_train(&classifier, train_data, train_byte_offset / SAMPLE_SIZE);

    // Loop through all the test samples and determine its label
    for (size_t i = 0; i < num_test; i++) {

        // Predict the test's label
        const int label = nn_predict(&classifier, &test_data[i * SAMPLE_SIZE + 1]);

        // See if it matches the actual label
        if (label == test_data[i * SAMPLE_SIZE]) {
            ncorrect++;
            fprintf(stdout, "%zd: %zd/%zd\n", i, ncorrect, num_test);
        }
    }

out:
    if (errno) {
        fprintf(stderr, "%s\n", strerror(errno));
    }

    return ret;
}
