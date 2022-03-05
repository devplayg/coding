#!/bin/sh

go build main.go

./main <<EOF
4 4
ABCD
EFAD
ASDA
NGHZ
EOF

./main <<EOF
2 4
CAAB
ADCB
EOF


./main <<EOF
3 6
HFDFFB
AJHGDH
DGAGEH
EOF

./main <<EOF
5 5
IEFCJ
FHFKC
FFALF
HFGCF
HMCHH
EOF
