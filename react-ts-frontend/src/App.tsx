import React, { useEffect, useState } from 'react'
import { getNotes, createNote, type Note } from './api'


export default function App() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [text, setText] = useState<string>('')

  useEffect(() => {
    loadNotes()
  }, [])

  async function loadNotes() {
    const n = await getNotes()
    setNotes(n)
  }

  async function handleAdd(e: React.FormEvent) {
    e.preventDefault()
    if (!text.trim()) return 
    const created = await createNote(text.trim())
    setNotes((s) => [...s, created])
    setText('')
  }

  return (
    <div style={{ fontFamily: 'system-ui, sans-serif', padding: 24 }}>
    <h1>React + Go Notes</h1>

    <section style={{ marginTop: 20 }}>
      <h2>Notes</h2>
      <form onSubmit={handleAdd}>
      <input
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Write a note"
        style={{ padding: 8, width: 300 }}
        />
          <button type="submit" style={{ marginLeft: 8, padding: '8px 12px' }}>
          Add
          </button>
      </form>


        <ul>
          {notes?.map(n => (
            <li key={n.id}>{n.text}</li>
          ))}
        </ul>
      </section>
    </div>
  )
}