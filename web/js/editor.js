function editor() {
    return {
        addTask(type) {
            return function() {
                this.tasks.push({
                    id: `task:${Date.now()}`,
                    type,
                    text: '',
                    single: {
                        answer: '',
                        options: []
                    },
                    multiple: {
                        answer: [],
                        options: []
                    },
                    open: ''
                })
            }
        },

        removeOptionMultiple(task, option) {
            this.tasks[task].multiple.options.splice(option, 1)

            const answerIndex = this.tasks[task].multiple.answer.indexOf(option.toString())

            if (answerIndex > -1) {
                this.tasks[task].multiple.answer.splice(answerIndex, 1)
            }
        },

        // Submit edited test.
        onSubmit() {
            const body = {
                title: this.title,
                description: this.description,
                subject: this.subject,
                author: this.author,
                target: this.target,
                institution: this.institution,
                shuffleTasks: this.shuffleTasks,
                tasks: [],
            }

            if (this.expiresAt) {
                body.expiresAt = new Date(this.expiresAt).toISOString()
            }

            for (task of this.tasks) {
                const bodyTask = {
                    type: task.type,
                    text: task.text,
                }

                switch (task.type) {
                    case 'single':
                        bodyTask.answer = task.single.answer
                        bodyTask.options = task.single.options
                        break
                    case 'multiple':
                        bodyTask.answer = task.multiple.answer.sort().join(',')
                        bodyTask.options = task.multiple.options
                        break
                    case 'open':
                        bodyTask.answer = task.open
                        break
                    default:
                        bodyTask.answer = ''
                        break
                }

                body.tasks.push(bodyTask)
            }

            fetch(document.location.href, {
                method: 'POST',
                body: JSON.stringify(body)
            })
                .then((response) => {
                    if (response.ok) {
                        window.location.href = '/teacher/tests'
                    }
                    return Promise.reject(response)
                })
                .catch(function(res) {
                    res.json().then(function(json) {
                        dispatchEvent(new CustomEvent('notify', {
                            detail: {
                                type: 'error',
                                message: json.message,
                            }
                        }))
                    })
                })
        }
    }
}
