import React, { useEffect, useState } from 'react'
import { getNotes, createNote, deleteNote, type Note } from './api'

export default function App() {
  const [notes, setNotes] = useState<Note[]>([])
  const [text, setText] = useState('')
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    loadNotes()
  }, [])

  async function loadNotes() {
    setLoading(true)
    try {
      const n = await getNotes()
      setNotes(n)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  async function handleAdd(e: React.FormEvent) {
    e.preventDefault()
    if (!text.trim()) return
    try {
      const created = await createNote(text.trim())
      setNotes((s) => [created, ...s])
      setText('')
    } catch (err) {
      console.error(err)
    }
  }

  async function handleDelete(id: number) {
    if (!confirm('Delete this note?')) return
    try {
      await deleteNote(id)
      setNotes((s) => s.filter((n) => n.id !== id))
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div style={{ fontFamily: 'system-ui, sans-serif', padding: 24 }}>
      <h1>📝 React + Go Notes</h1>

      <form onSubmit={handleAdd} style={{ marginTop: 16 }}>
        <input
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Write a note..."
          style={{ padding: 8, width: 300 }}
        />
        <button
          type="submit"
          style={{ marginLeft: 8, padding: '8px 12px' }}
          disabled={!text.trim()}
        >
          Add
        </button>
      </form>

      <section style={{ marginTop: 32 }}>
        <h2>Notes</h2>
        {loading ? (
          <p>Loading...</p>
        ) : notes.length === 0 ? (
          <p>No notes yet.</p>
        ) : (
          <ul style={{ listStyle: 'none', padding: 0 }}>
            {notes.map((n) => (
              <li
                key={n.id}
                style={{
                  padding: 12,
                  marginBottom: 8,
                  background: '#3C3C3C',
                  borderRadius: 8,
                  display: 'flex',
                  justifyContent: 'space-between',
                  alignItems: 'center',
                }}
              >
                <div>
                  <div>{n.text}</div>
                  <small style={{ color: '#666' }}>
                    {new Date(n.created_at).toLocaleString()}
                  </small>
                </div>
                <button
                  onClick={() => handleDelete(n.id)}
                  style={{
                    background: '#e74c3c',
                    color: 'white',
                    border: 'none',
                    borderRadius: 4,
                    padding: '6px 10px',
                    cursor: 'pointer',
                  }}
                >
                  Delete
                </button>
              </li>
            ))}
          </ul>
        )}
      </section>
    </div>
  )
}
