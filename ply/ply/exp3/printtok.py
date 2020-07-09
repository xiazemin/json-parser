import lex
# Test it out
data = '''
3 + 4 * 10
  + -20 *2
'''

# Give the lexer some input
lex.lexer.input(data)

# Tokenize
while True:
    tok = lex.lexer.token()
    if not tok: break      # No more input
    print tok