package foo

type StructFoo struct { //@item(StructFoo, "StructFoo", "struct{...}", "struct")
	Value int //@item(Value, "Value", "int", "field")
}

<<<<<<< HEAD
// TODO(rstambler): Create pre-set builtins?
=======
// Pre-set this marker, as we don't have a "source" for it in this package.
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
/* Error() */ //@item(Error, "Error()", "string", "method")

func Foo() { //@item(Foo, "Foo()", "", "func")
	var err error
	err.Error() //@complete("E", Error)
}

func _() {
	var sFoo StructFoo           //@complete("t", StructFoo)
	if x := sFoo; x.Value == 1 { //@complete("V", Value),typdef("sFoo", StructFoo)
		return
	}
}

<<<<<<< HEAD
//@complete("", Foo, IntFoo, StructFoo)
type IntFoo int //@item(IntFoo, "IntFoo", "int", "type")
=======
func _() {
	shadowed := 123
	{
		shadowed := "hi" //@item(shadowed, "shadowed", "string", "var")
		sha              //@complete("a", shadowed)
	}
}

type IntFoo int //@item(IntFoo, "IntFoo", "int", "type"),complete("", Foo, IntFoo, StructFoo)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
