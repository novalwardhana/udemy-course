package helper

type Cube struct {
	Length float64
	Wide   float64
	Height float64
}

type CubeInterface interface {
	Volume() float64
	SurfaceArea() float64
}

func (c *Cube) Volume() float64 {
	return c.Length * c.Wide * c.Height
}

func (c *Cube) SurfaceArea() float64 {
	return 2*c.Length*c.Wide + 2*c.Wide*c.Height + 2*c.Length*c.Height
}
