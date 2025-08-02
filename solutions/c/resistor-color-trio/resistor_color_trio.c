#include "resistor_color_trio.h"

#include <math.h>

resistor_value_t color_code(resistor_band_t *arr){
    int acc = 0;
    resistor_unit_t unit = 0;
    resistor_value_t res;

    acc = (arr[0] * 10) + arr[1];
    unit += arr[2] / 3;
    acc *= pow(10, (arr[2] % 3));
    while (acc > 1000)
    {
        acc /= 1000;
        unit++;
    }
    res.value = acc;
    res.unit = unit;
    
    return (res);
}