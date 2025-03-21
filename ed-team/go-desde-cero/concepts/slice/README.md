# Slice

 Are pointers from an array, don't have any data. The first index will contain the first element of the array but the last index don't.
 
If the values from an slice that points to an array are updated, also will be the values from the array.

## Capacity of a Slice

Number of elements of the original array, since the index created until the end, no matter if the slice size is n, it will have the capacity from the starting index.
If we keep appending values to the slice that surpass the capacity of the original array Go will create a new array with double size of the original slice capacity at max.

## Size of a Slice

Number of elements from the slice

## Slice Creation

To create slice we can:

```Go
pets := []string{"dog","cat"}
```

is like an array but you don't specify the size. Another way is with make:

```Go
pets := make(string, 2, 5)
```

make receives the type, size and capacity. It's a better approach if the slice is not created from an array.
