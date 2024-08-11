# Lexer Ideas

Invalid characters in source:
- '\0' null character
- '\a' alert/bell
- '\b' backspace
- '\f' form feed
- '\v' vertical tab
- ('\r' not part of a '\r\n' pair)

## Raw Tokens

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
