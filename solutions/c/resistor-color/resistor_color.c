#include "resistor_color.h"

int color_code(resistor_band_t color)
{
    return ((int) color);
}

resistor_band_t *colors(void)
{
    static resistor_band_t color_list[] = {
        BLACK, BROWN, RED, ORANGE, YELLOW,
        GREEN, BLUE, VIOLET, GREY, WHITE
    };
    return (color_list);
}