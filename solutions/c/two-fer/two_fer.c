#include <stdio.h>
#include <string.h>

#include "two_fer.h"

void two_fer(char *buffer, const char *name)
{
    if (!name || !*name)
        strcpy(buffer, "One for you, one for me.");
    else
        sprintf(buffer, "One for %s, one for me.", name);
}