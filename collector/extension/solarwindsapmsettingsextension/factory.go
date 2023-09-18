package solarwindsapmsettingsextension

import (
	"context"
	"github.com/open-telemetry/opentelemetry-lambda/collector/extension/solarwindsapmsettingsextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	DefaultEndpoint = "apm.collector.cloud.solarwinds.com:443"
	DefaultInterval = "1m"
)

func createDefaultConfig() component.Config {
	return &Config{
		Endpoint: DefaultEndpoint,
		Interval: DefaultInterval,
	}
}

func createExtension(ctx context.Context, settings extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	logger := settings.Logger
	extensionCfg := cfg.(*Config)

	extension := &solarwindsapmSettingsExtension{
		logger: logger,
		config: extensionCfg,
	}
	return extension, nil
}

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}
