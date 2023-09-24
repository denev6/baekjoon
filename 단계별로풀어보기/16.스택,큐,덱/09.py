import sys
from collections import deque

input = sys.stdin.readline
print = sys.stdout.write

deque_ = deque()
num_command = int(input())
res = list()

for _ in range(num_command):
    command, *n = input().strip().split()
    if command == "1":
        deque_.appendleft(n[0])  
    elif command == "2":
        deque_.append(n[0])  
    elif command == "3":
        res.append(deque_.popleft() if deque_ else -1)
    elif command == "4":
        res.append(deque_.pop() if deque_ else -1)
    elif command == "5":
        res.append(len(deque_))
    elif command == "6":
        res.append(0 if deque_ else 1)
    elif command == "7":
        res.append(deque_[0] if deque_ else -1)
    elif command == "8":
        res.append(deque_[-1] if deque_ else -1)

print("\n".join(map(str, res)))
