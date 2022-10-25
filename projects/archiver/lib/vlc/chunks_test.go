package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
        tests := []struct {
                name string
                args  args
                want BinaryChunks
        }{
                {
                        name: "split by chunks test",
			args: args{
				bStr: "001000100110100101",
				chunkSize: 8,
			},
                        want: BinaryChunks{"00100010", "01101001", "01000000"},
                },
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
                                t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
                        }
                })
        }
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "binary chunks to hex",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		data string
		want HexChunks
	}{
		{
			name: "hex chunks test",
			data: "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
