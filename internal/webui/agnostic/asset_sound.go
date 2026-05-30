package agnostic

import (
	"fmt"
	"os"
	"path/filepath"

	"elichika/internal/assetdata"
	"elichika/internal/config"

	"github.com/eman1can/sound_decrypt/awb"
	"github.com/eman1can/sound_decrypt/wav"
)

func loadPackBytes(packName string) ([]byte, error) {
	downloadData := assetdata.GetDownloadData(packName)

	data, err := os.ReadFile(filepath.Join(config.StaticDataPath, "packs", downloadData.File))
	if err != nil {
		return []byte{}, err
	}

	offset := downloadData.Start
	size := downloadData.Size

	return data[offset : offset+size], nil
}

// ConvertVoiceToWAV finds the AWB pack for sheetName, extracts and fixes the
// HCA, converts it to WAV via ffmpeg, and caches the result in static/sounds/.
// Returns the WAV path, or an error when the sound is unavailable.
func ConvertVoiceToWAV(sheetName string) (string, error) {
	path := filepath.Join(config.StaticDataPath, "sounds", "wav", sheetName+".wav")
	if _, err := os.Stat(path); err == nil {
		return path, nil // already cached
	}

	sound, ok := assetdata.SoundBySheetName[sheetName]
	if !ok {
		return "", fmt.Errorf("sound not found: %s", sheetName)
	}

	awbData, err := loadPackBytes(sound.AwbPackName)
	if err != nil {
		return "", err
	}

	awbFile, err := awb.LoadAWB(awbData, config.AssetAWBKey)
	if err != nil {
		return "", err
	}

	for _, hcaFile := range awbFile.Subfiles {
		f, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating %s: %v\n", path, err)
			return "", err
		}

		if err := wav.WriteWAV(hcaFile, f); err != nil {
			fmt.Fprintf(os.Stderr, "error writing %s: %v\n", path, err)
			f.Close()
			return "", err
		}

		f.Close()
		return path, nil
	}

	return "", fmt.Errorf("no such wav file: %s", path)
}
