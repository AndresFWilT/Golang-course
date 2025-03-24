# Maps

maps are key value structures

## Creation

There are two ways to create maps in go:

1. using make

```Go
music := make(map[string]string)
	music["Guitar"] = "Gibson"
	music["Magic"] = "Magician"
```

2. using literals

```Go
tech := map[string]string{
		"Guitar": "Guitar",
		"Magic":  "Magic",
	}
```

## Deletion

To delete objects use

```Go
delete(map, "key")
```

## Get Value

To get if a value exists a map returns:

```Go
var pets = make(map[string]string)
pets["Dog"] = "Danger"
value, ok := map["Dog"]
```