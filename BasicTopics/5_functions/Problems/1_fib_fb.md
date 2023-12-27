Correctness: 5/5
The iterative implementation appears correct. It initializes the first two Fibonacci numbers correctly and then iteratively computes the next numbers based on the sum of the two preceding ones.

Complexity: 4/5
The time complexity of this implementation is O(n), where n is the number of terms. It uses an array to store previously computed Fibonacci numbers, which contributes to space complexity of O(n). The space complexity could be improved by using only two variables to store the last two Fibonacci numbers, reducing it to O(1).