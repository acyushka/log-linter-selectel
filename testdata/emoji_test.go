package analyzer_test

import (
	"loglinter/internal/analyzer"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEmoji(t *testing.T) {
	tests := []Test{
		{
			name: "без эмодзи - OK",
			msg:  "user logged in",
			want: []string{},
		},
		{
			name: "с эмодзи - ERROR",
			msg:  "user logged in 👍",
			want: []string{errHasEmoji},
		},
		{
			name: "только эмодзи - ERROR",
			msg:  "👍",
			want: []string{errHasNotLower, errHasEmoji},
		},
		{
			name: "флаг страны - ERROR",
			msg:  "user from 🇷🇺",
			want: []string{errHasEmoji},
		},
		{
			name: "несколько эмодзи - ERROR",
			msg:  "hello 👍🌍🎉",
			want: []string{errHasEmoji},
		},
		{
			name: "эмодзи и русский - ERROR",
			msg:  "привет 👍",
			want: []string{errHasEmoji, errHasNotEnglish},
		},
		{
			name: "восклицательный знак - ERROR",
			msg:  "hello!",
			want: []string{errHasEmoji},
		},
		{
			name: "вопросительный знак - ERROR",
			msg:  "hello?",
			want: []string{errHasEmoji},
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
