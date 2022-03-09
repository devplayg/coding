#!/bin/sh

go build main.go

#./main <<EOF
#3 3
#010
#010
#010
#EOF

./main <<EOF
6 4
0100
1110
1000
0000
0111
0000
EOF
#
#./main <<EOF
#4 4
#0111
#1111
#1111
#1110
#EOF