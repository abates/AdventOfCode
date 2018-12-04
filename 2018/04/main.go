package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type EventType int

const (
	START EventType = iota
	ASLEEP
	AWAKE
)

type Event struct {
	ID   int
	Type EventType
}

func (e *Event) UnmarshalText(text []byte) error {
	str := strings.TrimSpace(string(text))
	if strings.HasPrefix(str, "Guard") {
		e.Type = START
		fmt.Sscanf(str, "Guard #%d begins shift", &e.ID)
	} else if strings.HasPrefix(str, "falls") {
		e.Type = ASLEEP
	} else if strings.HasPrefix(str, "wakes") {
		e.Type = AWAKE
	} else {
		return fmt.Errorf("Unknown event string %q", str)
	}
	return nil
}

func (e Event) String() string {
	switch e.Type {
	case START:
		return fmt.Sprintf("Guard #%d begins shift", e.ID)
	case ASLEEP:
		return "falls asleep"
	case AWAKE:
		return "wakes up"
	}
	return ""
}

type LogEntry struct {
	Time  time.Time
	Event Event
}

func (le *LogEntry) Less(other *LogEntry) bool {
	return le.Time.Before(other.Time)
}

func (le *LogEntry) UnmarshalText(text []byte) (err error) {
	if bytes.HasPrefix(text, []byte("[")) {
		i := bytes.Index(text, []byte("]"))
		dateString := text[0 : i+1]
		eventString := text[i+1:]
		le.Time, err = time.Parse("[2006-01-02 15:04]", string(dateString))
		if err == nil {
			err = le.Event.UnmarshalText(eventString)
		}
	} else {
		err = fmt.Errorf("Cannot parse %q", text)
	}
	return err
}

func (le *LogEntry) String() string {
	return fmt.Sprintf("%s %s", le.Time.Format("[2006-01-02 15:04]"), le.Event)
}

type Log struct {
	entries []*LogEntry
}

func (l *Log) Push(text []byte) error {
	le := &LogEntry{}
	err := le.UnmarshalText(text)
	if err == nil {
		l.entries = append(l.entries, le)
	}
	return err
}

func (l *Log) Sort() {
	sort.Sort(l)
}

func (l *Log) Len() int {
	return len(l.entries)
}

func (l *Log) Less(i, j int) bool {
	return l.entries[i].Less(l.entries[j])
}

func (l *Log) Swap(i, j int) {
	tmp := l.entries[i]
	l.entries[i] = l.entries[j]
	l.entries[j] = tmp
}

type Shift struct {
	sleepMinutes []int
}

func (s *Shift) AddSleep(start, end time.Time) {
	ssec := start.Minute()
	esec := end.Minute()
	for i := ssec; i < esec; i++ {
		s.sleepMinutes = append(s.sleepMinutes, i)
	}
}

func (s *Shift) TotalSleep() int {
	return len(s.sleepMinutes)
}

type Guard struct {
	shifts []*Shift
}

func (g *Guard) NewShift(time.Time) {
	g.shifts = append(g.shifts, &Shift{})
}

func (g *Guard) AddSleep(start, end time.Time) {
	g.shifts[len(g.shifts)-1].AddSleep(start, end)
}

func (g *Guard) TotalSleep() int {
	total := 0
	for _, shift := range g.shifts {
		total = shift.TotalSleep()
	}
	return total
}

func (g *Guard) MostSleep() (int, int) {
	max := -1
	maxTime := -1
	times := make(map[int]int)
	for _, shift := range g.shifts {
		for _, minute := range shift.sleepMinutes {
			times[minute]++
			if max == -1 || max < times[minute] {
				max = times[minute]
				maxTime = minute
			}
		}
	}
	return maxTime, max
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")

	log := &Log{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		err := log.Push(line)
		if err != nil {
			panic(err.Error())
		}
	}
	log.Sort()

	guards := make(map[int]*Guard)
	var currentGuard *Guard
	var sleepStart time.Time
	for _, entry := range log.entries {
		switch entry.Event.Type {
		case START:
			if _, found := guards[entry.Event.ID]; !found {
				guards[entry.Event.ID] = &Guard{}
			}
			currentGuard = guards[entry.Event.ID]
			currentGuard.NewShift(entry.Time)
		case ASLEEP:
			sleepStart = entry.Time
		case AWAKE:
			currentGuard.AddSleep(sleepStart, entry.Time)
		}
	}

	/* Part 1 */
	max := -1
	maxId := -1
	for id, guard := range guards {
		total := guard.TotalSleep()
		if max < total || max == -1 {
			max = total
			maxId = id
		}
	}

	commonMinute, _ := guards[maxId].MostSleep()

	fmt.Printf("Guard with most sleep: %d most common sleep minute: %d -- %d\n", maxId, commonMinute, maxId*commonMinute)

	/* Part 2 */
	max = -1
	maxId = -1
	maxMinute := -1
	for id, guard := range guards {
		minute, sleep := guard.MostSleep()
		if max < sleep || max == -1 {
			max = sleep
			maxMinute = minute
			maxId = id
		}
	}
	fmt.Printf("Guard with most common minute: %d most common sleep minute: %d -- %d\n", maxId, maxMinute, maxId*maxMinute)
}
