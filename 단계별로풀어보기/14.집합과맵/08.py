# 백준 결과: '틀렸습니다'
# 반례 못 찾음

import sys

input = sys.stdin.readline
print = sys.stdout.write

string = input().rstrip("\n")
len_string = len(string)
res = 0

for len_sub_string in range(1, len_string + 1):
    sub_strings = set(
        string[p : p + len_sub_string] 
        for p in range(len_string - len_sub_string + 1)
    )
    res += len(sub_strings)

print(str(res))
