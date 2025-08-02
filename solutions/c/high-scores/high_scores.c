#include <stdlib.h>

#include "high_scores.h"

int32_t latest(const int32_t *scores, size_t scores_len)
{
    return (scores[scores_len - 1]);
}

int32_t personal_best(const int32_t *scores, size_t scores_len)
{
    int max = 0;

    for (size_t i = 0; i < scores_len; i++)
    {
        if (scores[i] > max)
            max = scores[i];
    }

    return (max);
}

size_t personal_top_three(const int32_t *scores, size_t scores_len,
                          int32_t *output)
{
    if (scores_len == 0)
        return 0;
    
    int32_t first = INT32_MIN;
    int32_t second = INT32_MIN;
    int32_t third = INT32_MIN;
    
    for (size_t i = 0; i < scores_len; i++) 
    {
        if (scores[i] >= first) {
            third = second;
            second = first;
            first = scores[i];
        } else if (scores[i] >= second && scores[i] != first) {
            third = second;
            second = scores[i];
        } else if (scores[i] >= third && scores[i] != second && scores[i] != first) {
            third = scores[i];
        }
    }
    
    size_t count = 0;
    if (first != INT32_MIN) output[count++] = first;
    if (second != INT32_MIN) output[count++] = second;
    if (third != INT32_MIN) output[count++] = third;
    
    return count;
}