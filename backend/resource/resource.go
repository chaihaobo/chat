package resource

import (
	"context"

	"github.com/chaihaobo/chat/resource/config"
	"github.com/chaihaobo/chat/resource/logger"
	"github.com/chaihaobo/chat/resource/metric"
	"github.com/chaihaobo/chat/resource/tracer"
	"github.com/chaihaobo/chat/resource/validator"
)

type (
	Resource interface {
		Configuration() *config.Configuration
		Logger() logger.Logger
		Validator() validator.Validator
		Metric() metric.PrometheusMetric
		Tracer() tracer.Tracer
		Close() error
	}

	resource struct {
		configuration *config.Configuration
		logger        logger.Logger
		validator     validator.Validator
		metric        metric.PrometheusMetric
		tracer        tracer.Tracer
		loggerFlusher func() error
	}
)

func (r *resource) Tracer() tracer.Tracer {
	return r.tracer
}

func (r *resource) Metric() metric.PrometheusMetric {
	return r.metric
}

func (r *resource) Validator() validator.Validator {
	return r.validator
}

func (r *resource) Logger() logger.Logger {
	return r.logger
}

func (r *resource) Close() error {
	ctx := context.Background()
	closeFuncs := []func() error{
		r.loggerFlusher,
		func() error {
			return r.metric.Close(ctx)
		},
		func() error {
			return r.tracer.Close(ctx)
		},
	}
	for _, closeFun := range closeFuncs {
		if err := closeFun(); err != nil {
			return err
		}
	}

	return nil
}

func (r *resource) Configuration() *config.Configuration {
	return r.configuration
}

func New(configPath string) (Resource, error) {
	configuration, err := config.NewConfiguration(configPath)
	if err != nil {
		return nil, err
	}

	logConfig := configuration.Logger
	logger, f, err := logger.New(logger.Config{
		FileName:   logConfig.FileName,
		MaxSize:    logConfig.MaxSize,
		MaxAge:     logConfig.MaxAge,
		WithCaller: true,
		CallerSkip: 1,
	})
	if err != nil {
		return nil, err
	}

	validator, err := validator.NewValidator()
	if err != nil {
		return nil, err

	}

	prometheusMetric, err := metric.NewPrometheusMetric(configuration)
	if err != nil {
		return nil, err
	}

	tracer, err := tracer.NewTracer(configuration)
	if err != nil {
		return nil, err
	}

	return &resource{
		configuration: configuration,
		logger:        logger,
		validator:     validator,
		loggerFlusher: f,
		metric:        prometheusMetric,
		tracer:        tracer,
	}, nil
}
