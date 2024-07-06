# Arcane

Arcane is an Interpreter written in Go.

## File Structure
### Token
- Defines constants representing the lexical tokens of Arcane.

### AST (Abstract Syntax Tree)
- Declares the types used to represent the syntax tree.

### Lexer
- Takes the source code as input and output the tokens that represent the source code.
- Passes these tokens to the parser.

### Parser
- Uses a recursive decent parser, specifically the **Top Down Operator Precedence** (Pratt Parser) by Vaughan Pratt. [More Info](https://tdop.github.io)
- Takes the input from Lexer and builds the **AST** from it.

### REPL (Read Eval Print Loop)
- Similar to `console` or `interactive mode` in other programming languages.
- Reads input, send it to the interpreter to evaluation, print the result, and start again.