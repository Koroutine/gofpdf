package gofpdf

import (
	"testing"
)

func TestSVGBasicParse(t *testing.T) {

	svg, err := SVGBasicParse([]byte(`<svg height="200" width="350">
    <text x="150" y="80">Hello world</text>
    <text x="150" y="20">Hello again</text>
	<path d="M150 0 L75 200 L225 200 Z" />
</svg>`))

	if err != nil {
		t.Fatal(err)
	}

	init := InitType{
		OrientationStr: "P",
		UnitStr:        "pt",
		Size: SizeType{
			Wd: svg.Wd,
			Ht: svg.Ht,
		},
	}
	pdf := NewCustom(&init)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SVGBasicWrite(&svg, 1.0)
	err = pdf.OutputFileAndClose("pdf/SvgBasic.pdf")

	if err != nil {
		t.Fatal(err)
	}
}
