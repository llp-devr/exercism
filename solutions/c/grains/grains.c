#include "grains.h"

uint64_t square(uint8_t index)
{
    uint8_t i;
    uint64_t res;

    if (index == 0 || index > 64)
        return (0);
    i = 1;
    res = 1;
    while (i < index)
    {
        res *= 2;
        i++;
    }
    return res;
}

uint64_t total(void)
{
    uint8_t i;
    uint64_t res;

    res = 0;
    i = 1;
    while (i <= 64)
    {
        res += square(i);
        i++;
    }
    return (res);
}