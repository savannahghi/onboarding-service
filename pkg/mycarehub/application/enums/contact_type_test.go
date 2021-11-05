package enums

import (
	"bytes"
	"strconv"
	"testing"
)

func TestContactType_String(t *testing.T) {
	tests := []struct {
		name string
		e    ContactType
		want string
	}{
		{
			name: "PHONE",
			e:    PhoneContact,
			want: "PHONE",
		},
		{
			name: "EMAIL",
			e:    EmailContact,
			want: "EMAIL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("ContactType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    ContactType
		want bool
	}{
		{
			name: "valid type",
			e:    EmailContact,
			want: true,
		},
		{
			name: "invalid type",
			e:    ContactType("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("ContactType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContactType_UnmarshalGQL(t *testing.T) {
	pmtc := PhoneContact
	invalid := ContactType("invalid")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *ContactType
		args    args
		wantErr bool
	}{
		{
			name: "valid type",
			e:    &pmtc,
			args: args{
				v: "PHONE",
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			e:    &invalid,
			args: args{
				v: "this is not a valid type",
			},
			wantErr: true,
		},
		{
			name: "non string type",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ContactType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContactType_MarshalGQL(t *testing.T) {
	w := &bytes.Buffer{}
	tests := []struct {
		name  string
		e     ContactType
		b     *bytes.Buffer
		wantW string
		panic bool
	}{
		{
			name:  "valid type enums",
			e:     PhoneContact,
			b:     w,
			wantW: strconv.Quote("PHONE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.MarshalGQL(tt.b)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ContactType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
