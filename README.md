# cm2img
Image creator for Circuit Maker2 written in Go.

# Usage
Run:
go run . normal $FILENAME

It will output to the terminal
To save to file:

go run . normal $FILENAME > output

If you are feeling cool and want to have small, high quality images, use the fine mode.

go run . fine $FILENAME

NOTE: Only supports jpeg's and png's