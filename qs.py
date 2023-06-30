def quick_sort(arr, low, high):
    if low < high:
        # Partition the array and get the pivot index
        pivot_index = partition(arr, low, high)

        # Recursively sort the subarrays
        quick_sort(arr, low, pivot_index - 1)
        quick_sort(arr, pivot_index + 1, high)




def partition( arr, low, high):
    p = arr[high]
    i = low -1

    for j in range(low,high):
        if arr[j] <= p:
            i = i + 1
            arr[i],arr[j] = arr[j], arr[i]
    arr[i+1],arr[high] = arr[high],arr[i+1]

    return i + 1





'''
def partition(arr, low, high):
    p = arr[high] 
    i = low-1

    for j in range(low,high):
        if arr[j] <= p:
            i = i + 1
            arr[j],arr[i]=arr[i],arr[j]
   
            
    arr[i+1],arr[high]=arr[high],array[i+1]
    return i+1
 '''

 
 
    
# Example usage:
array = [7, 2, 1, 6, 8, 5, 3, 4]
quick_sort(array, 0, len(array) - 1)
print(array)
