import sys

input = sys.stdin.readline
print = sys.stdout.write 

num_point = int(input())
arr = [0 for _ in range(num_point)]

for i in range(num_point):
    arr[i] = list(map(int, input().split()))

sorted_arr = sorted(arr, key=lambda x: (x[0], x[1]))

for x, y in sorted_arr:
    print(f"{x} {y}\n")
