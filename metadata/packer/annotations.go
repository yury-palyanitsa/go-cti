package packer

import (
	"fmt"
	"log/slog"

	"github.com/acronis/go-cti/metadata"
	"github.com/acronis/go-cti/metadata/archiver"
)

func defaultAnnotationHandler(
	baseDir string, archiver archiver.Archiver, key metadata.GJsonPath, entity *metadata.Entity, a metadata.Annotations,
) error {
	// process asset annotation
	if a.Asset != nil {
		value := key.GetValue(entity.Values)
		assetPath := value.String()
		if assetPath == "" {
			slog.Warn("Empty asset path", slog.String("entity", entity.Cti), slog.String("key", value.Str))
			return nil
		}
		if err := archiver.WriteFile(baseDir, assetPath); err != nil {
			return fmt.Errorf("write asset %s: %w", assetPath, err)
		}
	}
	return nil
}
