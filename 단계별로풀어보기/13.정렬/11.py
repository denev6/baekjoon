import sys

input = sys.stdin.readline
print = sys.stdout.write 

_ = input()
ordered_numbers = list(map(int, input().split()))
sorted_numbers = sorted(set(ordered_numbers))
res = dict()

for i, n in enumerate(sorted_numbers):
  res[n] = i

print(" ".join(str(res[n]) for n in ordered_numbers))
