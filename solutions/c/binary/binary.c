#include <math.h>
#include <string.h>

#include "binary.h"

int convert(const char *input)
{   
    int acc = 0;
    int len = strlen(input);
    
    for (int i = 0; i < len; i++)
    {
        int digit = input[i] - '0';
        if (digit != 0 && digit != 1)
            return (-1);
        acc += digit * pow(2, (len - i - 1));
    }
    
    return (acc);
}