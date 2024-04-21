document.getElementById("search-form").onsubmit = (event) => {
    event.preventDefault()
    document.location.href = event.target.q.value
}
