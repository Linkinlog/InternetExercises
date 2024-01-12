#include "armstrong_numbers.h"
#include <math.h>
#include <stdlib.h>

bool is_armstrong_number(int candidate) {
  if (candidate == 0) {
    return true;
  }
  int sum = 0;
  int length = floor(log10(abs(candidate))) + 1;
  int candidateCopy = candidate;
  while (candidateCopy > 0) {
    int digit = candidateCopy % 10;
    sum += pow(digit ,length);
    candidateCopy /= 10;
  }
  return sum == candidate;
}
