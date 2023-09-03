import sys

input = sys.stdin.readline
print = sys.stdout.write 

num_word = int(input())
words = [0 for _ in range(num_word)]

for i in range(num_word):
    words[i] = input().rstrip('\n') 

words = list(set(words))
sorted_words = sorted(words, key=lambda x: (len(x), x))
print("\n".join(sorted_words))
