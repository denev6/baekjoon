import sys


def get_primes(max_range):
    """에라토스테네스의 체"""
    is_prime = [True for _ in range(max_range + 1)]
    is_prime[0] = False

    sqrt = int(max_range**0.5) + 1

    for i in range(2, sqrt + 1):
        if is_prime[i-1]:
            j = 2
            while i * j <= max_range:
                is_prime[i * j - 1] = False
                j += 1

    return is_prime


input = sys.stdin.readlines
print = sys.stdout.write

num_case, *cases = input()
res = [0 for _ in range(int(num_case))]
cases = list(map(int, cases))

max_range = max(cases)
is_prime = get_primes(max_range)

for idx, n in enumerate(cases):
    cnt = 0
    for i in range(1, n // 2 + 1):
        if is_prime[i - 1] and is_prime[n - i - 1]:
            cnt += 1
    res[idx] = cnt

print("\n".join(map(str, res)))
