package main

import (
	"image"
	"image/color"
	"runtime"
	"testing"
)

//func Test_getFrames(t *testing.T) {
//	testDatas := []struct {
//		imageName      string
//		numberOfFrames int
//	}{
//		{"1.gif", 25},
//		{"2.gif", 4},
//		{"3.gif", 8},
//	}
//
//	for _, td := range testDatas {
//
//		// read file
//		reader, err := os.Open("./GIF-Images/" + td.imageName)
//		if err != nil {
//			t.Errorf("failed due to error %s", err)
//		}
//		pf, err := getFrames(reader)
//
//		// check the number of frames
//		if len(pf) != td.numberOfFrames {
//			t.Fail()
//		}
//	}
//}

func Test_ClampUint8(t *testing.T) {
	var testData = []struct {
		in       int32
		expected uint8
	}{
		{0, 0},
		{255, 255},
		{128, 128},
		{-2, 0},
		{256, 255},
	}
	for _, test := range testData {
		actual := clampUint8(test.in)
		if actual != test.expected {
			t.Fail()
		}
	}
}

func Test_ClampUint16(t *testing.T) {
	var testData = []struct {
		in       int64
		expected uint16
	}{
		{0, 0},
		{65535, 65535},
		{128, 128},
		{-2, 0},
		{65536, 65535},
	}
	for _, test := range testData {
		actual := clampUint16(test.in)
		if actual != test.expected {
			t.Fail()
		}
	}
}

func Test_FloatToUint8(t *testing.T) {
	var testData = []struct {
		in       float32
		expected uint8
	}{
		{0, 0},
		{255, 255},
		{128, 128},
		{1, 1},
		{256, 255},
	}
	for _, test := range testData {
		actual := floatToUint8(test.in)
		if actual != test.expected {
			t.Fail()
		}
	}
}

func Test_FloatToUint16(t *testing.T) {
	var testData = []struct {
		in       float32
		expected uint16
	}{
		{0, 0},
		{65535, 65535},
		{128, 128},
		{1, 1},
		{65536, 65535},
	}
	for _, test := range testData {
		actual := floatToUint16(test.in)
		if actual != test.expected {
			t.Fail()
		}
	}
}

var img = image.NewGray16(image.Rect(0, 0, 3, 3))

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	img.Set(1, 1, color.White)
}

func Test_Param1(t *testing.T) {
	m := Resize(0, 0, img, NEAREST_NEIHHBOR)
	if m.Bounds() != img.Bounds() {
		t.Fail()
	}
}

func Test_Param2(t *testing.T) {
	m := Resize(100, 0, img, NEAREST_NEIHHBOR)
	if m.Bounds() != image.Rect(0, 0, 100, 100) {
		t.Fail()
	}
}

func Test_ZeroImg(t *testing.T) {
	zeroImg := image.NewGray16(image.Rect(0, 0, 0, 0))

	m := Resize(0, 0, zeroImg, NEAREST_NEIHHBOR)
	if m.Bounds() != zeroImg.Bounds() {
		t.Fail()
	}
}

func Test_HalfZeroImg(t *testing.T) {
	zeroImg := image.NewGray16(image.Rect(0, 0, 0, 100))

	m := Resize(0, 1, zeroImg, NEAREST_NEIHHBOR)
	if m.Bounds() != zeroImg.Bounds() {
		t.Fail()
	}

	m = Resize(1, 0, zeroImg, NEAREST_NEIHHBOR)
	if m.Bounds() != zeroImg.Bounds() {
		t.Fail()
	}
}

func Test_CorrectResize(t *testing.T) {
	zeroImg := image.NewGray16(image.Rect(0, 0, 256, 256))

	m := Resize(60, 0, zeroImg, NEAREST_NEIHHBOR)
	if m.Bounds() != image.Rect(0, 0, 60, 60) {
		t.Fail()
	}
}

func Test_SameColorWithRGBA(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetRGBA(x, y, color.RGBA{0x80, 0x80, 0x80, 0xFF})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.RGBA)
			if color.R != 0x80 || color.G != 0x80 || color.B != 0x80 || color.A != 0xFF {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_SameColorWithNRGBA(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetNRGBA(x, y, color.NRGBA{0x80, 0x80, 0x80, 0xFF})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.RGBA)
			if color.R != 0x80 || color.G != 0x80 || color.B != 0x80 || color.A != 0xFF {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_SameColorWithRGBA64(t *testing.T) {
	img := image.NewRGBA64(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetRGBA64(x, y, color.RGBA64{0x8000, 0x8000, 0x8000, 0xFFFF})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.RGBA64)
			if color.R != 0x8000 || color.G != 0x8000 || color.B != 0x8000 || color.A != 0xFFFF {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_SameColorWithNRGBA64(t *testing.T) {
	img := image.NewNRGBA64(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetNRGBA64(x, y, color.NRGBA64{0x8000, 0x8000, 0x8000, 0xFFFF})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.RGBA64)
			if color.R != 0x8000 || color.G != 0x8000 || color.B != 0x8000 || color.A != 0xFFFF {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_SameColorWithGray(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetGray(x, y, color.Gray{0x80})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.Gray)
			if color.Y != 0x80 {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_SameColorWithGray16(t *testing.T) {
	img := image.NewGray16(image.Rect(0, 0, 20, 20))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.SetGray16(x, y, color.Gray16{0x8000})
		}
	}
	out := Resize(10, 10, img, LANCZOS_3)
	for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
		for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
			color := out.At(x, y).(color.Gray16)
			if color.Y != 0x8000 {
				t.Errorf("%+v", color)
			}
		}
	}
}

func Test_Bounds(t *testing.T) {
	img := image.NewRGBA(image.Rect(20, 10, 200, 99))
	out := Resize(80, 80, img, LANCZOS_2)
	out.At(0, 0)
}

func Test_SameSizeReturnsOriginal(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	out := Resize(0, 0, img, LANCZOS_2)

	if img != out {
		t.Fail()
	}

	out = Resize(10, 10, img, LANCZOS_2)

	if img != out {
		t.Fail()
	}
}

func Test_PixelCoordinates(t *testing.T) {
	checkers := image.NewGray(image.Rect(0, 0, 4, 4))
	checkers.Pix = []uint8{
		255, 0, 255, 0,
		0, 255, 0, 255,
		255, 0, 255, 0,
		0, 255, 0, 255,
	}

	resized := Resize(12, 12, checkers, NEAREST_NEIHHBOR).(*image.Gray)

	if resized.Pix[0] != 255 || resized.Pix[1] != 255 || resized.Pix[2] != 255 {
		t.Fail()
	}

	if resized.Pix[3] != 0 || resized.Pix[4] != 0 || resized.Pix[5] != 0 {
		t.Fail()
	}
}

func Test_ResizeWithPremultipliedAlpha(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 4))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		// 0x80 = 0.5 * 0xFF.
		img.SetRGBA(0, y, color.RGBA{0x80, 0x80, 0x80, 0x80})
	}

	out := Resize(1, 2, img, METHCEL_NETRAVALI)

	outputColor := out.At(0, 0).(color.RGBA)
	if outputColor.R != 0x80 {
		t.Fail()
	}
}

func Test_ResizeWithTranslucentColor(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 2))

	// Set the pixel colors to an "invisible green" and white.
	// After resizing, the green shouldn't be visible.
	img.SetNRGBA(0, 0, color.NRGBA{0x00, 0xFF, 0x00, 0x00})
	img.SetNRGBA(0, 1, color.NRGBA{0x00, 0x00, 0x00, 0xFF})

	out := Resize(1, 1, img, BILINEARB)

	_, g, _, _ := out.At(0, 0).RGBA()
	if g != 0x00 {
		t.Errorf("%+v", g)
	}
}

const (
	// Use a small image size for benchmarks. We don't want memory performance
	// to affect the benchmark results.
	benchMaxX = 250
	benchMaxY = 250

	// Resize values near the original size require increase the amount of time
	// resize spends converting the image.
	benchWidth  = 200
	benchHeight = 200
)

func benchRGBA(b *testing.B, interp InterpolationFunction) {
	m := image.NewRGBA(image.Rect(0, 0, benchMaxX, benchMaxY))
	// Initialize m's pixels to create a non-uniform image.
	for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
		for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
			i := m.PixOffset(x, y)
			m.Pix[i+0] = uint8(y + 4*x)
			m.Pix[i+1] = uint8(y + 4*x)
			m.Pix[i+2] = uint8(y + 4*x)
			m.Pix[i+3] = uint8(4*y + x)
		}
	}

	var out image.Image
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Resize(benchWidth, benchHeight, m, interp)
	}
	out.At(0, 0)
}

// The names of some interpolation functions are truncated so that the columns
// of 'go test -bench' line up.
func Benchmark_Nearest_RGBA(b *testing.B) {
	benchRGBA(b, NEAREST_NEIHHBOR)
}

func Benchmark_Bilinear_RGBA(b *testing.B) {
	benchRGBA(b, BILINEARB)
}

func Benchmark_Bicubic_RGBA(b *testing.B) {
	benchRGBA(b, BICUBIC)
}

func Benchmark_Mitchell_RGBA(b *testing.B) {
	benchRGBA(b, METHCEL_NETRAVALI)
}

func Benchmark_Lanczos2_RGBA(b *testing.B) {
	benchRGBA(b, LANCZOS_2)
}

func Benchmark_Lanczos3_RGBA(b *testing.B) {
	benchRGBA(b, LANCZOS_3)
}

func benchYCbCr(b *testing.B, interp InterpolationFunction) {
	m := image.NewYCbCr(image.Rect(0, 0, benchMaxX, benchMaxY), image.YCbCrSubsampleRatio422)
	// Initialize m's pixels to create a non-uniform image.
	for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
		for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
			yi := m.YOffset(x, y)
			ci := m.COffset(x, y)
			m.Y[yi] = uint8(16*y + x)
			m.Cb[ci] = uint8(y + 16*x)
			m.Cr[ci] = uint8(y + 16*x)
		}
	}
	var out image.Image
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Resize(benchWidth, benchHeight, m, interp)
	}
	out.At(0, 0)
}

func Benchmark_Nearest_YCC(b *testing.B) {
	benchYCbCr(b, NEAREST_NEIHHBOR)
}

func Benchmark_Bilinear_YCC(b *testing.B) {
	benchYCbCr(b, BILINEARB)
}

func Benchmark_Bicubic_YCC(b *testing.B) {
	benchYCbCr(b, BICUBIC)
}

func Benchmark_Mitchell_YCC(b *testing.B) {
	benchYCbCr(b, METHCEL_NETRAVALI)
}

func Benchmark_Lanczos2_YCC(b *testing.B) {
	benchYCbCr(b, LANCZOS_2)
}

func Benchmark_Lanczos3_YCC(b *testing.B) {
	benchYCbCr(b, LANCZOS_3)
}

type Image interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func TestImage(t *testing.T) {
	testImage := []Image{
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio420),
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio422),
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio440),
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio411),
		newYCC(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio410),
	}
	for _, m := range testImage {
		if !image.Rect(0, 0, 10, 10).Eq(m.Bounds()) {
			t.Errorf("%T: want bounds %v, got %v",
				m, image.Rect(0, 0, 10, 10), m.Bounds())
			continue
		}
		m = m.SubImage(image.Rect(3, 2, 9, 8)).(Image)
		if !image.Rect(3, 2, 9, 8).Eq(m.Bounds()) {
			t.Errorf("%T: sub-image want bounds %v, got %v",
				m, image.Rect(3, 2, 9, 8), m.Bounds())
			continue
		}
		// Test that taking an empty sub-image starting at a corner does not panic.
		m.SubImage(image.Rect(0, 0, 0, 0))
		m.SubImage(image.Rect(10, 0, 10, 0))
		m.SubImage(image.Rect(0, 10, 0, 10))
		m.SubImage(image.Rect(10, 10, 10, 10))
	}
}

func TestConvertYCbCr(t *testing.T) {
	testImage := []Image{
		image.NewYCbCr(image.Rect(0, 0, 50, 50), image.YCbCrSubsampleRatio420),
		image.NewYCbCr(image.Rect(0, 0, 50, 50), image.YCbCrSubsampleRatio422),
		image.NewYCbCr(image.Rect(0, 0, 50, 50), image.YCbCrSubsampleRatio440),
		image.NewYCbCr(image.Rect(0, 0, 50, 50), image.YCbCrSubsampleRatio444),
		image.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio411),
		image.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio410),
	}
	for _, img := range testImage {
		m := img.(*image.YCbCr)
		for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
			for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
				yi := m.YOffset(x, y)
				ci := m.COffset(x, y)
				m.Y[yi] = uint8(16*y + x)
				m.Cb[ci] = uint8(y + 16*x)
				m.Cr[ci] = uint8(y + 16*x)
			}
		}

		// test conversion from YCbCr to ycc
		yc := imageYCbCrToYCC(m)
		for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
			for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
				ystride := 3 * (m.Rect.Max.X - m.Rect.Min.X)
				xstride := 3
				yi := m.YOffset(x, y)
				ci := m.COffset(x, y)
				si := (y * ystride) + (x * xstride)
				if m.Y[yi] != yc.Pix[si] {
					t.Errorf("Err Y - found: %d expected: %d x: %d y: %d yi: %d si: %d",
						m.Y[yi], yc.Pix[si], x, y, yi, si)
				}
				if m.Cb[ci] != yc.Pix[si+1] {
					t.Errorf("Err Cb - found: %d expected: %d x: %d y: %d ci: %d si: %d",
						m.Cb[ci], yc.Pix[si+1], x, y, ci, si+1)
				}
				if m.Cr[ci] != yc.Pix[si+2] {
					t.Errorf("Err Cr - found: %d expected: %d x: %d y: %d ci: %d si: %d",
						m.Cr[ci], yc.Pix[si+2], x, y, ci, si+2)
				}
			}
		}

		// test conversion from ycc back to YCbCr
		ym := yc.YCbCr()
		for y := m.Rect.Min.Y; y < m.Rect.Max.Y; y++ {
			for x := m.Rect.Min.X; x < m.Rect.Max.X; x++ {
				yi := m.YOffset(x, y)
				ci := m.COffset(x, y)
				if m.Y[yi] != ym.Y[yi] {
					t.Errorf("Err Y - found: %d expected: %d x: %d y: %d yi: %d",
						m.Y[yi], ym.Y[yi], x, y, yi)
				}
				if m.Cb[ci] != ym.Cb[ci] {
					t.Errorf("Err Cb - found: %d expected: %d x: %d y: %d ci: %d",
						m.Cb[ci], ym.Cb[ci], x, y, ci)
				}
				if m.Cr[ci] != ym.Cr[ci] {
					t.Errorf("Err Cr - found: %d expected: %d x: %d y: %d ci: %d",
						m.Cr[ci], ym.Cr[ci], x, y, ci)
				}
			}
		}
	}
}

func TestYCbCr(t *testing.T) {
	rects := []image.Rectangle{
		image.Rect(0, 0, 16, 16),
		image.Rect(1, 0, 16, 16),
		image.Rect(0, 1, 16, 16),
		image.Rect(1, 1, 16, 16),
		image.Rect(1, 1, 15, 16),
		image.Rect(1, 1, 16, 15),
		image.Rect(1, 1, 15, 15),
		image.Rect(2, 3, 14, 15),
		image.Rect(7, 0, 7, 16),
		image.Rect(0, 8, 16, 8),
		image.Rect(0, 0, 10, 11),
		image.Rect(5, 6, 16, 16),
		image.Rect(7, 7, 8, 8),
		image.Rect(7, 8, 8, 9),
		image.Rect(8, 7, 9, 8),
		image.Rect(8, 8, 9, 9),
		image.Rect(7, 7, 17, 17),
		image.Rect(8, 8, 17, 17),
		image.Rect(9, 9, 17, 17),
		image.Rect(10, 10, 17, 17),
	}
	subsampleRatios := []image.YCbCrSubsampleRatio{
		image.YCbCrSubsampleRatio444,
		image.YCbCrSubsampleRatio422,
		image.YCbCrSubsampleRatio420,
		image.YCbCrSubsampleRatio440,
	}
	deltas := []image.Point{
		image.Pt(0, 0),
		image.Pt(1000, 1001),
		image.Pt(5001, -400),
		image.Pt(-701, -801),
	}
	for _, r := range rects {
		for _, subsampleRatio := range subsampleRatios {
			for _, delta := range deltas {
				testYCbCr(t, r, subsampleRatio, delta)
			}
		}
		if testing.Short() {
			break
		}
	}
}

func testYCbCr(t *testing.T, r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio, delta image.Point) {
	// Create a YCbCr image m, whose bounds are r translated by (delta.X, delta.Y).
	r1 := r.Add(delta)
	img := image.NewYCbCr(r1, subsampleRatio)

	// Initialize img's pixels. For 422 and 420 subsampling, some of the Cb and Cr elements
	// will be set multiple times. That's OK. We just want to avoid a uniform image.
	for y := r1.Min.Y; y < r1.Max.Y; y++ {
		for x := r1.Min.X; x < r1.Max.X; x++ {
			yi := img.YOffset(x, y)
			ci := img.COffset(x, y)
			img.Y[yi] = uint8(16*y + x)
			img.Cb[ci] = uint8(y + 16*x)
			img.Cr[ci] = uint8(y + 16*x)
		}
	}

	m := imageYCbCrToYCC(img)

	// Make various sub-images of m.
	for y0 := delta.Y + 3; y0 < delta.Y+7; y0++ {
		for y1 := delta.Y + 8; y1 < delta.Y+13; y1++ {
			for x0 := delta.X + 3; x0 < delta.X+7; x0++ {
				for x1 := delta.X + 8; x1 < delta.X+13; x1++ {
					subRect := image.Rect(x0, y0, x1, y1)
					sub := m.SubImage(subRect).(*ycc)

					// For each point in the sub-image's bounds, check that m.At(x, y) equals sub.At(x, y).
					for y := sub.Rect.Min.Y; y < sub.Rect.Max.Y; y++ {
						for x := sub.Rect.Min.X; x < sub.Rect.Max.X; x++ {
							color0 := m.At(x, y).(color.YCbCr)
							color1 := sub.At(x, y).(color.YCbCr)
							if color0 != color1 {
								t.Errorf("r=%v, subsampleRatio=%v, delta=%v, x=%d, y=%d, color0=%v, color1=%v",
									r, subsampleRatio, delta, x, y, color0, color1)
								return
							}
						}
					}
				}
			}
		}
	}
}
