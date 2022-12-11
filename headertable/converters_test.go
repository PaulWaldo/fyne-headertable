package headertable

import (
	"testing"
	"time"
)

func TestDisplayAsString(t *testing.T) {
	testDate := time.Date(2000, 5, 12, 10, 12, 14, 216, time.UTC)
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "string type", args: args{i: "abc123"}, want: "abc123"},
		{name: "number type", args: args{i: -123.4522}, want: "%!s(float64=-123.4522)"},
		{name: "time type", args: args{i: testDate}, want: "2000-05-12 10:12:14.000000216 +0000 UTC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DisplayAsString(tt.args.i); got != tt.want {
				t.Errorf("DisplayAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisplayAsCurrency(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "round to cents", args: args{i: 123.4599}, want: "123.46"},
		{name: "display minus sign", args: args{i: -12.34}, want: "-12.34"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DisplayAsCurrency(tt.args.i); got != tt.want {
				t.Errorf("DisplayAsCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisplayAsISODate(t *testing.T) {
	testDate := time.Date(2000, 5, 12, 10, 12, 14, 216, time.UTC)
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "proper date", args: args{i: testDate}, want: "2000-05-12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DisplayAsISODate(tt.args.i); got != tt.want {
				t.Errorf("DisplayAsISODate() = %v, want %v", got, tt.want)
			}
		})
	}
}
