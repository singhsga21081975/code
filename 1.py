def find_pair_with_sum(nums, target):
    num_set = set()
    pairs = []
    for num in nums:
        c = target - num
        if c in num_set:
            pairs.append((c,num))
        num_set.add(num)
    return pairs




# Example usage:
numbers = [1, 2, 3, 4, 5, 6]
target_sum = 9

result = find_pair_with_sum(numbers, target_sum)

if result:
    print(f"Pair with sum {target_sum}: {result}")
else:
    print(f"No pair found with sum {target_sum}")
