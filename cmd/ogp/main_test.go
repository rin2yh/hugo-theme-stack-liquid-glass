package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"testing"
)

// pngDims decodes a PNG header and returns its pixel dimensions.
func pngDims(t *testing.T, path string) (int, int) {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer f.Close()
	cfg, err := png.DecodeConfig(f)
	if err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	return cfg.Width, cfg.Height
}

// writeSwatch writes a small test image using enc and returns its path.
func writeSwatch(t *testing.T, path string, enc func(*os.File, image.Image) error) string {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			img.SetRGBA(x, y, color.RGBA{uint8(x * 4), uint8(y * 4), 128, 255})
		}
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create %s: %v", path, err)
	}
	defer f.Close()
	if err := enc(f, img); err != nil {
		t.Fatalf("encode %s: %v", path, err)
	}
	return path
}

// TestRunDefaultCanvas checks the brand-neutral default renders at 1200x630.
func TestRunDefaultCanvas(t *testing.T) {
	out := filepath.Join(t.TempDir(), "base.png")
	if err := run(out, "", "", ""); err != nil {
		t.Fatalf("run: %v", err)
	}
	if w, h := pngDims(t, out); w != width || h != height {
		t.Fatalf("got %dx%d, want %dx%d", w, h, width, height)
	}
}

// TestRunAcceptsAvatarFormats is a regression test for the -avatar flag: the
// flag help and README advertise png/jpg/webp, so decoding must not fail for a
// JPEG avatar (previously the image/jpeg decoder was not registered).
func TestRunAcceptsAvatarFormats(t *testing.T) {
	dir := t.TempDir()
	pngPath := writeSwatch(t, filepath.Join(dir, "a.png"), func(f *os.File, im image.Image) error {
		return png.Encode(f, im)
	})
	jpgPath := writeSwatch(t, filepath.Join(dir, "a.jpg"), func(f *os.File, im image.Image) error {
		return jpeg.Encode(f, im, nil)
	})
	for _, avatar := range []string{pngPath, jpgPath} {
		out := filepath.Join(dir, "out.png")
		if err := run(out, "", "", avatar); err != nil {
			t.Fatalf("run with avatar %s: %v", avatar, err)
		}
		if w, h := pngDims(t, out); w != width || h != height {
			t.Fatalf("avatar %s: got %dx%d, want %dx%d", avatar, w, h, width, height)
		}
	}
}

// TestRunUnknownAvatarFormatErrors confirms an unsupported avatar surfaces an
// error rather than silently producing a background.
func TestRunUnknownAvatarFormatErrors(t *testing.T) {
	dir := t.TempDir()
	bad := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(bad, []byte("not an image"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(filepath.Join(dir, "out.png"), "", "", bad); err == nil {
		t.Fatal("expected an error for a non-image avatar, got nil")
	}
}

// TestRunWithSiteName exercises the text path using the bundled font.
func TestRunWithSiteName(t *testing.T) {
	fontPath := filepath.FromSlash("../../assets/ogp/NotoSansJP-Bold.ttf")
	if _, err := os.Stat(fontPath); err != nil {
		t.Skipf("bundled font not available: %v", err)
	}
	out := filepath.Join(t.TempDir(), "out.png")
	if err := run(out, fontPath, "テスト Blog", ""); err != nil {
		t.Fatalf("run: %v", err)
	}
	if w, h := pngDims(t, out); w != width || h != height {
		t.Fatalf("got %dx%d, want %dx%d", w, h, width, height)
	}
}

func TestClamp(t *testing.T) {
	cases := []struct{ v, lo, hi, want float64 }{
		{-1, 0, 1, 0},
		{2, 0, 1, 1},
		{0.5, 0, 1, 0.5},
		{0, 0, 1, 0},
		{1, 0, 1, 1},
	}
	for _, c := range cases {
		if got := clamp(c.v, c.lo, c.hi); got != c.want {
			t.Errorf("clamp(%v,%v,%v)=%v, want %v", c.v, c.lo, c.hi, got, c.want)
		}
	}
}

func TestSmoothstep(t *testing.T) {
	if got := smoothstep(0, 1, 0); got != 0 {
		t.Errorf("smoothstep at low edge = %v, want 0", got)
	}
	if got := smoothstep(0, 1, 1); got != 1 {
		t.Errorf("smoothstep at high edge = %v, want 1", got)
	}
	if got := smoothstep(0, 1, 0.5); math.Abs(got-0.5) > 1e-9 {
		t.Errorf("smoothstep midpoint = %v, want 0.5", got)
	}
	if got := smoothstep(0, 1, -5); got != 0 {
		t.Errorf("smoothstep below range = %v, want 0 (clamped)", got)
	}
	if got := smoothstep(0, 1, 5); got != 1 {
		t.Errorf("smoothstep above range = %v, want 1 (clamped)", got)
	}
}

func TestLerpBounds(t *testing.T) {
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	if got := lerp(black, white, 0); got != black {
		t.Errorf("lerp t=0 = %v, want %v", got, black)
	}
	if got := lerp(black, white, 1); got != white {
		t.Errorf("lerp t=1 = %v, want %v", got, white)
	}
	if mid := lerp(black, white, 0.5); mid.R < 126 || mid.R > 129 {
		t.Errorf("lerp midpoint R = %d, want ~127", mid.R)
	}
	// Out-of-range t must be clamped, so channels never wrap around uint8.
	for _, tt := range []float64{-1, 2} {
		got := lerp(black, white, tt)
		if got.R > 0xff || got.A != 0xff {
			t.Errorf("lerp t=%v produced %v (wrapped?)", tt, got)
		}
	}
}

func TestScreenNoOverflow(t *testing.T) {
	// screen brightens; with alpha in [0,1] no channel may wrap below the
	// background (a uint8 overflow would show up as a dark pixel).
	bg := color.RGBA{200, 200, 200, 0xff}
	for _, a := range []float64{0, 0.25, 0.55, 1} {
		got := screen(bg, color.RGBA{0xff, 0xff, 0xff, 0xff}, a)
		if got.R < bg.R || got.G < bg.G || got.B < bg.B {
			t.Errorf("screen(a=%v) darkened background: %v", a, got)
		}
	}
}

func TestInRoundRectCoverageRange(t *testing.T) {
	r := image.Rect(10, 10, 110, 60)
	pts := []struct{ x, y float64 }{
		{0, 0}, {60, 35}, {10.5, 10.5}, {109.5, 59.5}, {200, 200}, {10, 35},
	}
	for _, p := range pts {
		if cov := inRoundRect(p.x, p.y, r, 12); cov < 0 || cov > 1 {
			t.Errorf("inRoundRect(%v,%v) = %v, out of [0,1]", p.x, p.y, cov)
		}
	}
	if cov := inRoundRect(60, 35, r, 12); cov < 0.99 {
		t.Errorf("deep interior coverage = %v, want ~1", cov)
	}
	if cov := inRoundRect(500, 500, r, 12); cov != 0 {
		t.Errorf("far exterior coverage = %v, want 0", cov)
	}
}

func TestBlendPixelOutOfBounds(t *testing.T) {
	dst := image.NewRGBA(image.Rect(0, 0, 4, 4))
	red := color.RGBA{0xff, 0, 0, 0xff}
	// Out-of-bounds coordinates must be a no-op, not a panic.
	blendPixel(dst, -1, -1, red, 1)
	blendPixel(dst, 100, 100, red, 1)
	// In-bounds fully-opaque blend replaces the pixel.
	blendPixel(dst, 1, 1, red, 1)
	if got := dst.RGBAAt(1, 1); got.R != 0xff || got.G != 0 {
		t.Errorf("in-bounds blend = %v, want opaque red", got)
	}
	// The out-of-bounds calls left the origin untouched.
	if got := dst.RGBAAt(0, 0); got != (color.RGBA{}) {
		t.Errorf("origin modified by out-of-bounds blend: %v", got)
	}
}
