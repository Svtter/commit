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
			r := Shellout(tt.args.command, tt.args.args...)
			if (r.err != nil) != tt.wantErr {
				t.Errorf("Shellout() error = %v, wantErr %v", r.err, tt.wantErr)
				return
			}
			if r.out != tt.want {
				t.Errorf("Shellout() got = %v, want %v", r.out, tt.want)
			}
			if r.errout != tt.want1 {
				t.Errorf("Shellout() got1 = %v, want %v", r.errout, tt.want1)
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
		r RunResult
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r.OutputError()
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		r RunResult
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r.Output()
		})
	}
}
