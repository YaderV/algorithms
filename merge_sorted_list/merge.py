# Set a new list
# Set two pointers p_left, p_right
# If both are equal, copy both values and update both pointers + 1
# If the left number is less than right; so we copy the value and move the left pointer + 1
# If the right number less than less; so we copy the value and move the right pointer + 1
# when the l_pointer reach m or r_pointer reach n, stop and return list
# if p_left == m and p_right < n; copy right value and update pointer
# if p_right == m and p_right < n; copy right value and update pointer

def merge_lists(nums1, m, nums2, n):
    merged_nums = []
    p_left = 0
    p_right = 0

    while p_left < m or p_right < n:
        # If we reach the end of nums1, we copy all the nums2 values
        if p_left == m:
            merged_nums.append(nums2[p_right])
            p_right += 1

        # If we reach the end of nums2, we copy all the nums2 values
        elif p_right == n:
            merged_nums.append(nums1[p_left])
            p_left += 1

        elif nums1[p_left] == nums2[p_right]:
            merged_nums.append(nums1[p_left])
            merged_nums.append(nums2[p_right])
            p_left += 1
            p_right += 1

        elif nums1[p_left] < nums2[p_right]:
            merged_nums.append(nums1[p_left])
            p_left += 1

        elif nums2[p_right] < nums1[p_left]:
            merged_nums.append(nums2[p_left])
            p_right += 1

    return merged_nums

# Arrays de prueba 1
nums1 = [1,2,3,0,0,0]
m = 3
nums2 = [2,5,6]
n = 3
response = merge_lists(nums1, m, nums2, n)
print(response)

# Arrays de prueba 2
nums1 = [1,2,3,0,0,0,0]
m = 3
nums2 = [-4,2,3,9]
n = 4
response = merge_lists(nums1, m, nums2, n)
print(response)
