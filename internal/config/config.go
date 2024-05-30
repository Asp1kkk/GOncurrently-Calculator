package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AdditionDuration       time.Duration
	SubstractionDuration   time.Duration
	MultiplicationDuration time.Duration
	DivisionDuration       time.Duration

	ComputingPower int
}

func LoadFromEnv() (*Config, error) {
	conf := &Config{}

	add, err := strconv.Atoi(os.Getenv("TIME_ADDITION_MS"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as int: %w", err)
	}

	sub, err := strconv.Atoi(os.Getenv("TIME_SUBSTRACTION_MS"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as int: %w", err)
	}

	mul, err := strconv.Atoi(os.Getenv("TIME_MULTIPLICATION_MS"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as int: %w", err)
	}

	div, err := strconv.Atoi(os.Getenv("TIME_DIVISION_MS"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as int: %w", err)
	}

	cp, err := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as int: %w", err)
	}

	conf.AdditionDuration = time.Duration(add)
	conf.SubstractionDuration = time.Duration(sub)
	conf.MultiplicationDuration = time.Duration(mul)
	conf.DivisionDuration = time.Duration(div)
	conf.ComputingPower = cp

	return conf, nil
}
