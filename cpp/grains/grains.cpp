#include "grains.h"

namespace grains {
    unsigned long long square(int squarePosition) {
        return static_cast<unsigned long long>(std::pow(2, squarePosition-1));
    }

    unsigned long long total() {
        unsigned long long grains {};

        for (int i {1}; i <= 64; i++) {
            grains += square(i);
        }
        return grains;
    }
}  // namespace grains
