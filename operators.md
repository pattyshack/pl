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
