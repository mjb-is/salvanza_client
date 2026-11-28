// node src/historyKey.test.mjs
import assert from 'node:assert/strict'
import { historyKey, findHistory } from './historyKey.js'

// Two older entries with the same ID (two sets that started in the same
// second, #173) and a newer one with its own.
const history = [
  { id: '1764280800-2', name: 'FILES_C', timestamp: '2026-11-27T23:05:00Z' },
  { id: '1764280800', name: 'FILES_A', timestamp: '2026-11-27T23:04:17Z' },
  { id: '1764280800', name: 'FILES_B', timestamp: '2026-11-27T23:00:45Z' },
]
const keys = history.map(historyKey)
assert.equal(new Set(keys).size, 3, 'every entry has its own key')

// Each one is found by its own key, so only it shows as selected.
for (const h of history) assert.equal(findHistory(history, historyKey(h)), h)
assert.equal(history.filter(h => historyKey(h) === keys[2]).length, 1)

// Nothing picked, or a key no longer in the list: the newest.
assert.equal(findHistory(history, null), history[0])
assert.equal(findHistory(history, 'gone|2026-01-01T00:00:00Z'), history[0])
assert.equal(findHistory([], null), null)
assert.equal(findHistory(undefined, null), null)

// An entry without a time still has a key.
assert.equal(historyKey({ id: 'x' }), 'x|')
assert.equal(historyKey(null), '')
console.log('historyKey: ok')
