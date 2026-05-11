package main

import (
	"elichika/gui/graphic"
	"elichika/gui/graphic/button"

	"fmt"
)

var gridSizes = []int{150, 200, 300}

// SetupGrid continues the navigation row started by SetupSQL, adding the grid
// toggle, page buttons, size toggle, and page label — all on the same line —
// then terminates the row with NewLine and computes GridCellsPerPage.
func (t *TextureDisplay) SetupGrid(w *graphic.Window) {
	t.gridPlaceholder = graphic.RGBATexture(0x2a2a2aff)      // dark grey: loading
	t.gridErrorPlaceholder = graphic.RGBATexture(0x5a1a1aff) // dark red:  failed

	t.gridSizeIdx = 0
	t.GridCellSize = gridSizes[t.gridSizeIdx]
	t.GridCols = t.GetWidth() / t.GridCellSize

	// --- Grid View toggle ---
	t.GridViewButton = &button.RectButton{
		Width:   200,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridViewButton.Text = graphic.NewText(t.GridViewButton, "Grid View: OFF")
	t.GridViewButton.LeftClickHandler = func() {
		t.GridView = !t.GridView
		if t.GridView {
			t.GridViewButton.Text.SetText("Grid View: ON")
			t.startLoadingGrid(w)
		} else {
			t.GridViewButton.Text.SetText("Grid View: OFF")
		}
		graphic.InvalidateRenderCache(t)
	}

	// --- Page navigation helpers ---
	goToPage := func(page int) {
		if !t.GridView {
			return
		}
		maxPage := (t.SQLLength - 1) / t.GridCellsPerPage
		if page < 0 {
			page = 0
		}
		if page > maxPage {
			page = maxPage
		}
		if page == t.GridPage {
			return
		}
		t.GridPage = page
		t.startLoadingGrid(w)
		t.updatePageLabel()
		graphic.InvalidateRenderCache(t)
	}

	// --- Prev Page ---
	t.GridPrevButton = &button.RectButton{
		Width:   120,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridPrevButton.Text = graphic.NewText(t.GridPrevButton, "< Prev")
	t.GridPrevButton.LeftClickHandler = func() { goToPage(t.GridPage - 1) }

	// --- Prev 5 Pages ---
	t.GridPrev5Button = &button.RectButton{
		Width:   120,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridPrev5Button.Text = graphic.NewText(t.GridPrev5Button, "<< -5")
	t.GridPrev5Button.LeftClickHandler = func() { goToPage(t.GridPage - 5) }

	// --- Next Page ---
	t.GridNextButton = &button.RectButton{
		Width:   120,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridNextButton.Text = graphic.NewText(t.GridNextButton, "Next >")
	t.GridNextButton.LeftClickHandler = func() { goToPage(t.GridPage + 1) }

	// --- Next 5 Pages ---
	t.GridNext5Button = &button.RectButton{
		Width:   120,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridNext5Button.Text = graphic.NewText(t.GridNext5Button, "+5 >>")
	t.GridNext5Button.LeftClickHandler = func() { goToPage(t.GridPage + 5) }

	// --- Size toggle ---
	t.GridSizeButton = &button.RectButton{
		Width:   150,
		Height:  50,
		Texture: graphic.RGBATexture(0x7f7f7fff),
	}
	t.GridSizeButton.Text = graphic.NewText(t.GridSizeButton, t.sizeLabel())
	t.GridSizeButton.LeftClickHandler = func() {
		t.gridSizeIdx = (t.gridSizeIdx + 1) % len(gridSizes)
		t.GridCellSize = gridSizes[t.gridSizeIdx]
		t.GridCols = t.GetWidth() / t.GridCellSize
		t.recomputeCellsPerPage()
		t.GridSizeButton.Text.SetText(t.sizeLabel())
		if t.GridView {
			t.startLoadingGrid(w)
		}
		graphic.InvalidateRenderCache(t)
	}

	// --- Page label ---
	t.GridPageLabel = graphic.NewText(t, "")

	t.AddObject(t.GridViewButton, 5)
	t.AddObject(t.GridPrev5Button, 5)
	t.AddObject(t.GridPrevButton, 5)
	t.AddObject(t.GridNextButton, 5)
	t.AddObject(t.GridNext5Button, 5)
	t.AddObject(t.GridSizeButton, 5)
	t.AddObject(t.GridPageLabel, 0)
	t.NewLine()

	t.recomputeCellsPerPage()
}

func (t *TextureDisplay) sizeLabel() string {
	return fmt.Sprintf("Size: %dpx", gridSizes[t.gridSizeIdx])
}

// recomputeCellsPerPage recalculates how many cells fit after the control rows.
func (t *TextureDisplay) recomputeCellsPerPage() {
	gridAreaHeight := t.GetHeight() - t.y
	gridRows := gridAreaHeight / t.GridCellSize
	t.GridCellsPerPage = gridRows * t.GridCols
}

// updatePageLabel refreshes the "Page X / Y" text.
func (t *TextureDisplay) updatePageLabel() {
	if t.SQLLength == 0 || t.GridCellsPerPage == 0 {
		t.GridPageLabel.SetText("No results")
		return
	}
	totalPages := (t.SQLLength + t.GridCellsPerPage - 1) / t.GridCellsPerPage
	t.GridPageLabel.SetText(fmt.Sprintf("Page %d / %d  (%d results)", t.GridPage+1, totalPages, t.SQLLength))
	graphic.InvalidateRenderCache(t.GridPageLabel)
}

// startLoadingGrid resets the grid textures for the current page and launches
// a goroutine to decrypt and upload them.
// Must be called on the main thread.
func (t *TextureDisplay) startLoadingGrid(w *graphic.Window) {
	t.gridGeneration++
	gen := t.gridGeneration

	pageStart := t.GridPage * t.GridCellsPerPage
	pageEnd := pageStart + t.GridCellsPerPage
	if pageEnd > t.SQLLength {
		pageEnd = t.SQLLength
	}
	pageSize := pageEnd - pageStart

	results := make([]Texture, pageSize)
	copy(results, t.SQLResult[pageStart:pageEnd])

	t.freeGridTextures()
	t.GridTextures = make([]*graphic.Texture, pageSize)
	t.GridErrors = make([]bool, pageSize)
	t.updatePageLabel()

	fmt.Printf("[grid] loading page %d: results[%d:%d] (gen %d)\n", t.GridPage, pageStart, pageEnd, gen)

	go func() {
		failed := 0
		for i := 0; i < pageSize; i++ {
			if t.gridGeneration != gen {
				return
			}
			idx := i
			file := GetStaticFile(results[idx].PackName)
			if file == "" {
				failed++
				w.InternalEvent(func() {
					if t.gridGeneration == gen {
						t.GridErrors[idx] = true
						graphic.InvalidateRenderCache(t)
					}
				})
				continue
			}
			raw := LoadUnencrypted(file, results[idx].RawTexture())
			w.InternalEvent(func() {
				if t.gridGeneration != gen {
					return
				}
				tex := &graphic.Texture{}
				tex.LoadFromMemory(raw)
				t.GridTextures[idx] = tex
				graphic.InvalidateRenderCache(t)
			})
		}
		statusChannel <- fmt.Sprintf("Page %d: %d loaded, %d failed", t.GridPage+1, pageSize-failed, failed)
	}()
}

// freeGridTextures releases all GPU textures for the current page.
// Must be called on the main thread.
func (t *TextureDisplay) freeGridTextures() {
	for _, tex := range t.GridTextures {
		if tex != nil {
			tex.Free()
		}
	}
	t.GridTextures = nil
	t.GridErrors = nil
}

// DrawGrid renders the current page's thumbnails into the canvas.
func (t *TextureDisplay) DrawGrid() {
	if len(t.GridTextures) == 0 {
		return
	}
	for i, tex := range t.GridTextures {
		col := i % t.GridCols
		row := i / t.GridCols
		x := col * t.GridCellSize
		y := t.y + row*t.GridCellSize
		if y+t.GridCellSize > t.GetHeight() {
			break
		}
		if tex != nil {
			tex.StyleType = graphic.StyleTypeFitContainer
			t.Canvas.DrawTexture(tex, x, y, t.GridCellSize, t.GridCellSize)
		} else if t.GridErrors[i] {
			t.Canvas.DrawTexture(t.gridErrorPlaceholder, x, y, t.GridCellSize, t.GridCellSize)
		} else {
			t.Canvas.DrawTexture(t.gridPlaceholder, x, y, t.GridCellSize, t.GridCellSize)
		}
	}
}

// OnClick implements graphic.Clickable. Clicking a thumbnail selects that
// texture and switches back to single-view mode.
func (t *TextureDisplay) OnClick(w *graphic.Window, e graphic.MouseButtonDownEvent) bool {
	if !t.GridView || len(t.GridTextures) == 0 {
		return false
	}
	if e.Y < t.y {
		return false
	}
	col := e.X / t.GridCellSize
	row := (e.Y - t.y) / t.GridCellSize
	localIdx := row*t.GridCols + col
	if col >= t.GridCols || localIdx < 0 || localIdx >= len(t.GridTextures) {
		return false
	}
	actualIdx := t.GridPage*t.GridCellsPerPage + localIdx

	t.GridView = false
	t.GridViewButton.Text.SetText("Grid View: OFF")
	t.SQLIndex = actualIdx

	go func() {
		texture, err := LoadTextureDirect(t.SQLResult[actualIdx].RawTexture())
		if err != nil {
			statusChannel <- fmt.Sprintf("Failed to load: %s, error: %v", t.SQLResult[actualIdx].AssetPath, err)
			return
		}
		w.InternalEvent(func() {
			if t.AssetTexture != nil {
				t.AssetTexture.Free()
			}
			t.AssetTexture = texture
			graphic.InvalidateRenderCache(t)
		})
		statusChannel <- fmt.Sprintf("%s from %s (index %d of %d)",
			t.SQLResult[actualIdx].AssetPath, t.SQLResult[actualIdx].PackName, actualIdx+1, t.SQLLength)
	}()
	return true
}
