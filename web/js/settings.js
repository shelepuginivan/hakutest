function settings() {
    return {
        onSubmit() {
            // NOTE: Alpine converts number to string.
            this.general.port = parseInt(this.general.port) || this.general.port

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
