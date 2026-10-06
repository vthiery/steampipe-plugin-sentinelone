package sentinelone

import (
	"testing"
	"time"
)

func intPtr(i int) *int { return &i }

func TestPageSize(t *testing.T) {
	cases := []struct {
		name     string
		cfg      sentineloneConfig
		fallback int
		want     int
	}{
		{"unset falls back to table default", sentineloneConfig{}, 1000, 1000},
		{"configured value wins", sentineloneConfig{PageSize: intPtr(200)}, 1000, 200},
		{"zero falls back", sentineloneConfig{PageSize: intPtr(0)}, 1000, 1000},
		{"negative falls back", sentineloneConfig{PageSize: intPtr(-5)}, 1000, 1000},
		{"above API max is clamped", sentineloneConfig{PageSize: intPtr(5000)}, 1000, 1000},
	}
	for _, c := range cases {
		if got := pageSize(c.cfg, c.fallback); got != c.want {
			t.Errorf("%s: pageSize(%v, %d) = %d, want %d", c.name, c.cfg.PageSize, c.fallback, got, c.want)
		}
	}
}

func TestRequestTimeout(t *testing.T) {
	cases := []struct {
		name string
		cfg  sentineloneConfig
		want time.Duration
	}{
		{"unset defaults to 30s", sentineloneConfig{}, 30 * time.Second},
		{"configured value wins", sentineloneConfig{RequestTimeout: intPtr(300)}, 300 * time.Second},
		{"zero defaults", sentineloneConfig{RequestTimeout: intPtr(0)}, 30 * time.Second},
		{"negative defaults", sentineloneConfig{RequestTimeout: intPtr(-1)}, 30 * time.Second},
	}
	for _, c := range cases {
		if got := requestTimeout(c.cfg); got != c.want {
			t.Errorf("%s: requestTimeout = %v, want %v", c.name, got, c.want)
		}
	}
}
