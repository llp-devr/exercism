#include "collatz_conjecture.h"

int steps(int start)
{
    int count;

    if (start < 1)
        return (-1);
    count = 0;
    while (start != 1)
    {
        if ((start & 1) == 0)
            start = start / 2;
        else
        {
            start = ((3 * start) + 1) / 2 ;
            count++;
        }
        count++;
    }
    return (count);
}