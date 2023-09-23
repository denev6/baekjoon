import sys
from collections import deque

input = sys.stdin.readline
print = sys.stdout.write

stack = deque()
num_command = int(input())
res = list()

for _ in range(num_command):
    command = input().strip()
    if command[0] == "1":
        n = command[2:]
        stack.append(n)  
    elif command == "2":
        res.append(stack.pop() if stack else -1)
    elif command == "3":
        res.append(len(stack))
    elif command == "4":
        res.append(0 if stack else 1)
    elif command == "5":
        res.append(stack[-1] if stack else -1)

print("\n".join(map(str, res)))
