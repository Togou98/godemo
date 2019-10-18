package tmpconv
import "fmt"

type Ctmp float64
type Ftmp float64

const (
	Absctmp = -273.15
	Freezctmp = 0
	BoilingC = 100
	)

	func (c Ctmp) String() string { return fmt.Sprintf("%g C",c) }
	func (f Ftmp) String() string { return fmt.Sprintf("%g C",f) }
}

