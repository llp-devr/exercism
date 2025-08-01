#include "space_age.h"

float age(planet_t planet, int64_t seconds)
{
    if (planet == MERCURY)
        return (seconds / 31557600.0f / 0.2408467f);
    if (planet == VENUS)
        return (seconds / 31557600.0f / 0.61519726f);
    if (planet == EARTH)
        return (seconds / 31557600.0f / 1.0f);
    if (planet == MARS)
        return (seconds / 31557600.0f / 1.8808158f);
    if (planet == JUPITER)
        return (seconds / 31557600.0f / 11.862615f);
    if (planet == SATURN)
        return (seconds / 31557600.0f / 29.447498f);
    if (planet == URANUS)
        return (seconds / 31557600.0f / 84.016846f);
    if (planet == NEPTUNE)
        return (seconds / 31557600.0f / 164.79132f);
    return (-1);
}