#include <math.h>

#include "eliuds_eggs.h"

int egg_count(int eggs)
{
    int acc = 0;

    for (int i = 32; i >= 0; i--)
    {
        if (eggs < pow(2, i))
            acc += 0;
        else
        {
            eggs -= pow(2, i);
            acc += 1;
        }
    }

    return (acc);
}