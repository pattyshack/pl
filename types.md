# Types

## Named Tuple

```
()  // implicit tuple type declaration.  Can be used when the context is unambiguous
tuple() // explicit tuple type declaration.  Can be used anywhere.
```

anonymous tuple type definition within a function body must use the explicit form. e.g.,
```
func blah(i int, f float) (int, float) {
  tuple(int, float)(i, f)
}
```

all tuple fields are named, either implicitly or explicitly.  explicitly named fields must be uniquely named, whereas implicitly named field, which takes on the type's name, may be ambiguous as long as the fields are only accessed via pattern matching (error otherwise).

```
(int, float, third string) // implicitly named first field as int, second as float, third as third
```

support go style embedded field promotion

tuple initialization: allow both positional and associative name mapping by default, which assigns the value to the corresponding fields.  positional and associative mapping intialization are operators and can be override (to support vectors / maps).  need option to disallow operators (vector shouldn't support map style initialization, etc)

idea?: tuples are unordered.  tuple A is equivalent / super type of tuple B if all of B's fields / methods B unambiguously maps to fields in fields / methods of A.  a named field of callable type in A is a supertype of a B's method if it has the same name and the method signature is a subtype of the callable's type.
