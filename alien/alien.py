# words = ["habito","hacer","lectura","sonreir"]
# order = "hlabcdfgijkmnopqrstuvwxyz"
# Given a alphabet, calculate the order. dictiorary, map, list index
# Given a list of words, compare the 1rt with the 2nd, 2nd with 3
# If the result if false at any point, stop the program and return false
# If the result is true, continue with the next sequence
# Given two words, set a pointer at the begining of each word
# Compare the p1 with p2, if they are equal continue
# If they are sorted, continue
# If the are not, return false
# If a pointer reach word length, return true

words = ["hello","leetcode"]
order = "hlabcdefgijkmnopqrstuvwxyz"

def compare_words(word_left, word_right, order):
    pointer = 0
    while pointer < len(word_left):
        if len(word_left) > len(word_right):
            return False

        if order[word_left[pointer]] == order[word_right[pointer]]:
            pointer += 1
            continue

        elif order[word_left[pointer]] < order[word_right[pointer]]:
            return True

        return False
    return True


def check_word_list(alphabet, words):
    order = {}
    index = 1
    for letter in alphabet:
        order[letter] = index
        index += 1

    for i in range(0, len(words) - 1):
        if compare_words(words[i], words[i + 1], order):
            continue
        return False

    return True


if __name__ == "__main__":
    result = check_word_list(order, words)
    print(result)



