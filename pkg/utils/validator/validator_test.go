package validator

import (
	"github.com/stretchr/testify/assert"
	"go-skeleton/internal/model/enum"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		in   interface{}
		want string
	}{
		{
			in: struct {
				Input string `json:"input" validate:"job_validator"`
			}{Input: "test"},
			want: "Pekerjaan (test) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"job_validator"`
			}{Input: enum.REGISTRATION_JOB_PNS.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"education_validator"`
			}{Input: "test"},
			want: "Edukasi (test) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"education_validator"`
			}{Input: enum.REGISTRATION_EDUCATION_TIDAK_SEKOLAH.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"marital_status_validator"`
			}{Input: "test"},
			want: "Status Pernikahan (test) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"marital_status_validator"`
			}{Input: enum.REGISTRATION_MARITAL_BELUM_KAWIN.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"religion_validator"`
			}{Input: "test"},
			want: "Agama (test) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"religion_validator"`
			}{Input: enum.REGISTRATION_RELIGION_ISLAM.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"gender_validator"`
			}{Input: "test"},
			want: "Jenis kelamin (test) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"gender_validator"`
			}{Input: enum.REGISTRATION_MALE.String()},
			want: "",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"0000000", "123123123"}},
			want: "IP (0000000) tidak valid",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"0.0.0.0", "123123123"}},
			want: "IP (0.0.0.0) tidak valid",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"127.0.1.1", "123123123"}},
			want: "IP (123123123) tidak valid",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"hello world"}},
			want: "IP (hello world) tidak valid",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"::1"}},
			want: "IP (::1) tidak valid",
		},
		{
			in: struct {
				Input []string `json:"input" validate:"whitelist_ip_validator"`
			}{Input: []string{"127.0.0.1"}},
			want: "",
		},
		{
			in: struct {
				Input float64 `json:"input" validate:"multiple_of_validator"`
			}{Input: 10000},
			want: "Nominal (10000) harus kelipatan 50000",
		},
		{
			in: struct {
				Input float64 `json:"input" validate:"multiple_of_validator"`
			}{Input: 125000},
			want: "Nominal (125000) harus kelipatan 50000",
		},
		{
			in: struct {
				Input float64 `json:"input" validate:"multiple_of_validator"`
			}{Input: 100000},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"alphaspace_validator"`
			}{Input: "qweqwewe123123"},
			want: "Value (qweqwewe123123) tidak valid, hanya boleh mengandung karakter alphabet dan spasi",
		},
		{
			in: struct {
				Input string `json:"input" validate:"alphaspace_validator"`
			}{Input: "123123"},
			want: "Value (123123) tidak valid, hanya boleh mengandung karakter alphabet dan spasi",
		},
		{
			in: struct {
				Input string `json:"input" validate:"alphaspace_validator"`
			}{Input: "qweqwewe 123123"},
			want: "Value (qweqwewe 123123) tidak valid, hanya boleh mengandung karakter alphabet dan spasi",
		},
		{
			in: struct {
				Input string `json:"input" validate:"alphaspace_validator"`
			}{Input: "hello world"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "qweqwewe123123"},
			want: "Value (qweqwewe123123) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "qweqwewe 123123"},
			want: "Value (qweqwewe 123123) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "qweqwewe.123123"},
			want: "Value (qweqwewe.123123) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "123123qweqwewe"},
			want: "Value (123123qweqwewe) tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "hello world"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "hello world, SPD"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "hello world,. SPD"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"name_validator"`
			}{Input: "hello world., SPD"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"date_yyyymmdd_validator"`
			}{Input: "20220101"},
			want: "Tanggal (20220101) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"date_yyyymmdd_validator"`
			}{Input: "hello world"},
			want: "Tanggal (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"date_yyyymmdd_validator"`
			}{Input: "2022-01-01"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"callback_event_type_validator"`
			}{Input: "hello world"},
			want: "Callback event type (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"callback_event_type_validator"`
			}{Input: enum.CALLBACK_EVENT_LOAN_APPROVED.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"registration_status_validator"`
			}{Input: "hello world"},
			want: "Registation status (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"registration_status_validator"`
			}{Input: enum.REGISTRATION_STATUS_REQUESTED.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"emergency_contact_relationship_validator"`
			}{Input: "hello world"},
			want: "Emergency contact relationship (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"emergency_contact_relationship_validator"`
			}{Input: enum.EMERGENCY_CONTACT_RELATIONSHIP_SUAMI.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"registration_image_type_validator"`
			}{Input: "hello world"},
			want: "Jenis foto (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"registration_image_type_validator"`
			}{Input: enum.REGISTRATION_IMAGE_TYPE_LINK_KTP.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"se_type_validator"`
			}{Input: "hello world"},
			want: "SE type (hello world) tidak valid.",
		},
		{
			in: struct {
				Input string `json:"input" validate:"se_type_validator"`
			}{Input: enum.SE_TYPE_ONLINE.String()},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "hello world"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "helloworld"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "hello1"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "hello1"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "12hello"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "%$heahaha"},
			want: "Password tidak valid. Minimum 8 karakter dan maksimum 64 karakter (alphanumeric)",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "gasd%13%$"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"password_validator"`
			}{Input: "helloworld123"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "-91"},
			want: "Latitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "91"},
			want: "Latitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "+91"},
			want: "Latitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "-19"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "-90"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "90"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"latitude_validator"`
			}{Input: "+90"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "+181"},
			want: "Longitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "181"},
			want: "Longitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "-181"},
			want: "Longitude tidak valid",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "-19"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "-180"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "180"},
			want: "",
		},
		{
			in: struct {
				Input string `json:"input" validate:"longitude_validator"`
			}{Input: "+180"},
			want: "",
		},
	}

	for i, tc := range testCases {
		msg, err := Validate(tc.in)
		if err != nil {
			assert.NotNil(t, err, "Validate: should be not nil at index %v", i)
		} else {
			assert.Nil(t, err, "Validate: should be nil at index %v", i)
		}
		assert.Equal(t, tc.want, msg, "Validate: should be equals at index %v", i)
	}
}
