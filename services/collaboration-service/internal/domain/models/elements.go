package models

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Color struct {
	Value uint32 `json:"value"`
}

type ElementStroke struct {
	Points []Point `json:"points"`
	Width  float64 `json:"width"`
	Color  Color   `json:"color"`
}

type ElementLine struct {
	Start Point   `json:"start"`
	End   Point   `json:"end"`
	Width float64 `json:"width"`
	Color Color   `json:"color"`
}

type ElementCircle struct {
	Center Point   `json:"center"`
	Radius float64 `json:"radius"`
	Width  float64 `json:"width"`
	Color  Color   `json:"color"`
}
