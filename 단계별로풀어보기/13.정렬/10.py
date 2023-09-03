import sys

input = sys.stdin.readline
print = sys.stdout.write 

num_member = int(input())
members = [0 for _ in range(num_member)]

for id_ in range(num_member):
    age, name = input().split()
    members[id_] = [id_, int(age), name]

sorted_members = sorted(members, key=lambda x: (x[1], x[0]))

for _, age, name in sorted_members:
    print(f"{age} {name}\n")
