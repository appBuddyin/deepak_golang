package service_test

import (
	"calc-api/Service"
	"testing"
)

func TestOperationAdd(t *testing.T) {
	type args struct {
		operand  float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{
			name: "test_add_func",
			o:    Service.CreateOperationInstance(),
			args: args{
				operand:  5,
				operand2: 9.33,
			},
			want: 14.33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Add(tt.args.operand, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationSubtract(t *testing.T) {
	type args struct {
		operand  float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{
			name: "test_subtract_func",
			o:    Service.CreateOperationInstance(),
			args: args{
				operand:  10,
				operand2: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Subtract(tt.args.operand, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationMultiply(t *testing.T) {
	type args struct {
		operand  float64
		operand2 float64
	}
	tests := []struct {
		name string
		o    Service.Operations
		args args
		want float64
	}{
		{
			name: "test_multiply_func",
			o:    Service.CreateOperationInstance(),
			args: args{
				operand:  3,
				operand2: 7,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Multiply(tt.args.operand, tt.args.operand2); got != tt.want {
				t.Errorf("operation.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationDivide(t *testing.T) {
	type args struct {
		operand  float64
		operand2 float64
	}
	tests := []struct {
		name    string
		o       Service.Operations
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "test_divide_func",
			o:    Service.CreateOperationInstance(),
			args: args{
				operand:  10,
				operand2: 2,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "test_divide_by_zero",
			o:    Service.CreateOperationInstance(),
			args: args{
				operand:  10,
				operand2: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Divide(tt.args.operand, tt.args.operand2)
			if (err != nil) != tt.wantErr {
				t.Errorf("operation.Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("operation.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
