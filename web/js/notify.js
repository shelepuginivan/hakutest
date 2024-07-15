function notificationHandler() {
    return {
        notifications: [],
        notify(params) {
            const id = `notification:${Date.now()}`
            const notification = {
                id,
                type: params.type,
                message: params.message,
                show: true,
            }

            this.notifications.push(notification)

            setTimeout(() => {
                const index = this.notifications.findIndex((n) => n.id === id)

                if (index > -1) {
                    this.notifications[index].show = false
                }
            }, 2000)
        }
    }
}
