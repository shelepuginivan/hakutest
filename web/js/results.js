// Partially separated `x-data` for `templates/statistics_menu.gohtml`.
function statisticsMenu() {
    return {
        selected: new Set(),
        selectedAll: false,

        selectAll() {
            this.results.forEach((t) => this.selected.add(t))
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
    }
}
