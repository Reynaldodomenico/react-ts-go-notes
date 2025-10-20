export interface Note {
    id: number
    text: string
}

export async function getNotes(): Promise<Note[]> {
    const res = await fetch('/api/notes')
    return res.json()
}

export async function createNote(text: string): Promise<Note> {
        const res = await fetch('/api/notes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text })
    })
    return res.json()
}