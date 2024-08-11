# Lexer Ideas

Invalid characters in source:
- '\0' null character
- '\a' alert/bell
- '\b' backspace
- '\f' form feed
- '\v' vertical tab
- ('\r' not part of a '\r\n' pair)

TODO: improve non-ascii unicode handling

## Tokens

### Spaces
Sequence of ' ' or '\t'.  Only serve to separate tokens.  Spaces are ignored by multi-staged lexing.

### Newlines
Sequence of '\n' or '\r\n'.  Serve to separate tokens and may terminate a statement.

XXX: go explicitly inserts ';' to terminate statement.  maybe make lexer smarter and conditionally remove newlines instead.

### Comment
Raw comments are folded into other tokens via multi-staged lexing.

Line comments. Terminated by endline or EOF.  Note that the newline/EOF is not part of the comment.
```
// single line comment
```

Block comments.  Begin with /* and end with */.  Block comments are scoped. 
```
/*
block comment 
/*
nested scope block comment
*/
*/
```

### Boolean Literals
- true
- false

### Integer Literals
- decimal: 0|\[1-9\](\_?\[0-9\])*
- binary: 0\[bB\](\_?[01])+
- octal: 0\[oO\]?(\_?\[0-7\])+
- hexadecimal: 0\[xX\](\_?[0-9a-fA-F])+

### Float Literals
- (decimal integer)\\.(\[0-9\](\_?\[0-9\])*)?(\[eE\]\[+-\]?(decimal integer))?
- (decimal integer)\[eE\]\[+-\]?(decimal integer)
- \\.\[0-9\](\_?\[0-9\])*(\[eE\]\[+-\]?(decimal integer))?

XXX: support hex float representation?

### Rune Literals
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
\" double quote (only within string literals)
```

### String Literal
Strings are sequence of characters surrounded by matching pair of string delimiter.  Valid delimiters are ", """, \`, and \`\`\`.  There are two forms: interpreted vs. raw string.  Raw strings are prefixed by `r`.  Raw string does not interpret escape character sequence, whereas interpreted string does.

In all cases, string may span multiple lines, i.e., contain unescaped newline '\n'.  Escaped carriage return sequence "\\\\r" is preserved by interpreted string but unescaped carriage return '\r' is discarded from all string.

```
"escaped tab \t"

`string can span
multiple lines`

"""string can use " without escaping because of the triple-quote delimiter"""  // similarly for backticks, etc.

r"raw string\n" == "raw string\\n"

r"""multi-line
raw string with "
character"""

r`can use " """ with backtick delimiter`

Note: no example for triple-backticks cuz github's markdown can't escape triple-backticks correctly
```

### Other potential literals?

Complex / Imaginary number
- probably not worth supporting as language feature.  just provide a std lib.  typed tuple with polymorphic operator maybe a good alternative.

### Literal Types
- Don't add built-in types into parser / lexer.  Specify "built-in" types as compiler schema, the types would provide api to determine how literal are converted to typed values.  A default type must be specified as part of the schema (for example, a integer literal can be an int, int16, int64. by default it should be an int)
