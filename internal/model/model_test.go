package model

import "testing"

func TestModel_AdvanceTo(t *testing.T) {
	type fields struct {
		Position Position
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "advance within 6 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 3,
				Y: 3,
			}},
			wantErr: false,
		},
		{
			name: "advance over 6 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 6,
				Y: 6,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				Position: tt.fields.Position,
			}
			if err := m.AdvanceTo(tt.args.position); (err != nil) != tt.wantErr {
				t.Errorf("AdvanceTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_ChargeTo(t *testing.T) {
	type fields struct {
		Position Position
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "charge within 12 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 6,
				Y: 6,
			}},
			wantErr: false,
		},
		{
			name: "charge over 12 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 12,
				Y: 12,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				Position: tt.fields.Position,
			}
			if err := m.ChargeTo(tt.args.position); (err != nil) != tt.wantErr {
				t.Errorf("ChargeTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_RushTo(t *testing.T) {
	type fields struct {
		Position Position
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "rush within 12 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 6,
				Y: 6,
			}},
			wantErr: false,
		},
		{
			name: "rush over 12 inches",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{position: Position{
				X: 12,
				Y: 12,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				Position: tt.fields.Position,
			}
			if err := m.RushTo(tt.args.position); (err != nil) != tt.wantErr {
				t.Errorf("RushTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_moveTo(t *testing.T) {
	type fields struct {
		Position Position
	}
	type args struct {
		position Position
		limit    float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "move within limit",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{
				position: Position{
					X: 6,
					Y: 6,
				},
				limit: 100,
			},
			wantErr: false,
		},
		{
			name: "move over limit",
			fields: fields{Position: Position{
				X: 0,
				Y: 0,
			}},
			args: args{
				position: Position{
					X: 12,
					Y: 12,
				},
				limit: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				Position: tt.fields.Position,
			}
			if err := m.moveTo(tt.args.position, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("moveTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
