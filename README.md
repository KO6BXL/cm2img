# cm2img
Image creator for Circuit Maker2 written in Go.

# Usage
Run:
go run . -m $MODE -i $FILENAME

It will output to the terminal
To save to file:

go run . -m $MODE -i $FILENAME > output

If you are feeling cool and want to have small, high quality images, use the fine mode.

go run . -m fine -i $FILENAME

Or if you are feeling classy, use the normal mode to make each pixel one full block

go run . -m normal -i $FILENAME

NOTE: Only supports jpeg's and png's
