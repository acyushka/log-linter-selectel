package analyzer_test

import (
	"loglinter/internal/analyzer"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnglish(t *testing.T) {
	tests := []Test{
		{
			name: "только английский - OK",
			msg:  "user logged in",
			want: []string{},
		},
		{
			name: "английский с цифрами - OK",
			msg:  "user logged in 123",
			want: []string{},
		},
		{
			name: "русские буквы - ERROR",
			msg:  "user вошел",
			want: []string{errHasNotEnglish},
		},
		{
			name: "смесь английского и русского - ERROR",
			msg:  "user вошел in",
			want: []string{errHasNotEnglish},
		},
		{
			name: "китайские иероглифы - ERROR",
			msg:  "user 你好",
			want: []string{errHasNotEnglish},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := analyzer.ValidateMsg(tt.msg)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ValidateMsg(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
