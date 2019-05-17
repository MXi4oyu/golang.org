// A comment just to push the positions out

package a

<<<<<<< HEAD
=======
import "fmt"

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
type A string //@A

func Stuff() { //@Stuff
	x := 5
	Random2(x) //@godef("dom2", Random2)
	Random()   //@godef("()", Random)
<<<<<<< HEAD
=======

	var err error         //@err
	fmt.Printf("%v", err) //@godef("err", err)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
