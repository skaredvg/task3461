package arithmetic

import (
	"reflect"
	"testing"
)

func TestParsing(t *testing.T) {
	type args struct {
		expr string
		tmpl string
	}
	tests := []struct {
		name    string
		args    args
		want    Statement
		wantErr bool
	}{
		{name: "5+3=?",
			args:    args{expr: "5+3=?", tmpl: `^([0-9]+)[[:space:]]*([\+\-*/]{1})[[:space:]]*([0-9]+)[[:space:]]*(=)[[:space:]]*([?])[[:space:]]*$`},
			want:    Statement{op1: 5, op2: 3, sign: "+"},
			wantErr: false},
		{name: "5+3=",
			args:    args{expr: "5+3=", tmpl: `^([0-9]+)[[:space:]]*([\+\-*/]{1})[[:space:]]*([0-9]+)[[:space:]]*(=)[[:space:]]*([?])[[:space:]]*$`},
			want:    Statement{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parsing(tt.args.expr, tt.args.tmpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parsing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parsing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	type args struct {
		st Statement
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
		wantErr bool
	}{
		{name: "5+3=8",
			args:    args{st: Statement{op1: 5, op2: 3, sign: "+"}},
			wantRes: 8,
			wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Calculate(tt.args.st)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Calculate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
