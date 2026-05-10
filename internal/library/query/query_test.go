package query

import "testing"

func TestPageParamDefaultsAndBounds(t *testing.T) {
	t.Parallel()

	param := &PageParam{Page: -1, Limit: -10}
	if got := param.GetPageIndex(); got != 1 {
		t.Fatalf("GetPageIndex() = %d, want 1", got)
	}
	if got := param.GetPageSize(); got != 20 {
		t.Fatalf("GetPageSize() = %d, want 20", got)
	}

	param = &PageParam{Page: 2, Limit: 101}
	if got := param.GetPageIndex(); got != 2 {
		t.Fatalf("GetPageIndex() = %d, want 2", got)
	}
	if got := param.GetPageSize(); got != 100 {
		t.Fatalf("GetPageSize() = %d, want 100", got)
	}
}

func TestPageParamGetOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		param PageParam
		want  string
	}{
		{
			name:  "safe field",
			param: PageParam{Sort: "sys_user.id", Order: "DESC"},
			want:  "sys_user.id desc",
		},
		{
			name:  "invalid direction",
			param: PageParam{Sort: "id", Order: "delete"},
			want:  "",
		},
		{
			name:  "unsafe field",
			param: PageParam{Sort: "id desc;drop table sys_user", Order: "asc"},
			want:  "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.param.GetOrder(); got != tt.want {
				t.Fatalf("GetOrder() = %q, want %q", got, tt.want)
			}
		})
	}
}
