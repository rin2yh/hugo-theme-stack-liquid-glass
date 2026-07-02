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
	cases := []struct {
		name      string
		v, lo, hi float64
		want      float64
	}{
		{"below", -1, 0, 1, 0},
		{"above", 2, 0, 1, 1},
		{"inside", 0.5, 0, 1, 0.5},
		{"at-low", 0, 0, 1, 0},
		{"at-high", 1, 0, 1, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := clamp(c.v, c.lo, c.hi); got != c.want {
				t.Errorf("clamp(%v,%v,%v)=%v, want %v", c.v, c.lo, c.hi, got, c.want)
			}
		})
	}
}

func TestSmoothstep(t *testing.T) {
	// All cases interpolate over [0,1]; the result is smooth and clamped.
	cases := []struct {
		name string
		x    float64
		want float64
	}{
		{"low-edge", 0, 0},
		{"high-edge", 1, 1},
		{"midpoint", 0.5, 0.5},
		{"quarter", 0.25, 0.15625},
		{"below-clamped", -5, 0},
		{"above-clamped", 5, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := smoothstep(0, 1, c.x); math.Abs(got-c.want) > 1e-9 {
				t.Errorf("smoothstep(0,1,%v)=%v, want %v", c.x, got, c.want)
			}
		})
	}
}

func TestLerp(t *testing.T) {
	// Interpolating black->white; out-of-range t is clamped, so channels
	// never wrap around uint8.
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	cases := []struct {
		name string
		t    float64
		want color.RGBA
	}{
		{"at-0", 0, black},
		{"at-1", 1, white},
		{"midpoint", 0.5, color.RGBA{127, 127, 127, 0xff}},
		{"below-clamped", -1, black},
		{"above-clamped", 2, white},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := lerp(black, white, c.t); got != c.want {
				t.Errorf("lerp(black,white,%v)=%v, want %v", c.t, got, c.want)
			}
		})
	}
}

func TestScreenNoOverflow(t *testing.T) {
	// screen brightens; with alpha in [0,1] no channel may wrap below the
	// background (a uint8 overflow would show up as a dark pixel).
	bg := color.RGBA{200, 200, 200, 0xff}
	cases := []struct {
		name string
		a    float64
	}{
		{"zero", 0},
		{"quarter", 0.25},
		{"blob-max", 0.55},
		{"full", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := screen(bg, color.RGBA{0xff, 0xff, 0xff, 0xff}, c.a)
			if got.R < bg.R || got.G < bg.G || got.B < bg.B {
				t.Errorf("screen(a=%v) darkened background: %v", c.a, got)
			}
		})
	}
}

func TestInRoundRectCoverageRange(t *testing.T) {
	// Coverage must always land in [lo,hi]; interior is ~1, far exterior 0,
	// and every point stays within the anti-aliased [0,1] range.
	r := image.Rect(10, 10, 110, 60)
	cases := []struct {
		name   string
		x, y   float64
		lo, hi float64
	}{
		{"interior", 60, 35, 0.99, 1},
		{"far-exterior", 500, 500, 0, 0},
		{"corner-outside", 0, 0, 0, 1},
		{"near-min-corner", 10.5, 10.5, 0, 1},
		{"near-max-corner", 109.5, 59.5, 0, 1},
		{"left-edge", 10, 35, 0, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if cov := inRoundRect(c.x, c.y, r, 12); cov < c.lo || cov > c.hi {
				t.Errorf("inRoundRect(%v,%v)=%v, want in [%v,%v]", c.x, c.y, cov, c.lo, c.hi)
			}
		})
	}
}

func TestBlendPixelOutOfBounds(t *testing.T) {
	dst := image.NewRGBA(image.Rect(0, 0, 4, 4))
	red := color.RGBA{0xff, 0, 0, 0xff}
	// Every out-of-bounds coordinate must be a no-op, not a panic.
	oob := []struct {
		name string
		x, y int
	}{
		{"both-negative", -1, -1},
		{"both-beyond", 100, 100},
		{"x-negative", -1, 2},
		{"y-beyond", 2, 100},
	}
	for _, c := range oob {
		t.Run(c.name, func(t *testing.T) {
			blendPixel(dst, c.x, c.y, red, 1)
		})
	}
	if got := dst.RGBAAt(0, 0); got != (color.RGBA{}) {
		t.Errorf("out-of-bounds blend modified in-bounds pixel: %v", got)
	}
	// A fully-opaque in-bounds blend replaces the target pixel.
	blendPixel(dst, 1, 1, red, 1)
	if got := dst.RGBAAt(1, 1); got != red {
		t.Errorf("in-bounds blend = %v, want %v", got, red)
	}
}
