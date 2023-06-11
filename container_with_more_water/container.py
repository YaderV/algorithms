# - Pointer left and a pointer right; pl = 0, pr = 1
# - We move the right pointer, until we reach length - 1
# - We move left pointer + 1 and right pointer (pl + 1)
# - Calc would be (pr - pl) x min height
# - If the calc of the water is greater than the previous want; save it as max
# - When left_pointer reach the lenght - 1 stop and return max value

def max_area(heights):
    p_left = 0
    p_right = p_left + 1
    max_area = 0

    while p_left < len(heights) - 1:
        area = (p_right - p_left) * min(heights[p_left], heights[p_right])
        if area > max_area:
            max_area = area

        if p_right < len(heights) - 1:
            p_right += 1
        else:
            p_left += 1
            p_right = p_left + 1

    return max_area

if __name__ == "__main__":
    heights = [8, 1, 6, 2, 5, 4, 1, 3, 7]
    print(max_area(heights))
