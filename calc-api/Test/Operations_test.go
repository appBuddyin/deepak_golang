package Test

import (
	"calc-api/Service"
	"testing"
)

func Test_operation_add(t *testing.T) {
	type args struct {
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{name: "test_add_func",
			o:    Service.CreateOperationInstance(),
			args: args{operand1: 5, operand2: 9.33},
			want: 14.33},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Add(tt.args.operand1, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operation_subtract(t *testing.T) {
	type args struct {
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{name: "test_subtract_func",
			o:    Service.CreateOperationInstance(),
			args: args{operand1: 9.33, operand2: 4.33},
			want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Subtract(tt.args.operand1, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operation_multiply(t *testing.T) {
	type args struct {
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{name: "test_multiply_func",
			o:    Service.CreateOperationInstance(),
			args: args{operand1: 5, operand2: 9.33},
			want: 46.65},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Multiply(tt.args.operand1, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operation_divide(t *testing.T) {
	type args struct {
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{name: "test_divide_func",
			o:    Service.CreateOperationInstance(),
			args: args{operand1: 9.33, operand2: 3},
			want: 3.11},
			{
				name: "test_divide_func2",
				args: args{operand1: 10,operand2: 0},
				want: 0,
			},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Divide(tt.args.operand1, tt.args.operand2)
			if err != nil {
				t.Errorf("operation.Divide() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("operation.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
