package goblaze

import (
    "fmt"
    "strconv"
    "time"
)

type TimeView struct {
    ComponentBase
    Offset time.Duration
}

func NewTimeView() *TimeView {
    t := &TimeView{}
    t.RegisterEvents(map[string]EventHandler{
        "advance": t.Advance,
        "reset":   t.Reset,
    })
    return t
}

// Advance accepts payloads like "m5" or "s1" to add minutes or seconds.
// Examples: "m5" adds 5 minutes, "s1" adds 1 second, empty payload adds 1 minute.
func (t *TimeView) Advance(payload string) {
    if payload == "" {
        t.Offset += time.Minute
        return
    }

    // parse syntax: first char is unit (m=minutes, s=seconds), rest is number
    if len(payload) < 2 {
        return
    }

    unit := payload[0]
    numStr := payload[1:]
    n, err := strconv.Atoi(numStr)
    if err != nil {
        return
    }

    switch unit {
    case 'm':
        t.Offset += time.Duration(n) * time.Minute
    case 's':
        t.Offset += time.Duration(n) * time.Second
    }
}

// Reset clears the offset, resetting to real time.
func (t *TimeView) Reset(_ string) {
    t.Offset = 0
}

func (t *TimeView) Render() Node {
    realNow := time.Now()
    sessionNow := realNow.Add(t.Offset)

    return Div(
        H1("Server Time (per-session)"),
        Div(
            Div(Text(fmt.Sprintf("Real: %s", realNow.Format("2006-01-02 15:04:05")))),
            Div(Text(fmt.Sprintf("Session: %s", sessionNow.Format("2006-01-02 15:04:05")))),
        ),
        Div(
            Button("+5m", "onclick", "advance|m5"),
            Button("+1s", "onclick", "advance|s1"),
            Button("Reset", "onclick", "reset"),
        ),
    )
}

func (t *TimeView) Tick() {
    // no-op; Render computes time from now + offset
}
