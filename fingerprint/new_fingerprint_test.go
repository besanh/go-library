package fingerprint

import (
	"net/http"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestSanitize_SingleString(t *testing.T) {
	// Printable ASCII remains unchanged
	input := "HelloWorld123"
	out := Sanitize(input)
	if out != input {
		t.Errorf("Sanitize(%q) = %q; want %q", input, out, input)
	}

	// Control character \x01 replaced by tofu U+FFFE
	input = "A\x01B"
	out = Sanitize(input)
	if !strings.ContainsRune(out, '􏿮') {
		t.Errorf("Sanitize(%q) = %q; missing tofu replacement", input, out)
	}

	// Invalid rune (beyond MaxRune) replaced by replacement char
	input = string([]rune{utf8.MaxRune + 1})
	out = Sanitize(input)
	if !strings.ContainsRune(out, '�') {
		t.Errorf("Sanitize(%q) = %q; missing replacement char", input, out)
	}
}

func TestSanitize_MultipleStrings(t *testing.T) {
	s1, s2 := "foo", "bar"
	out := Sanitize(s1, s2)
	// Expect format: [foo, bar]
	if !strings.HasPrefix(out, "[") || !strings.HasSuffix(out, "]") {
		t.Errorf("Sanitize(slice) = %q; want bracketed list", out)
	}
	if !strings.Contains(out, "foo") || !strings.Contains(out, "bar") {
		t.Errorf("Sanitize(slice) = %q; missing elements", out)
	}
}

func TestSafeHeader_NoHeader(t *testing.T) {
	r := &http.Request{Header: http.Header{}}
	if got := SafeHeader(r, "X-Not-Here"); got != "" {
		t.Errorf("SafeHeader(no header) = %q; want empty", got)
	}
}

func TestSafeHeader_SingleValue(t *testing.T) {
	r := &http.Request{Header: http.Header{"X-Test": {"val1"}}}
	out := SafeHeader(r, "X-Test")
	if out != "val1" {
		t.Errorf("SafeHeader(single) = %q; want %q", out, "val1")
	}
}

func TestSafeHeader_MultipleValues(t *testing.T) {
	r := &http.Request{Header: http.Header{"X-Multi": {"a", "b"}}}
	out := SafeHeader(r, "X-Multi")
	if !strings.HasPrefix(out, "[") || !strings.Contains(out, "a") || !strings.Contains(out, "b") {
		t.Errorf("SafeHeader(multi) = %q; invalid format", out)
	}
}

func TestRequestFingerprint(t *testing.T) {
	r := &http.Request{Header: http.Header{
		"Accept-Language":           {"en-US"},
		"User-Agent":                {"Go-http-client/1.1"},
		"Referer":                   {"http://example.com"},
		"Accept":                    {"text/html"},
		"Accept-Encoding":           {"gzip, deflate"},
		"Connection":                {"keep-alive"},
		"Cache-Control":             {"no-cache"},
		"Upgrade-Insecure-Requests": {"1"},
		"Via":                       {"1.1 proxy"},
		"DNT":                       {"1"},
		"Cookie":                    {"session=abc"},
	}}
	r.RemoteAddr = "127.0.0.1:1234"
	line := RequestFingerprint(r)

	checks := []string{
		"en-US",
		"Go-http-client/1.1",
		"R=http://example.com",
		"A=text/html",
		"E=gzip, deflate",
		"keep-alive",
		"no-cache",
		"Via=1.1 proxy",
		"session=abc",
	}
	for _, chk := range checks {
		if !strings.Contains(line, chk) {
			t.Errorf("RequestFingerprint missing %q; got %q", chk, line)
		}
	}
	// Both UIR and DNT headers with value "1" should be omitted
	if strings.Contains(line, "UIR=") || strings.Contains(line, "DNT") {
		t.Errorf("RequestFingerprint should skip UIR and DNT; got %q", line)
	}
}

func TestFingerprintMD(t *testing.T) {
	r := &http.Request{Header: http.Header{
		"Accept":     {"application/json"},
		"User-Agent": {"TestAgent"},
	}}
	r.RemoteAddr = "10.0.0.1:8080"
	md := FingerprintMD(r)
	// Should contain IP and header lines
	if !strings.Contains(md, "**IP**: 10.0.0.1:8080") {
		t.Errorf("FingerprintMD missing IP; got %q", md)
	}
	if !strings.Contains(md, "- **Accept**: application/json") {
		t.Errorf("FingerprintMD missing Accept; got %q", md)
	}
	if !strings.Contains(md, "- **User-Agent**: TestAgent") {
		t.Errorf("FingerprintMD missing User-Agent; got %q", md)
	}
}

func TestIPMethodURL(t *testing.T) {
	r := &http.Request{Method: "POST", RequestURI: "/path?x=1"}
	r.RemoteAddr = "192.168.0.1:5555"
	out := IPMethodURL(r)
	expected := "--> 192.168.0.1:5555 POST /path?x=1"
	if out != expected {
		t.Errorf("IPMethodURL = %q; want %q", out, expected)
	}
}
