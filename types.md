# Type ideas

## proloaded "primitive" types

no primitive types built into the parser. hardcode some primitive types post-parsing for now / switch to schema/config based preloading later.  primitive types are just struct.

## struct

array also behave like normal opaque struct must pass with immut / mut ref

syntactic sugar for data only tuples (primarily intended for used within func args list or lookup index)?
```
expr1:expr2:expr3 -> (expr1, expr2, expr3)

: -> ((), ())
<expr>: -> (<expr>, ())
:<expr> -> ((), <expr>)
:: -> ((), (), ())

```


PROBABLY NOT A GOOD IDEA TO BAKE INTERFACE INTO STRUCT.  REDO THIS WHOLE SECTIOn

A struct define an interface signature and optionally concrete implementation.  The interface is generalized to allow fields as well as trait method definitions.

struct may extend other struct signatures.  the extended struct "inherit" the parent signature's fields and method declarations/definitions.

how to deal with "multi-inheritance" conflict?
- field declarations must be identically typed
- allow duplicated method declaration if one signature is a subtype of the another signature (contravariant parameter, covariant return type).  supertype declaration / definition is used
- struct must manually implement method if there are multiple indentically typed parent method definitions
- compile error otherwise

```
// implicit struct type declaration.  Can be used when the context is unambiguous
(<struct declaration>*)

// explicit struct type declaration.  Can be used anywhere.
struct <generic parameters>* [extends <comma seperated signature list>]* (
  // all declarations outside the concrete implementation scope is part of the interface signature
  // helper methods need not be declared

  <field or method declaration / definition>*

  // specific field needed by the concrete implementation.  may optionally declare/define helper methods
  // fields / methods from the outer scope as well fields / methods from the extends signatures are automatically part of the concrete implementation.
  implementation (
    <field or method declaration / definition>*
  )
)

<struct type> refers to the interface whereas !<struct type> refers to the concrete implementation of the struct type
```

interface signatures are unordered, but are ordered for the concrete implementation (ordered by parent decl, then interface decl, then implementation decl.  dups removed)

anonymous struct type definition within a function body must use the explicit form. e.g.,
```
func blah(i int, f float) (int, float) {
  struct(int, float)(i, f)
}
```

all struct fields are named, either implicitly or explicitly.  explicitly named fields must be uniquely named, whereas implicitly named field, which takes on the type's name, may be ambiguous as long as the fields are only accessed via pattern matching (error otherwise).

```
(int, float, third string) // implicitly named first field as int, second as float, third as third
```

support go style embedded field promotion

struct initialization: allow both positional and associative name mapping by default, which assigns the value to the corresponding fields.  positional and associative mapping intialization are operators and can be override (to support vectors / maps).  need option to disallow operators (vector shouldn't support map style initialization, etc)

idea?: structs are unordered.  struct A is equivalent to struct B if all of B's fields / methods B unambiguously maps to fields in fields / methods of A.  a named field of callable type in A equivalent to B's method if it has the same name and the method signature is a subtype of the callable's type.

comma separator only needed for declaring multiple fields on the same line.  newline acts as implicit separator.

method can be defined in struct body, or declare without body / defined outside of struct body


method signature need a way to specify constraint on self: pass by value, or pass by reference, or don't care

====

enum type
- enum value can have data
- if all enum values exposes the same (or eqv?) method signature, the method is directly acessible via the enum type (don't need to manually match the enum value before call the method)
- allow a default enum value
- "?" operator macro can operator on any enum type, rather than just result and option.  returns the value if it is the default value

```
type1|name2: type2|type3
```

is implicitly
```
enum(
  type1(type1),  // implicitly named
  name2(type2),  // explicitly named
  type3(type3),  // implicitly named
)
```
