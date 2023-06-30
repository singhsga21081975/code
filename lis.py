def lis(nums):
    n = len(nums)
    dp= [ 0] * n-1 
    prev = [ -1] * n-1

    for i in range (n-1):
        for j in  range(i):
            if nums[j] < nums[i] and dp[i] < dp[j] + 1 :
                dp[i] = dp[j-1] + 1
                prev[i] = j

    maxLen = max(dp)

    for i in range(maxLen):
        print( nums[prev[i]])
