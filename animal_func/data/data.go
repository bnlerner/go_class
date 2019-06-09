package data

type Point struct {
	x float64
	y float64
}

/*
func (p Point) DistToOrigin() float64 {
	ret := math.Pow(p.x, 2) + math.Pow(p.y, 2)

	return math.Sqrt(ret)
}
*/
func (p *Point) InitPoint(xn, yn float64) {
	p.x = xn
	p.y = yn
}
