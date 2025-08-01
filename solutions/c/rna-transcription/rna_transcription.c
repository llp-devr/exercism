#include "rna_transcription.h"



char *to_rna(const char *dna)
{
    char *rna;
    uint32_t i;

    rna = (char *) malloc(sizeof(char) * (strlen(dna) + 1));
    if (!rna)
        return (NULL);
    i = 0;
    while (dna[i] != 0)
    {
        if (dna[i] == 'G')
            rna[i] = 'C';
        if (dna[i] == 'C')
            rna[i] = 'G';
        if (dna[i] == 'T')
            rna[i] = 'A';
        if (dna[i] == 'A')
            rna[i] = 'U';
        i++;
    }
    rna[i] = 0;
    return (rna);
}