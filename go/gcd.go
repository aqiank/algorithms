package main

func gcd(a, b int) int {
    for b != 0 {
        remainder := a % b
        a = b
        b = remainder
    }

    return a
}
