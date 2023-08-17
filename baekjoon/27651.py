# 벌레컷

import sys

num_parts = int(sys.stdin.readline())
parts = sys.stdin.readline()


def get_index_of_close_lower_value(target, arr):
    # Binary Search: O(longN)
    start = 0
    end = len(arr) - 1

    while start <= end:
        mid = (start + end) // 2

        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            start = mid + 1
        else:
            end = mid - 1

    return end


parts = [int(x) for x in parts.split()]
partial_sum = []
res = 0

temp = 0
for i in parts:
    temp += i
    partial_sum.append(temp)


"""Perfix Sum: O(N)

가슴(thorax) > 배(abdomen) > 머리(head)

   sum[j] - sum[i] > sum[-1] - sum[j] 
&& sum[-1] - sum[j] > sum[i]

   sum[j] > (sum[i] + sum[-1]) / 2 
&& sum[-1] - sum[i] > sum[j]

Thus, 
  sum[-1] - sum[i] > sum[j] > (sum[i] + sum[-1]) / 2
"""

for i in range(num_parts):
    upper_bound = partial_sum[-1] - partial_sum[i]
    lower_bound = (partial_sum[i] + partial_sum[-1]) / 2

    if upper_bound <= lower_bound:
        break

    upper_index = get_index_of_close_lower_value(upper_bound, partial_sum)
    lower_index = get_index_of_close_lower_value(lower_bound, partial_sum)
    num_cases = upper_index - lower_index
    if partial_sum[upper_index] == upper_bound:
        num_cases -= 1

    res += num_cases

sys.stdout.write(str(res))
