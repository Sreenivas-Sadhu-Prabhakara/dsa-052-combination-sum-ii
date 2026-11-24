# Combination Sum II

> Striver SDE Sheet 052 · Day 9 Recursion — single-use elements (`i+1`), duplicate values skipped at the same depth, same sorted cutoff. Go, stdlib only, tested.

## Problem

Like Combination Sum, but each array element may be used **once**, and the array itself may contain duplicate values — yet the output must contain no duplicate combinations. `[2,5,2,1,2]` target 5 yields `[1,2,2]` and `[5]`, not three copies of `[1,2,2]`.

## Approach

Two one-line changes to Combination Sum carry all the new semantics. Recurse with `i+1`: each element is consumed. And at each loop level skip values equal to their left neighbor (`i > start`): picking the second identical 2 where the first was available would regenerate the same combination subtree. Vertical repetition (a 2 after taking a 2) stays legal — it uses *different* elements. The sorted-order break on `a[i] > remaining` prunes as before.

**Complexity:** O(2ⁿ) worst case (output-bound), O(n) depth.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
