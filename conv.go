package tmpconv

func Ctof( c Ctmp) Ftmp { return Ftmp( c * 9 / 5 + 32)}
func Ftoc (f Ftmp) Ctmp { return Ctmp( (f - 32) * 5 /9 )}

