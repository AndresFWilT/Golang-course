# Encapsulation

We can make encapsulation with fields name as exportable fields or non-exportable fields instead of access identifiers. 

Sample:

```Go
package main

type Courses struct {
	Name            string //exportable
	isValidClass    bool // non-exportable
}
```


