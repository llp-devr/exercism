#include "pangram.h"
#include <ctype.h>
#include <string.h>

bool is_pangram(const char *sentence)
{
    if (!sentence)
        return (false);
    
    bool c[26] = { false };
    int len = strlen(sentence);
    
    for (int i = 0; i < len; i++)
    {
        int tmp = 100;
        
        if (65 <= sentence[i] && sentence[i] <= 90)
            tmp = sentence[i] - 65;
        else if (97 <= sentence[i] && sentence[i] <= 122)
            tmp = sentence[i] - 97;
        if (tmp >= 0 && tmp < 26)
            c[tmp] = true;
    }

    for (int i = 0; i < 26; i++)
    {
        if (c[i] == false)
            return (false);
    }
    return (true);
}
