package utils

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func File(c echo.Context, fileField string, maxBytes int64) ([]byte, error) {
	fileHeader, err := c.FormFile(fileField)
	if err != nil {
		return nil, fmt.Errorf("%s is required: %w", fileField, err)
	}

	if fileHeader.Size > maxBytes {
		return nil, fmt.Errorf("file exceeds %d bytes limit", maxBytes)
	}

	source, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer func() {
		_ = source.Close()
	}()

	raw, err := io.ReadAll(io.LimitReader(source, maxBytes))
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	return raw, nil
}

