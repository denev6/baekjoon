# 증가 배열 만들기

user_input = input()
n, m, k = [int(i) for i in user_input.split()]

if k < n + m - 1:
    print("NO")
else:
    print("YES")
    for i in range(n):
        for j in range(1, m + 1):
            print(i + j, end=" ")
        print()
