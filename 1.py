def find_pair_with_sum(nums, target):
    num_set = set()

    for num in nums:
        complement = target - num
        if complement in num_set:
            return (complement, num)
        num_set.add(num)

    return None

# Example usage:
numbers = [1, 2, 3, 4, 5, 6]
target_sum = 9

result = find_pair_with_sum(numbers, target_sum)

if result:
    print(f"Pair with sum {target_sum}: {result[0]} and {result[1]}")
else:
    print(f"No pair found with sum {target_sum}")
