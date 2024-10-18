package clockface_test

import (
	"bytes"
	"encoding/xml"
	clockface "hello/maths"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg" json:"svg,omitempty"`
	Xmlns   string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Width   string   `xml:"width,attr" json:"width,omitempty"`
	Height  string   `xml:"height,attr" json:"height,omitempty"`
	ViewBox string   `xml:"viewBox,attr" json:"viewbox,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Circle  Circle   `xml:"circle" json:"circle,omitempty"`
	Line    []Line   `xml:"line" json:"line,omitempty"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr" json:"cx,omitempty"`
	Cy float64 `xml:"cy,attr" json:"cy,omitempty"`
	R  float64 `xml:"r,attr" json:"r,omitempty"`
}

type Line struct {
	X1 float64 `xml:"x1,attr" json:"x1,omitempty"`
	Y1 float64 `xml:"y1,attr" json:"y1,omitempty"`
	X2 float64 `xml:"x2,attr" json:"x2,omitempty"`
	Y2 float64 `xml:"y2,attr" json:"y2,omitempty"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {

			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {

			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {

			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2024, time.January, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
