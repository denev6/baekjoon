# 토큰나이저

import re

sentence = input()

# 패턴과 일치할 경우 앞뒤 공백 생성
# e.g. 'a< b' -> 'a <  b'
sentence = re.sub(r"[<>\(\)]|&&|\|\|", r" \g<0> ", sentence)

words = sentence.split()
res = " ".join(words)

print(res)
