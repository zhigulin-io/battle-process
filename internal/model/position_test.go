package model

import "testing"

func TestPosition_DistanceTo(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "distance zero to positive position",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{position: Position{
				X: 3,
				Y: 4,
			}},
			want: 5,
		},
		{
			name: "distance from zero to negative position",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{position: Position{
				X: -3,
				Y: -4,
			}},
			want: 5,
		},
		{
			name: "distance from non zero to positive position",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{position: Position{
				X: 4,
				Y: 5,
			}},
			want: 5,
		},
		{
			name: "distance from non zero to negative position",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{position: Position{
				X: -2,
				Y: -3,
			}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.DistanceTo(tt.args.position); got != tt.want {
				t.Errorf("DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
