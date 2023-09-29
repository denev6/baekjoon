import sys
from collections import deque

input = sys.stdin.readline
print = sys.stdout.write 

input()  # Skip
structures = input().rstrip("\n").split(" ")
items = input().rstrip("\n").split(" ")
input()  # Skip
new_items = input().rstrip("\n").split(" ")

qs = deque()

for i, flag in enumerate(structures):
    if flag == "0":
        qs.append(items[i])

res = [None for i in range(len(new_items))]

for i, new in enumerate(new_items):
    qs.appendleft(new)
    res[i] = qs.pop() 
print(" ".join(res))
