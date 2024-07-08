// Partially separated `x-data` for `templates/statistics.gohtml`.
function statistics() {
    return {
        sortedBy: 'student',
        ascending: true,

        sortByStudent() {
            if (this.sortedBy === 'student') {
                this.ascending = !this.ascending
            }
            this.sortedBy = 'student'

            this.results.sort((a, b) => {
                if (this.ascending) {
                    return a.student >= b.student
                }
                return a.student <= b.student
            })
        },

        sortByPoints() {
            if (this.sortedBy === 'points') {
                this.ascending = !this.ascending
            }
            this.sortedBy = 'points'

            this.results.sort((a, b) => {
                if (this.ascending) {
                    return a.points >= b.points
                }
                return a.points <= b.points
            })
        },

        sortByPercentage() {
            if (this.sortedBy === 'percentage') {
                this.ascending = !this.ascending
            }
            this.sortedBy = 'percentage'

            this.results.sort((a, b) => {
                if (this.ascending) {
                    return a.percentage >= b.percentage
                }
                return a.percentage <= b.percentage
            })
        },

        sortBySubmissionTime() {
            if (this.sortedBy === 'submittedAt') {
                this.ascending = !this.ascending
            }
            this.sortedBy = 'submittedAt'

            this.results.sort((a, b) => {
                if (this.ascending) {
                    return a.submittedAt >= b.submittedAt
                }
                return a.submittedAt <= b.submittedAt
            })
        },

        sortByTask(task) {
            if (this.sortedBy === `task:${task}`) {
                this.ascending = !this.ascending
            }
            this.sortedBy = `task:${task}`

            this.results.sort((a, b) => {
                if (this.ascending) {
                    return a.answers[task].correct >= b.answers[task].correct
                }
                return a.answers[task].correct <= b.answers[task].correct
            })
        },

        sortIndicator() {
            const iconName = this.ascending ? 'i arrow-up' : 'i arrow-down'
            return `sort__indicator ${iconName}`
        }
    }
}
