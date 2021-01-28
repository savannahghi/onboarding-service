package domain_test

import (
	"bytes"
	"strconv"
	"testing"

	"gitlab.slade360emr.com/go/profile/pkg/onboarding/domain"
)

func TestKYCProcessStatus_String(t *testing.T) {
	tests := []struct {
		name string
		e    domain.KYCProcessStatus
		want string
	}{
		{
			name: "approved",
			e:    domain.KYCProcessStatusApproved,
			want: "APPROVED",
		},
		{
			name: "rejected",
			e:    domain.KYCProcessStatusRejected,
			want: "REJECTED",
		},
		{
			name: "pending",
			e:    domain.KYCProcessStatusPending,
			want: "PENDING",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("KYCProcessStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKYCProcessStatus_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    domain.KYCProcessStatus
		want bool
	}{
		{
			name: "valid approved",
			e:    domain.KYCProcessStatusApproved,
			want: true,
		},
		{
			name: "invalid kyc process status",
			e:    domain.KYCProcessStatus("this is not a real kyc process status"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("KYCProcessStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_UnmarshalGQL(t *testing.T) {
	valid := domain.KYCProcessStatusApproved
	invalid := domain.KYCProcessStatus("this is not a real kyc process status")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *domain.KYCProcessStatus
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "APPROVED",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not a real kyc process status",
			},
			wantErr: true,
		},
		{
			name: "non string",
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
				t.Errorf("KYCProcessStatus.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGender_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     domain.KYCProcessStatus
		wantW string
	}{
		{
			name:  "valid",
			e:     domain.KYCProcessStatusPending,
			wantW: strconv.Quote("PENDING"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("KYCProcessStatus.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestFivePointRating_String(t *testing.T) {
	tests := []struct {
		name string
		e    domain.FivePointRating
		want string
	}{
		{
			name: "poor",
			e:    domain.FivePointRatingPoor,
			want: "POOR",
		},
		{
			name: "unsatisfactory",
			e:    domain.FivePointRatingUnsatisfactory,
			want: "UNSATISFACTORY",
		},
		{
			name: "average",
			e:    domain.FivePointRatingAverage,
			want: "AVERAGE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("FivePointRating.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFivePointRating_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    domain.FivePointRating
		want bool
	}{
		{
			name: "valid",
			e:    domain.FivePointRatingPoor,
			want: true,
		},
		{
			name: "invalid",
			e:    domain.FivePointRating("this is not a real rating"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("FivePointRating.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFivePointRating_UnmarshalGQL(t *testing.T) {
	valid := domain.FivePointRatingPoor
	invalid := domain.FivePointRating("this is not a real rating")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *domain.FivePointRating
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "SATISFACTORY",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not a real five points rating",
			},
			wantErr: true,
		},
		{
			name: "non string",
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
				t.Errorf("FivePointRating.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFivePointRating_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     domain.FivePointRating
		wantW string
	}{
		{
			name:  "valid",
			e:     domain.FivePointRatingAverage,
			wantW: strconv.Quote("AVERAGE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("FivePointRating.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestBeneficiaryRelationship_String(t *testing.T) {
	tests := []struct {
		name string
		e    domain.BeneficiaryRelationship
		want string
	}{
		{
			name: "spouse",
			e:    domain.BeneficiaryRelationshipSpouse,
			want: "SPOUSE",
		},
		{
			name: "child",
			e:    domain.BeneficiaryRelationshipChild,
			want: "CHILD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeneficiaryRelationship_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    domain.BeneficiaryRelationship
		want bool
	}{
		{
			name: "valid",
			e:    domain.BeneficiaryRelationshipSpouse,
			want: true,
		},
		{
			name: "invalid",
			e:    domain.BeneficiaryRelationship("this is not real"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeneficiaryRelationship_UnmarshalGQL(t *testing.T) {
	valid := domain.BeneficiaryRelationshipSpouse
	invalid := domain.BeneficiaryRelationship("this is not real")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *domain.BeneficiaryRelationship
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "SPOUSE",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not real",
			},
			wantErr: true,
		},
		{
			name: "non string",
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
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBeneficiaryRelationship_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     domain.BeneficiaryRelationship
		wantW string
	}{
		{
			name:  "valid",
			e:     domain.BeneficiaryRelationshipSpouse,
			wantW: strconv.Quote("SPOUSE"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestIdentificationDocType_String(t *testing.T) {
	tests := []struct {
		name string
		e    domain.IdentificationDocType
		want string
	}{
		{
			name: "NATIONALID",
			e:    domain.IdentificationDocTypeNationalid,
			want: "NATIONALID",
		},
		{
			name: "PASSPORT",
			e:    domain.IdentificationDocTypePassport,
			want: "PASSPORT",
		},
		{
			name: "MILITARY",
			e:    domain.IdentificationDocTypeMilitary,
			want: "MILITARY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentificationDocType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    domain.IdentificationDocType
		want bool
	}{
		{
			name: "valid",
			e:    domain.IdentificationDocTypeMilitary,
			want: true,
		},
		{
			name: "invalid",
			e:    domain.IdentificationDocType("this is not real"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentificationDocType_UnmarshalGQL(t *testing.T) {
	valid := domain.IdentificationDocTypeNationalid
	invalid := domain.IdentificationDocType("this is not real")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *domain.IdentificationDocType
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "NATIONALID",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not real",
			},
			wantErr: true,
		},
		{
			name: "non string",
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
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIdentificationDocType_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     domain.IdentificationDocType
		wantW string
	}{
		{
			name:  "valid",
			e:     domain.IdentificationDocTypeNationalid,
			wantW: strconv.Quote("NATIONALID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestOrganizationType_String(t *testing.T) {
	tests := []struct {
		name string
		e    domain.OrganizationType
		want string
	}{
		{
			name: "LIMITED_COMPANY",
			e:    domain.OrganizationTypeLimitedCompany,
			want: "LIMITED_COMPANY",
		},
		{
			name: "TRUST",
			e:    domain.OrganizationTypeTrust,
			want: "TRUST",
		},
		{
			name: "UNIVERSITY",
			e:    domain.OrganizationTypeUniversity,
			want: "UNIVERSITY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganizationType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    domain.OrganizationType
		want bool
	}{
		{
			name: "valid",
			e:    domain.OrganizationTypeLimitedCompany,
			want: true,
		},
		{
			name: "invalid",
			e:    domain.OrganizationType("this is not real"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganizationType_UnmarshalGQL(t *testing.T) {
	valid := domain.OrganizationTypeUniversity
	invalid := domain.OrganizationType("this is not real")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *domain.OrganizationType
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "UNIVERSITY",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not real",
			},
			wantErr: true,
		},
		{
			name: "non string",
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
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrganizationType_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     domain.OrganizationType
		wantW string
	}{
		{
			name:  "valid",
			e:     domain.OrganizationTypeTrust,
			wantW: strconv.Quote("TRUST"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
