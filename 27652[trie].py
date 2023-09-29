# AB
### 시간 초과 (실패) ###

import sys


class Node(object):
    def __init__(self, key):
        self.key = key
        self.children = {}
        self.count = 0


class Trie(object):
    def __init__(self):
        self._head = Node(None)

    def add(self, string):
        current_node = self._head

        for char in string:
            if char not in current_node.children:
                current_node.children[char] = Node(char)
            current_node = current_node.children[char]
            current_node.count += 1

    def delete(self, string):
        # string이 반드시 존재한다고 가정 (시간 단축)
        current_node = self._head
        for char in string:
            previous_node = current_node
            current_node = current_node.children[char]
            if current_node.count == 1:
                previous_node.children.pop(char)
                return
            current_node.count -= 1

    def count_char(self, target):
        current_node = self._head
        for char in target:
            if char in current_node.children:
                current_node = current_node.children[char]
            else:
                return 0
        return current_node.count


a_trie = Trie()
b_trie = Trie()


def add(trie, string):
    if trie == "A":
        a_trie.add(string)
    elif trie == "B":
        b_trie.add(string[::-1])


def delete(trie, string):
    if trie == "A":
        a_trie.delete(string)
    elif trie == "B":
        b_trie.delete(string[::-1])


def find(target):
    cnt = 0
    for i in range(1, len(target)):
        perfix = target[:i]
        suffix = target[i:][::-1]
        cnt_perfix = a_trie.count_char(perfix)
        if cnt_perfix != 0:
            cnt_suffix = b_trie.count_char(suffix)
            cnt += cnt_perfix * cnt_suffix
    return cnt


def execute_query(query: str):
    command, *args = query.split()
    if command == "add":
        add(args[0], args[1])
    elif command == "delete":
        delete(args[0], args[1])
    elif command == "find":
        cnt = find(args[0])
        sys.stdout.write(f"{cnt}\n")


num_queries = int(sys.stdin.readline())

for _ in range(num_queries):
    query = sys.stdin.readline()
    execute_query(query)
