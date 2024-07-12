function editor() {
    return {
        addTask() {
            this.tasks.push({
                id: `task:${Date.now()}`,
                type: 'single',
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
            }).finally(function() {
                window.location.href = '/teacher/tests'
            })
        }
    }
}
