package script

var codeTemplate = `package {{.PackageName}}

func {{.FuncName}}(input any) any {
    // TODO: implement your solution here
    return nil
}
`

var testTemplate = `package {{.PackageName}}

import (
	"testing"
	"reflect"
)

func {{.TestFuncName}}(t *testing.T) {
	tests := []struct {
		name   string
		input  any
		output any
	}{
		// Add your test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := {{.FuncName}}(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("{{.FuncName}}() = %v, want %v", got, tt.output)
			}
		})
	}
}
`
