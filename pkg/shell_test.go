package pkg

import "testing"

func TestShellout(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Shellout(tt.args.command, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shellout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shellout() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Shellout() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReadFromCommand(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFromCommand(); got != tt.want {
				t.Errorf("ReadFromCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputError(t *testing.T) {
	type args struct {
		out    string
		errout string
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OutputError(tt.args.out, tt.args.errout, tt.args.err)
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		out    string
		errout string
		err    error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Output(tt.args.out, tt.args.errout, tt.args.err)
		})
	}
}
