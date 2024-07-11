// Partially separated `x-data` for `templates/tests.gohtml`.
function tests() {
    return {
        selected: new Set(),
        selectedAll: false,

        selectAll() {
            this.tests.forEach((t) => this.selected.add(t))
        },

        unselectAll() {
            this.selected.clear()
        },

        toggleSelectAll() {
            if (this.selectedAll) {
                this.unselectAll()
            } else {
                this.selectAll()
            }
 
            this.selectedAll = !this.selectedAll
        },

        toggleEntry(entry) {
            if (this.selected.has(entry)) {
                this.selected.delete(entry)
            } else {
                this.selected.add(entry)
            }
        },

        downloadSelectedAction() {
            const query = new URLSearchParams()
            for (const s of this.selected) {
                query.append('tests', s)
            }
            return `/teacher/tests/selected?${query.toString()}`
        }
    }
}
