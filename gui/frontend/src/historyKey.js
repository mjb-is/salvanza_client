// historyKey tells Reports entries apart. New entries have an ID of their own,
// but older ones from two Backup Sets that started in the same second share
// one (#173), so the time the run ended goes with it.
export const historyKey = (entry) => (entry ? `${entry.id}|${entry.timestamp || ''}` : '')

// findHistory returns the entry with that key, or the newest one when none has
// it (nothing picked yet, or the picked one has dropped out of the list).
export const findHistory = (history, key) =>
  (history || []).find(h => historyKey(h) === key) || (history && history[0]) || null
