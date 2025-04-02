package processor

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

// LockProvider defines a distributed lock provider.
type LockProvider interface {
	Acquire(ctx context.Context, key string) (bool, func() error, error)
}

// Processor defines a movie processor.
type Processor struct {
	logger       *zap.Logger
	lockProvider LockProvider
}

// New creates a new movie processor.
func New(logger *zap.Logger, lockProvider LockProvider) *Processor {
	return &Processor{logger, lockProvider}
}

// Start starts the movie processor.
func (p *Processor) Start(ctx context.Context) error {
	p.logger.Info("Starting the movie processor")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		_, release, err := p.lockProvider.Acquire(ctx, "locks/service/movie/processor")
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			p.logger.Error("Unable to acquire lock, retrying in 10 seconds", zap.Error(err))
			time.Sleep(time.Second * 10)
			continue
		}
		p.logger.Info("Lock has been acquired, starting processing")
		if err := p.process(ctx); err != nil {
			p.logger.Error("Processing error", zap.Error(err))
		} else {
			p.logger.Info("Processing completed successfully")
		}
		p.logger.Info("Releasing the lock")
		if err := release(); err != nil {
			p.logger.Error("Failed to release the lock", zap.Error(err))
		}
	}
}

func (p *Processor) process(_ context.Context) error {
	// Here we would process the movie videos.
	// We omit the implementation of movie processing logic for simplicity and just use time.Sleep to simulate processing.
	// Example of some real processing logic would be:
	// return exec.Command("bash", "-c", `for file in *.mp4; do ffmpeg -i "$file" "${file%.mp4}.mp3" && rm file; done`).Run()
	time.Sleep(time.Second * 10)
	return nil
}
