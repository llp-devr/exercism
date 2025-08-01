#include "dnd_character.h"
#include <stdlib.h>

int ability(void)
{
    int min = 10;
    int total = 0;

    for (int i = 0; i < 4; i++)
    {
        int dice = (rand() % 6) + 1;
        if (i < 3)
            total += dice;
        else
        {
            if (dice > min)
                total += (dice - min);
        }
        if (dice < min)
            min = dice;
    }

    return total;
}

int modifier(int score)
{
    int val = (score - 10) / 2;
    if (score  < 10 && ((score - 10) % 2) != 0)
        val -= 1;
    return (val);
}

dnd_character_t make_dnd_character(void)
{
    dnd_character_t tmp;

    tmp.strength = ability();
    tmp.dexterity = ability();
    tmp.constitution = ability();
    tmp.intelligence = ability();
    tmp.wisdom = ability();
    tmp.charisma = ability();
    tmp.hitpoints = 10 + modifier(tmp.constitution);
    
    return (tmp);
}