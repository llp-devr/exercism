#include "hamming.h"

uint16_t ft_strlen(const char *s)
{
    uint16_t i;

    i = 0;
    while (s[i] != 0)
        i++;
    return (i);
}

int compute(const char *lhs, const char *rhs)
{
    int d;
    uint16_t i;

    if (ft_strlen(lhs) != ft_strlen(rhs))
        return (-1);
    d = 0;
    i = 0;
    while (lhs[i] != 0 || rhs[i] != 0)
    {
        if (lhs[i] != rhs[i])
            d++;
        i++;
    }
    return (d);
}