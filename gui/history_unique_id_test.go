package main

import "testing"

func TestUniqueHistoryID(t *testing.T) {
	h := []JobHistory{{ID: "1764280800"}, {ID: "1764280800-2"}, {ID: "restore-5"}}
	cases := []struct{ in, want string }{
		{"1764280801", "1764280801"},   // free: kept as it is
		{"1764280800", "1764280800-3"}, // taken twice: the next free one
		{"restore-5", "restore-5-2"},
	}
	for _, c := range cases {
		if got := uniqueHistoryID(c.in, h); got != c.want {
			t.Errorf("uniqueHistoryID(%q) = %q, want %q", c.in, got, c.want)
		}
	}
	if got := uniqueHistoryID("", nil); got == "" {
		t.Error("an empty ID was left empty")
	}
}

// Two Backup Sets that start in the same second each keep their own entry in
// Reports (#173), and the first keeps its ID as it was.
func TestHistorySameSecondRunsGetTheirOwnIDs(t *testing.T) {
	useTempConfigDir(t)
	a := &App{}
	for _, name := range []string{"FILES_A", "FILES_B", "FILES_C"} {
		if err := a.appendJobHistory(JobHistory{ID: "1764280800", Name: name, Status: "success"}); err != nil {
			t.Fatal(err)
		}
	}
	history, err := a.GetJobHistory()
	if err != nil {
		t.Fatal(err)
	}
	if len(history) != 3 {
		t.Fatalf("history has %d entries, want 3", len(history))
	}
	want := map[string]string{"FILES_A": "1764280800", "FILES_B": "1764280800-2", "FILES_C": "1764280800-3"}
	for _, h := range history {
		if h.ID != want[h.Name] {
			t.Errorf("%s has ID %q, want %q", h.Name, h.ID, want[h.Name])
		}
	}
}
