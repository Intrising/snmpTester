package task

import "fmt"

// Stats is the Statistics
type Stats struct {
	total       int
	pass        int
	failed      int
	marked      int
	unmarkedOID []string
}

var stats *Stats

// AddunmarkedOID asdf
func (s *Stats) AddunmarkedOID(oid string) {
	s.unmarkedOID = append(s.unmarkedOID, oid)
}

// AddPass asdf
func (s *Stats) AddPass() {
	s.pass++
	s.total++
}

// AddFailed asdf
func (s *Stats) AddFailed() {
	s.failed++
	s.total++
}

// AddMarked asdfa
func (s *Stats) AddMarked() {
	s.marked++
}

// PrintStats asdf
func PrintStats() {
	fmt.Println("\n\n=================== Stats " + deviceIP + "=============================")
	fmt.Println("\n\n=================== Failed OIDs =============================")
	if stats.marked != stats.failed {
		for _, val := range stats.unmarkedOID {
			fmt.Println(val)
		}
	}
	fmt.Println("Pass:        ", stats.pass)
	fmt.Println("Failed:      ", stats.failed)
	fmt.Println("Fail marked: ", stats.marked)
	fmt.Println("Total:      ", stats.total)
}

func init() {
	stats = new(Stats)
	stats.total = 0
	stats.pass = 0
	stats.failed = 0
	stats.marked = 0
}
