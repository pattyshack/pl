## Operators

### Logical operators
```
and    conditional and   p && q  is "if p then q else false"
or     conditional or    p || q  is "if p then true else q"
not    not               !p      is "not p"
```

- expressions are lazily evaluated from left to right
- does not support operator overloading due to lazy eval
- compiler should have a built-in bool primitive type for control flow purpose

XXX: should we allow other types to auto convert to / behave like bool? probably not.

### Comparison operations
```
==  equal
!=  not equal 
<   less
<=  less or equal
>   greater
>=  greater or equal
```

implementing Comparable interface enables == and != operator usage.
- the Comparable interface only implments equal, != is implicit defined.
- Comparable operation must be commutative, `a == b` eqv. `b == a`

implementing Ordered interface enables, <, <=, >, >= operator usage
- the Ordered interface only implments less and lessOrEqual, > and >= are implicit defined.
- < and <= must be associative, `a < b < c` eqv. `(a < b) < c` eqv. `a < (b < c)`
- <= must be reflexive, `x <= x` is true

default behavior:
- since basic types such as int, float, etc. are not built into the language, basic types are comparable/ordered if the basic type implement Comparable/Ordered interface.
- interface of the same type are comparable.  return true if two interfaces refers to the same object.
- structs of the same type are comparable/ordered if all fields are comparable/ordered.  Fields are compared in source order.
- arrays are comparable if array element type is comparable/ordered.  The array is lexicographically ordered
- enum are comparable/ordered if enum data types are comparable / ordered.  enum type are first ordered by enum source order, follow by enum's data order

custom behavior:
- a type that implement Comparable / Ordered overrides default behavior.  non-homogeneous type comparison is allowed by parametric polymorphism as long as the operator is uniquely implemented by one of the type (not by both types).  the compiler may reorder terms while perserving expression eval order.  e.g., suppose type2 implements Ordered `type2.Less(type1)`, the compiler may rewrite `type1_expr > type2_expr` as `{t1_val := type1_expr; t2_val := type2_expr; t2_val < v1_val}`

allow python style lazy eval chained comparison, e.g.,
- `1 < x < y <= 10` eqv to `{x_val := eval x expression;  1 < x_val && { y_val := eval y expression; x_val < y_val && y_val <= 10}}`
- 32 > char < 127  // non-printable ascii character

\<operation\>= - operation assignment always yield unit

+=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>= ++, --, etc are statements with unit as return value


### Arithmetic operations
```
+     sum
-     different
*     product
\     quotient
%     remainder
~     bitwise negation
&     bitwise and
|     bitwise or
^     bitwise xor
<<    bitwise left shift
>>    bitwsie right shift
```
a type implementing a operator interface enable that operator for the type.  TODO: need to figure out the op interface, in particular, how to specify arithmetic relationships such as commutativity `a + b == b + a`  associativity `(a + b) + c == a + (b + c)` (distributive property, `a * (b + c) == a * b + a * c`, is harder to deal with since it involves multiple ops, but probably is not needed)

`<arithmetic operations>=` op assignment statements, including `++` / `--`. , always yield unit.  not usable directly inside another expression (but can be use in a nest expr block)
