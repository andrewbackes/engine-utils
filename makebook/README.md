# makebook

*Makebook* is a tool for creating polyglot opening books (bin format) from PGNs. 

### Usage
In the console:
`makebook [-f filter] pgn`

#### Example:
`makebook -f "BlackElo>=2700" -f "WhiteElo>=2700" -f "Result=1/2-1/2" -m 14 millionbase2.2.pgn`