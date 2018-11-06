package report

import "testing"

func TestGenerate(t *testing.T) {
	Generate("test", "select *..", "title1, title2")
}
