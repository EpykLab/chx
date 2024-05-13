/*
Copyright © 2024 EpykLab contact@epyklab.com
*/

package sources

import "testing"

func TestGetCrowdSecSmoke(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCrowdSecSmoke(tt.args.ip)
		})
	}
}
