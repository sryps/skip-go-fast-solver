package metrics

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/go-kit/kit/metrics"
	prom "github.com/go-kit/kit/metrics/prometheus"
	stdprom "github.com/prometheus/client_golang/prometheus"
	math2 "math"
)

const (
	chainIDLabel              = "chain_id"
	gasBalanceLevelLabel      = "gas_balance_level"
	sourceChainIDLabel        = "source_chain_id"
	destinationChainIDLabel   = "destination_chain_id"
	successLabel              = "success"
	chainNameLabel            = "chain_name"
	sourceChainNameLabel      = "source_chain_name"
	destinationChainNameLabel = "destination_chain_name"
	chainEnvironmentLabel     = "chain_environment"
	gasTokenSymbolLabel       = "gas_token_symbol"
)

type Metrics interface {
	SetGasBalance(string, string, string, string, big.Int, big.Int, big.Int, uint8)

	AddSolverLoop()
	SolverLoopLatency(time.Duration)
	AddTransactionSubmitted(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string)
	AddTransactionRetryAttempt(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string)
	AddTransactionConfirmed(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string)

	FillLatency(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string, latency time.Duration)
}

type metricsContextKey struct{}

func ContextWithMetrics(ctx context.Context, metrics Metrics) context.Context {
	return context.WithValue(ctx, metricsContextKey{}, metrics)
}

func FromContext(ctx context.Context) Metrics {
	metricsFromContext := ctx.Value(metricsContextKey{})
	if metricsFromContext == nil {
		return NewNoOpMetrics()
	} else {
		return metricsFromContext.(Metrics)
	}
}

var _ Metrics = (*PromMetrics)(nil)

type PromMetrics struct {
	gasBalance                    metrics.Gauge
	totalSolverLoops              metrics.Counter
	latencyPerSolverLoop          metrics.Histogram
	totalTransactionSubmitted     metrics.Counter
	totalTransactionRetryAttempts metrics.Counter
	totalTransactionsConfirmed    metrics.Counter
	latencyPerFill                metrics.Histogram
}

func NewPromMetrics() Metrics {
	return &PromMetrics{
		gasBalance: prom.NewGaugeFrom(stdprom.GaugeOpts{
			Namespace: "solver",
			Name:      "gas_balance_gauge",
			Help:      "gas balances, paginated by chain id and gas balance level",
		}, []string{chainIDLabel, chainNameLabel, gasTokenSymbolLabel, chainEnvironmentLabel, gasBalanceLevelLabel}),
		totalSolverLoops: prom.NewCounterFrom(stdprom.CounterOpts{
			Namespace: "solver",
			Name:      "total_solver_loops_counter",
			Help:      "number of solver loops",
		}, []string{}),
		latencyPerSolverLoop: prom.NewHistogramFrom(stdprom.HistogramOpts{
			Namespace: "solver",
			Name:      "latency_per_solver_loop",
			Help:      "latency per solver loop in milliseconds",
			Buckets:   []float64{5, 10, 25, 50, 75, 100, 150, 200, 300, 500, 750, 1000, 1500, 3000, 5000, 10000, 20000},
		}, []string{}),
		totalTransactionSubmitted: prom.NewCounterFrom(stdprom.CounterOpts{
			Namespace: "solver",
			Name:      "total_transactions_submitted_counter",
			Help:      "number of transactions submitted, paginated by success status and source and destination chain id",
		}, []string{successLabel, sourceChainIDLabel, destinationChainIDLabel, sourceChainNameLabel, destinationChainNameLabel, chainEnvironmentLabel}),
		totalTransactionRetryAttempts: prom.NewCounterFrom(stdprom.CounterOpts{
			Namespace: "solver",
			Name:      "total_transaction_retry_attempts_counter",
			Help:      "number of transactions retried, paginated by source and destination chain id",
		}, []string{sourceChainIDLabel, destinationChainIDLabel, sourceChainNameLabel, destinationChainNameLabel, chainEnvironmentLabel}),
		totalTransactionsConfirmed: prom.NewCounterFrom(stdprom.CounterOpts{
			Namespace: "solver",
			Name:      "total_transactions_confirmed_counter",
			Help:      "number of transactions confirmed, paginated by success status and source and destination chain id",
		}, []string{successLabel, sourceChainIDLabel, destinationChainIDLabel, sourceChainNameLabel, destinationChainNameLabel, chainEnvironmentLabel}),
		latencyPerFill: prom.NewHistogramFrom(stdprom.HistogramOpts{
			Namespace: "solver",
			Name:      "latency_per_fill",
			Help:      "latency from source transaction to fill completion, paginated by source and destination chain id",
			Buckets:   []float64{30, 60, 300, 600, 900, 1200, 1500, 1800, 2400, 3000, 3600},
		}, []string{sourceChainIDLabel, destinationChainIDLabel, sourceChainNameLabel, destinationChainNameLabel, chainEnvironmentLabel}),
	}
}

func (m *PromMetrics) SetGasBalance(chainID, chainName, gasTokenSymbol, chainEnvironment string, gasBalance, warningThreshold, criticalThreshold big.Int, gasTokenDecimals uint8) {
	gasBalanceLevel := "ok"
	if gasBalance.Cmp(&warningThreshold) < 0 {
		gasBalanceLevel = "warning"
	}
	if gasBalance.Cmp(&criticalThreshold) < 0 {
		gasBalanceLevel = "critical"
	}
	// We compare the gas balance against thresholds locally rather than in the grafana alert definition since
	// the prometheus metric is exported as a float64 and the thresholds reach Wei amounts where precision is lost.
	gasBalanceFloat, _ := gasBalance.Float64()
	gasTokenAmount := gasBalanceFloat / (math2.Pow10(int(gasTokenDecimals)))
	m.gasBalance.With(chainIDLabel, chainID, chainNameLabel, chainName, gasTokenSymbolLabel, gasTokenSymbol, chainEnvironmentLabel, chainEnvironment, gasBalanceLevelLabel, gasBalanceLevel).Set(gasTokenAmount)
}

func (m *PromMetrics) AddSolverLoop() {
	m.totalSolverLoops.Add(1)
}

func (m *PromMetrics) SolverLoopLatency(latency time.Duration) {
	m.latencyPerSolverLoop.Observe(float64(latency.Milliseconds()))
}

func (m *PromMetrics) AddTransactionSubmitted(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
	m.totalTransactionSubmitted.With(successLabel, fmt.Sprint(success), sourceChainIDLabel, sourceChainID, destinationChainIDLabel, destinationChainID, sourceChainNameLabel, sourceChainName, destinationChainNameLabel, destinationChainName, chainEnvironmentLabel, chainEnvironment).Add(1)
}

func (m *PromMetrics) AddTransactionRetryAttempt(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
	m.totalTransactionRetryAttempts.With(sourceChainIDLabel, sourceChainID, destinationChainIDLabel, destinationChainID, sourceChainNameLabel, sourceChainName, destinationChainNameLabel, destinationChainName, chainEnvironmentLabel, chainEnvironment).Add(1)
}

func (m *PromMetrics) AddTransactionConfirmed(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
	m.totalTransactionsConfirmed.With(successLabel, fmt.Sprint(success), sourceChainIDLabel, sourceChainID, destinationChainIDLabel, destinationChainID, sourceChainNameLabel, sourceChainName, destinationChainNameLabel, destinationChainName, chainEnvironmentLabel, chainEnvironment).Add(1)
}

func (m *PromMetrics) FillLatency(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string, latency time.Duration) {
	m.latencyPerFill.With(sourceChainIDLabel, sourceChainID, destinationChainIDLabel, destinationChainID, sourceChainNameLabel, sourceChainName, destinationChainNameLabel, destinationChainName, chainEnvironmentLabel, chainEnvironment).Observe(float64(latency.Seconds()))
}

type NoOpMetrics struct{}

func (n NoOpMetrics) SetGasBalance(s string, s2 string, s3 string, s4 string, b big.Int, b2 big.Int, b3 big.Int, u uint8) {
}

func (n NoOpMetrics) AddSolverLoop() {}

func (n NoOpMetrics) SolverLoopLatency(duration time.Duration) {}

func (n NoOpMetrics) AddTransactionSubmitted(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
}

func (n NoOpMetrics) AddTransactionRetryAttempt(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
}

func (n NoOpMetrics) AddTransactionConfirmed(success bool, sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string) {
}

func (n NoOpMetrics) FillLatency(sourceChainID, destinationChainID, sourceChainName, destinationChainName, chainEnvironment string, latency time.Duration) {
}

func NewNoOpMetrics() Metrics {
	return &NoOpMetrics{}
}
