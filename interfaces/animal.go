package interfaces

type Animal interface {
	Weight() float64
}

type Cow struct {
	mass float64
}

type Whale struct {
	mass        float64
	waterWeight float64
}

func (c *Cow) Weight() float64 {
	return c.mass * 9.8
}

func (w *Whale) Weight() float64 {
	return w.mass*9.8 - w.waterWeight
}

func CalculateWeight(a Animal) float64 {
	return a.Weight()
}
