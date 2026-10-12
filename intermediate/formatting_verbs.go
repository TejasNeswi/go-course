package intermediate

import "fmt"

func main() {

	i := 111_505.5
	string := "hello world"

	fmt.Printf("%v\n", i)  // Prints the value in default format
	fmt.Printf("%#v\n", i) // Prints the value in go syntax format
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	int := 18
	fmt.Printf("%b\n", int)   // binary
	fmt.Printf("%d\n", int)   // base 10
	fmt.Printf("%+d\n", int)  // signed in base 10
	fmt.Printf("%o\n", int)   // base 8
	fmt.Printf("%O\n", int)   // base 8 adds 0o prefix
	fmt.Printf("%x\n", int)   // hex in lowercase like ff
	fmt.Printf("%X\n", int)   // hex in uppercase like FF
	fmt.Printf("%#x\n", int)  // hex with 0x prefix
	fmt.Printf("%4d\n", int)  // right justified
	fmt.Printf("%-4d\n", int) // left justified
	fmt.Printf("%04d\n", int) // zero prefix

	txt := "world"

	fmt.Printf("%s\n", txt)   // plain string
	fmt.Printf("%q\n", txt)   // double quotted string
	fmt.Printf("%8s\n", txt)  // plain string width 8 right justified
	fmt.Printf("%-8s\n", txt) // plain string width 8 left justified
	fmt.Printf("%x\n", txt)   // hex
	fmt.Printf("% x\n", txt)  // hex with spaces

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)


	flt:=9.18

	fmt.Printf("%e\n", flt) // scientific notation
	fmt.Printf("%f\n", flt) // Decimanl point no exponent
	fmt.Printf("%.2f\n", flt) // Default width precision 2
	fmt.Printf("%6.2f\n", flt) // width 6 precision 2
	fmt.Printf("%g\n", flt) // exponent as needed only necessary digits

}
