#include "darts.h"
#include <math.h>

uint8_t score(coordinate_t pos)
{
    float distance;
    
    distance = sqrt((pos.x * pos.x) + (pos.y * pos.y));
    if (distance <= 1.0f)
        return (10);
    if (distance <= 5.0f)
        return (5);
    if (distance <= 10.0f)
        return (1);
    return (0);
}