#include "perfect_numbers.h"

kind classify_number(int n)
{
    int aliquot = 0;

    if (n < 1)
        return (ERROR);

    for (int i = 1; i < ((n / 2) + 1); i++)
    {
        if (n % i == 0)
            aliquot += i;
    }
    if (aliquot == n)
        return (PERFECT_NUMBER);
    if (aliquot < n)
        return (DEFICIENT_NUMBER);
    if (aliquot > n)
        return (ABUNDANT_NUMBER);
    return (ERROR);
}