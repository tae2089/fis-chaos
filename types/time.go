package types

type TimeChaosClockID string

const (
	CLOCKREALTIME         TimeChaosClockID = "CLOCK_REALTIME"
	CLOCKMONOTONIC        TimeChaosClockID = "CLOCK_MONOTONIC"
	CLOCKPROCESSCPUTIMEID TimeChaosClockID = "CLOCK_PROCESS_CPUTIME_ID"
	CLOCKTHREADCPUTIMEID  TimeChaosClockID = "CLOCK_THREAD_CPUTIME_ID"
	CLOCKMONOTONICRAW     TimeChaosClockID = "CLOCK_MONOTONIC_RAW"
	CLOCKREALTIMECOARSE   TimeChaosClockID = "CLOCK_REALTIME_COARSE"
	CLOCKMONOTONICCOARSE  TimeChaosClockID = "CLOCK_MONOTONIC_COARSE"
	CLOCKBOOTTIME         TimeChaosClockID = "CLOCK_BOOTTIME"
	CLOCKREALTIMEALARM    TimeChaosClockID = "CLOCK_REALTIME_ALARM"
	CLOCKBOOTTIMEALARM    TimeChaosClockID = "CLOCK_BOOTTIME_ALARM"
)
