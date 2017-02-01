
## Golang solution for "The Change Problem"

```
What is the minimum number of coins to give back as change ?
Input: the change N and a set of coins cut T.
Output: a list of coins cut that summed equals to N. That list must contains the
        minimum number of elements.
```

For example:
let the coins cut be `T = {1, 2, 5, 10, 20, 50, 100, 200}` (in cents of Euro)
and you want to give a change of `N = 27`.

Well, the minimum number of coins are 3:
2
5
20

## Try yourself
There are two file that solve the same problem, one use dynamic programming and
the other a greedy choice.

Change T and N as you prefer in one of the file and run:
```bash
go run change_xxxx.go
```
