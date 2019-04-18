package mypackage

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}
func (r Rect) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// `Rect`의 값을 전달받는 method. `Rect`를 변경해도 원본 객체는 영향받지 않는다.
func (r Rect) area1() float64 { // (r Rect) 부분을 `receiver`를 지정한다 고 한다. `receiver`는 method 내에서 parameter 처럼 사용할 수 있다.
	r.Width++
	return r.Width * r.Height
}

// `Rect`의 참조를 전달받는 method. `Rect`를 변경하면 원본 객체에 영향을 준다.
func (r *Rect) area2() float64 {
	r.Width++
	return r.Width * r.Height
}

/*
func main() {
	r := Rect{10, 10}
	area := r.area()
	fmt.Println(r.Width, area)  // 10, 110
	area2 := r.area2()
	fmt.Println(r.Width, area2) // 11, 110
}
*/
