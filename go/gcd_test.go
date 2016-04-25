package main

import (
    "testing"
)

func TestGCD(t *testing.T) {
    a, b := 54, 36

    if value := gcd(a, b); value != 18 {
        t.Error("gcd")
    }
}
