# Shapes

## Start
Every shape diagram must begin with start shape.


## Variable 
It's the place where variables are declared. 
```
x = 0
y = 2
```


## If 
Evaluated logic expression. Accepts '&&' and "||" operators as and, or respectively. Doesn't evaluate 'and'/'or' keywords.

```
t[i] < d || t[5] > 2
```


## Action
There is 4 types of action in this shape:

- Print, prints variable to the debug console on the screen
```
print x
```

- Math operations, performs operations with assignment, variable must be previously declared. 
```
x++
x--
x += 3
x -= x/2
x /= t[i]*2
x = y
```

- Swap, swaps values of two variables 
```
swap t[i], x
```

- Rand, randomize value in range with assignment. Rand is always floating number
```
x = rand 2, 5
```


## Stop
Every shape diagram must end with stop shape.
