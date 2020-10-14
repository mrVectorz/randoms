row_count = 0
total_numbers = 0
odds = 0

def pascals_triangle(n_rows):
    results = [] # a container to collect the rows
    for _ in range(n_rows):
        row = [1] # a starter 1 in the row
        if results: # then we're in the second row or beyond
            last_row = results[-1] # reference the previous row
            # this is the complicated part, it relies on the fact that zip
            # stops at the shortest iterable, so for the second row, we have
            # nothing in this list comprension, but the third row sums 1 and 1
            # and the fourth row sums in pairs. It's a sliding window.
            row.extend([sum(pair) for pair in zip(last_row, last_row[1:])])
            # finally append the final 1 to the outside
            row.append(1)
        results.append(row) # add the row to the results.
    return results

for i in pascals_triangle(10000):
    row_count += 1
    total_numbers += len(i)
    for x in i:
        if x%2 != 0:
            odds += 1
    print("{}: {}/{} = {}".format(row_count, odds, total_numbers, odds/total_numbers))
