// Keyboard hotkeys for Markdown editor.
function hotkey(event) {
    if (!event.ctrlKey) {
        return
    }

    switch (event.code) {
        case 'KeyB':    // Bold.
            event.preventDefault()
            applyFormatting(event.target, "**", "**", -2, -2)
            break
        case 'KeyI':    // Italic.
            event.preventDefault()
            applyFormatting(event.target, "*", "*", -1, -1)
            break
        case 'KeyK':    // Link.
            event.preventDefault()
            applyFormatting(event.target, "[", "]()", -3, -1)
            break
    }

    if (!event.shiftKey) {
        return
    }
    
    switch (event.code) {
        case 'KeyX':    // Strikethrough.
            event.preventDefault()
            applyFormatting(event.target, "~~", "~~", -2, -2)
            break
        case 'KeyM':    // Monospace.
            event.preventDefault()
            applyFormatting(event.target, "`", "`", -1, -1)
            break
    }
}

function applyFormatting(textarea, openToken, closeToken, moveCursorNormal, moveCursorSelection) {
    const text = textarea.value
    const start = textarea.selectionStart
    const end = textarea.selectionEnd

    const before = text.substring(0, start)
    const selection = text.substring(start, end)
    const after = text.substring(end)

    textarea.value = `${before}${openToken}${selection}${closeToken}${after}`

    // Move cursor according to the mode.
    if (start === end) {
        // Normal mode.
        textarea.selectionEnd += moveCursorNormal
    } else {
        // Selection mode.
        textarea.selectionEnd += moveCursorSelection
    }
    
    textarea.selectionStart = textarea.selectionEnd
}
