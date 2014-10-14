package uas

import "testing"

func benchmarkParser(b *testing.B, expectedName string, ua string) {
	agent := manifest.Parse(ua)
	if agent.BrowserVersion.Browser.Name != expectedName {
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		manifest.Parse(ua)
	}
}

// first browser regex in list
func BenchmarkParse_FindOperaMobileBrowser(b *testing.B) {
	ua := "Mozilla/5.0 (Linux; Android 2.3.4; MT11i Build/4.0.2.A.0.62) AppleWebKit/537.22 " +
		"(KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315"
	benchmarkParser(b, "Opera Mobile", ua)
}

// middle of the list browser regex
func BenchmarkParse_FindCurlBrowserWithForwardSlash(b *testing.B) {
	benchmarkParser(b, "cURL", "cURL/1.2.3.4")
}

// /WinHttp/si
// last browser regex in the list
func BenchmarkParse_FindLastWinHTTPBrowser(b *testing.B) {
	benchmarkParser(b, "WinHTTP", "WinHttp 1.2.3.4")
}
