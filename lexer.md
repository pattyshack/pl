# Lexer Ideas


## Raw Tokens
All byte sequences are tokenizable to the following tokens.  Unexpected characters, and other lex errors such as missing string end delimiter, are captured by parse error tokens rather than lexer returning an error.

### Spaces
Sequence of ' ' or '\t'.  Only serve to separate tokens.  Spaces are ignored by multi-staged lexing.

### Newlines
Sequence of '\n' or '\r\n' ('\r' not paired with '\n' is not a newline).  Serve to separate tokens and may terminate a statement.

go explicitly inserts ';' to terminate statement.  we'll conditionally emit newline to the parser the same way. newline is emitted if the line's final token is
1. an identifier
2. a literal
3. a jump label
4. one of the keywords: `break`, `continue`, `fallthrough`, `return`, `true`, or `false`
5. one of: `++`, `--`, `)`, `}`, or `]`

### Comment
Raw comments are folded into other tokens via multi-staged lexing.

Line comments. Terminated by '\n', '\r', or EOF.  Note that the '\n'/'\r'/EOF is not part of the comment.
```
// single line comment
```

Block comments.  Begin with /* and end with */.  Block comments are scoped.
```
/*
block comment
  /*
  nested scope block comment.  The indentation is only for clarity
  */
*/
```

### Boolean Literals
- true
- false

### Integer Literals

Same as golang

- let hex = [0-9a-fA-F]

- decimal: 0|\[1-9\](\_?\d)*
- binary: 0\[bB\](\_?[01])+
- octal: 0\[oO\]?(\_?\[0-7\])+
- hexadecimal: 0\[xX\](\_?hex)+

### Float Literals
Same as golang

- (decimal integer)\\.(\d(\_?\d)*)?(\[eE\]\[+-\]?(decimal integer))?
- (decimal integer)\[eE\]\[+-\]?(decimal integer)
- \\.\d(\_?\d)*(\[eE\]\[+-\]?(decimal integer))?


- (hex integer)\\.(hex(\_?hex)*)*[pP][+-]?(hex(\_?hex)*)
- (hex integer)[pP][+-]?hex(\_?hex)*
- 0x\\.hex(\_?hex)*[pP][+-]?hex(\_?hex)*

### Rune Literals
Same as golang (except additional escape characters)

- '(non-escaped unicode char)'
- '(escaped char)'
- 1 byte oct rune '\\[0-8]{3}'
- 1 byte hex rune '\\x[0-9a-fA-F]{2}'
- 2 byte hex rune: '\\u[0-9a-fA-F]{4}'
- 4 byte hex rune: '\\U\[0-9a-fA-F\]{8}'

Special escaped characters
```
\a alert
\b backspace
\f form feed
\n newline
\r carriage return
\t horizontal tab
\v vertical tab
\\ backslash
\' single quote
\" double quote
\` back quote
\{newline} line continuation (only within multi-line strings)
```

### String Literal
Strings are sequence of characters surrounded by matching pair of string delimiter.  Single line strings are delimited using matching " or \'.
 Swift-style multi-line strings (string start on next line, leading indentation is strip based on spacing before the closing delimiter, etc.) are delimited using matching """ or \`\`\`.  There are two modes: interpreted vs. raw string.  Raw strings are prefixed by `r`.  Raw string does not interpret escape character sequence, whereas interpreted string does.

Escaped carriage return sequence "\\\\r" is preserved by interpreted string but unescaped carriage return '\r' is discarded from all string.

```
"escaped tab \t"
`can use " without escaping in back quote string, but must escape \` `
"can use ` without escaping in double quote string, but must escape \" "

"""
string can span
multiple lines
"""

"""
string can use " without escaping because of the triple-quote delimiter
"""  // similarly for backticks, etc.

r"raw string\n" == "raw string\\n"

r"""
multi-line
raw string with " character
\n are two separate characters
"""

r`can use " """ with backtick delimiter`

Note: no example for triple-backticks cuz github's markdown can't escape triple-backticks correctly


// The whitespace before the closing quotation marks tells the compiler what whitespace
// to ignore before all of the other lines.

abc = """
    My string
      line 2
        line \
    3
    ""

// is eqv to

asdf = """
My string
  line 2
    line \
3
"""

// is eqv to

xyz = """
My string
  line 2
    line 3
"""

// also eqv to "My string\n  line 2\n    line 3"
```

### Other potential literals?

Complex / Imaginary number
- probably not worth supporting as language feature.  just provide a std lib.  typed tuple with polymorphic operator maybe a good alternative.

### Literal Types
- Don't add built-in types into parser / lexer.  Specify "built-in" types as compiler schema, the types would provide api to determine how literal are converted to typed values.  A default type must be specified as part of the schema (for example, a integer literal can be an int, int16, int64. by default it should be an int)

### Identifier

```
letter = unicode-letter | '_'

identifer = letter (letter | unicode-number)*
```

### Keywords
```
and      as           async
break
case     continue
default  defer        do
else     enum
false    fallthrough  for       func
if       in           implements  import
let
not
or
package
return
struct   switch
trait    true         type
unsafe
var
```

### Operators and punctuations
```
// @ label decl and jump label
// _ identifier
// ' rune
// " string
// ` string
// unused ascii symbols: \ #

+     +=    ++
-     -=    --
*     *=
/     /=
%     %=
~     ~=    ~~
&     &=
|     |=
^     ^=
=     ==
!=    !     !!
<     <=    <<    <<=
>     >=    >>    >>=
(
)
[
]
{
}
.     ...
,
:
;
?
!
$[

```

XXX: consider adding overflow guard to arithmetic operators, and adding swift style overflow-able arithmetic operator &* &+ &-

### Label
Structured goto/jump label for `return` / `break` / `continue` statements.

```
<identifier>@    label declaration

@<identifier>    jump to label
```

e.g.,
```
var x = myscope@{
  if ... {
    for ... do {
      if ... {
        return @myscope 1
      }
      ...
    }
  }
}
```

```
outer@ for ... {
  for ... do {
    if ... {
      break @outer  // break the outer loop
    }
  }
}
```

```
outer@ for ... do {
  for ... do {
    if ... {
      continue @outer  // continue onto the next outer loop iteration
    }
  }
}
```
