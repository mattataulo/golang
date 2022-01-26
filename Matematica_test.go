package main

import "testing"

func Test_calcular(t *testing.T) {
	type args struct {
		operacao string
		numero1  int
		numero2  int
	}
	tests := []struct {
		name      string
		args      args
		wantSaida int
	}{
		{
			name: "Soma",
			args: args{
				operacao: "+",
				numero1:  1,
				numero2:  2,
			},
			wantSaida: 3,
		},
		{
			name: "Subtração",
			args: args{
				operacao: "-",
				numero1:  5,
				numero2:  3,
			},
			wantSaida: 2,
		}, {
			name: "Divisão",
			args: args{
				operacao: "/",
				numero1:  10,
				numero2:  5,
			},
			wantSaida: 2,
		}, {
			name: "Multiplicação",
			args: args{
				operacao: "*",
				numero1:  10,
				numero2:  2,
			},
			wantSaida: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSaida := calcular(tt.args.operacao, tt.args.numero1, tt.args.numero2); gotSaida != tt.wantSaida {
				t.Errorf("calcular() = %v, want %v", gotSaida, tt.wantSaida)
			}
		})
	}
}
