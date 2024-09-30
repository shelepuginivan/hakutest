const navShortcuts = {
    "1": "/teacher/dashboard",
    "2": "/teacher/tests",
    "3": "/teacher/statistics",
    "4": "/teacher/settings",
}

document.addEventListener("keydown", function(event) {
    const newPath = navShortcuts[event.key]

    if (!newPath) {
        return
    }

    if (document.location.pathname !== newPath) {
        document.location.pathname = newPath
    }
})
