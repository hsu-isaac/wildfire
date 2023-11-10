package handler

import (
	"testing"
	"wildfire/api/services"
)

func Test_combineNameAndJoke(t *testing.T) {
	type args struct {
		joke string
		name services.Name
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Happy path",
			args: args{
				joke: "John Doe uses canvas in IE.",
				name: services.Name{
					FirstName: "Isaac",
					LastName:  "Hsu",
				},
			},
			want:    "Isaac Hsu uses canvas in IE.",
			wantErr: false,
		},
		{
			name: "Incorrect joke parameter, missing 'John Doe'",
			args: args{
				joke: "firstName lastName uses canvas in IE.",
				name: services.Name{
					FirstName: "Isaac",
					LastName:  "Hsu",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := combineNameAndJoke(tt.args.joke, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("combineNameAndJoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("combineNameAndJoke() = %v, want %v", got, tt.want)
			}
		})
	}
}
