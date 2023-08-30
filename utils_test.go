package utils

import "testing"

func TestDesensitizePhone(t *testing.T) {
	cases := []struct {
		phone string
		want  string
	}{
		{"13312345678", "133****5678"},
		{"1331234567", "1331234567"},
		{"133123456789", "133****56789"},
	}
	for _, c := range cases {
		if got := DesensitizePhone(c.phone); got != c.want {
			t.Errorf("DesensitizePhone(%q) == %q, want %q", c.phone, got, c.want)
		}
	}
}

func TestDesensitizePhoneScope(t *testing.T) {
	cases := []struct {
		phone string
		start int
		end   int
		want  string
	}{
		{"13312345678", 3, 7, "133****5678"},
		{"1331234567", 3, 7, "1331234567"},
		{"133123456789", 3, 7, "133****56789"},
		{"13312345678", 3, 9, "133******78"},
		{"13312345678", 3, 3, "13312345678"},
		{"13312345678", 3, 2, "13312345678"},
	}
	for _, c := range cases {
		if got := DesensitizePhoneScope(c.phone, c.start, c.end); got != c.want {
			t.Errorf("DesensitizePhoneScope(%q, %d, %d) == %q, want %q", c.phone, c.start, c.end, got, c.want)
		}
	}
}

func TestDesensitizeEmail(t *testing.T) {
	cases := []struct {
		email string
		want  string
	}{
		{"test@gmail.com", "t***@gmail.com"},
		{"testabc@gmail.com", "t******@gmail.com"},
	}
	for _, c := range cases {
		if got := DesensitizeEmail(c.email); got != c.want {
			t.Errorf("DesensitizeEmail(%q) == %q, want %q", c.email, got, c.want)
		}
	}
}

func TestDesensitizeName(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"张三", "张*"},
		{"张三丰", "张*丰"},
	}
	for _, c := range cases {
		if got := DesensitizeName(c.name); got != c.want {
			t.Errorf("DesensitizeName(%q) == %q, want %q", c.name, got, c.want)
		}
	}
}
