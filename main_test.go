package main

import (
	"testing"

	interpolator "github.com/judoDsl/interpolator"
)

func Test_joke_Execute(t *testing.T) {
	type fields struct {
		vars interpolator.VarsContent
	}
	tests := []struct {
		name    string
		fields  fields
		want    interpolator.VarsContent
		wantErr bool
	}{
		{name: "Should retrive a joke of Chuck Norris",
			fields:  fields{vars: interpolator.VarsContent{"category": "celebrity"}},
			want:    interpolator.VarsContent{"categories": "celebrity"},
			wantErr: false},
		{name: "Should fail, the empty category are not allowed",
			fields:  fields{vars: interpolator.VarsContent{}},
			want:    interpolator.VarsContent{},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			joker := joke{
				vars: tt.fields.vars,
			}
			got, err := joker.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("joke.Execute() error = %v, wantErr %v, name: %s", err, tt.wantErr, tt.name)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Errorf("joke.Execute() = %v, want %v", got, tt.want["categories"])
			}
		})
	}
}
