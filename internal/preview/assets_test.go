package preview

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"image/png"
	"testing"

	"ebe-boilerplate/internal/play"
)

func TestCatalogImages(t *testing.T) {
	data, err := Files.ReadFile("assets/manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		Games []string
		Files map[string]struct {
			SHA256 string `json:"sha256"`
			Size   []int
		}
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatal(err)
	}
	if len(manifest.Games) != play.Count {
		t.Fatal("game catalog needs regenerated previews")
	}
	for id, name := range play.IDs {
		if Size*Size*Frames(id)*4 > CacheBytes {
			t.Fatalf("%s sheet exceeds the cache budget", name)
		}
		if manifest.Games[id] != name {
			t.Fatalf("stable ID mismatch for %s", name)
		}
		for style := 0; style < 3; style++ {
			for _, sheet := range []bool{false, true} {
				if style == 0 && sheet {
					continue
				}
				path := Name(id, style, sheet)
				im, err := Decode(id, style, sheet)
				if err != nil {
					t.Fatal(err)
				}
				if sheet {
					first, last := FrameRect(id, 0), FrameRect(id, Frames(id)-1)
					if !first.In(im.Bounds()) || !last.In(im.Bounds()) {
						t.Fatal("frames exceed sheet", path)
					}
					changed := false
					for frame := 1; frame < Frames(id) && !changed; frame++ {
						r := FrameRect(id, frame)
						for y := 0; y < Size && !changed; y++ {
							for x := 0; x < Size; x++ {
								a, b, c, d := im.At(x, y).RGBA()
								e, f, g, h := im.At(x+r.Min.X, y+r.Min.Y).RGBA()
								if a != e || b != f || c != g || d != h {
									changed = true
									break
								}
							}
						}
					}
					if !changed {
						t.Fatal("animation contains no gameplay motion", path)
					}
				}
				data, _ := Files.ReadFile(path)
				sum := sha256.Sum256(data)
				entry, ok := manifest.Files[path[len("assets/"):]]
				if !ok || hex.EncodeToString(sum[:]) != entry.SHA256 {
					t.Fatal("manifest checksum", path)
				}
			}
		}
	}
	if len(manifest.Files) != play.Count*5 {
		t.Fatal("incomplete or extra image variants")
	}
}

func BenchmarkDecodeSheet(b *testing.B) {
	for b.Loop() {
		f, _ := Files.Open(Name(play.TerritoryID, 1, true))
		_, err := png.Decode(f)
		f.Close()
		if err != nil {
			b.Fatal(err)
		}
	}
}
