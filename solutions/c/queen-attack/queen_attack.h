#ifndef QUEEN_ATTACK_H
#define QUEEN_ATTACK_H

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef enum { CAN_NOT_ATTACK, CAN_ATTACK, INVALID_POSITION } attack_status_t;

typedef struct {
   uint8_t row;
   uint8_t column;
} position_t;

attack_status_t can_attack(position_t queen_1, position_t queen_2);
bool can_attack_diagonal(int x1, int y1, int x2, int y2);

#endif
