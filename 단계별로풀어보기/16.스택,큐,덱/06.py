import sys
from collections import deque

input = sys.stdin.readline
print = sys.stdout.write

queue = deque()
num_command = int(input())
res = list()

for _ in range(num_command):
    command, *n = input().strip().split()
    if command == "push":
        queue.append(n[0])  
    elif command == "pop":
        res.append(queue.popleft() if queue else -1)
    elif command == "size":
        res.append(len(queue))
    elif command == "empty":
        res.append(0 if queue else 1)
    elif command == "front":
        res.append(queue[0] if queue else -1)
    elif command == "back":
        res.append(queue[-1] if queue else -1)

print("\n".join(map(str, res)))
