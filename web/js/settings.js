function settings() {
    return {
        onSubmit() {
            fetch(document.location.href, {
                method: "POST",
                body: JSON.stringify(this)
            })
                .then((response) => {
                    if (response.ok) {
                        location.reload()
                        return
                    }
                    return Promise.reject(response)
                })
                .catch(function(res) {
                    res.json().then(function(json) {
                        notify("error", json.message)
                    })
                })
        }
    }
}
