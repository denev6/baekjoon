# 시간초과

import sys
from collections import deque

input = sys.stdin.readline
print = sys.stdout.write 

def solution():
    num_case = int(input().rstrip("\n"))
    
    for _ in range(num_case):
        res = ""
        commands = input().rstrip("\n")
        len_queue = int(input().rstrip("\n"))
        
        if len_queue == 0:
          queue = deque()
        else:
          str_queue = input()
          str_queue = str_queue[1:len(str_queue)-2]
          queue = deque(str_queue.split(","))

        for cmd in commands:
            if cmd == "R":
              if len(queue) > 1:
                queue.reverse()
            elif cmd == "D":
                if len(queue) == 0:
                    res = "error"
                    break
                queue.popleft()

        if len(res) == 0:
            res = f"[{','.join(queue)}]"
        print(f"{res}\n")

solution()