export interface Note {
    id: number
    text: string
}

export async function getNotes(): Promise<Note[]> {
    const res = await fetch('/api/notes')

    if (!res.ok) {
        throw new Error(`Failed to fetch notes: ${res.statusText}`)
    }

    const data = await res.json()
    return Array.isArray(data) ? data : []
}

export async function createNote(text: string): Promise<Note> {
    const res = await fetch('/api/notes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text })
    })

    if (!res.ok) {
        throw new Error(`Failed to create note: ${res.statusText}`)
    }

    const data = await res.json()

    if (!data || typeof data.id !== 'number' || typeof data.text !== 'string') {
        throw new Error('Invalid note returned from server')
    }

    return data
}
