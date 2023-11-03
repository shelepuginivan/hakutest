const tasksSection = document.getElementById("tasks")
const buttonAddTask = document.getElementById("button-add-task")
const numberOfTasksInput = document.getElementById("number-of-tasks-input")

let taskIndex = 0

const addTask = (taskIndex) => {
    const newTask = document.createElement('fieldset')

    newTask.classList.add("test-task", "task-${taskIndex}")
    newTask.innerHTML = `
        <legend>Task ${taskIndex + 1}</legend>
        <select class="input-select" name="${taskIndex}-type">
            <option value="single">Single answer</option>
            <option value="multiple">Multiple answers</option>
            <option value="open">Open question</option>
        </select>
        <input
            class="input-text"
            type="text"
            name="${taskIndex}-text"
        >
        <input
            class="input-text"
            type="text"
            name="${taskIndex}-answer"
        >
        <div>
            <input
                class="input-checkbox"
                type="checkbox"
                name="${taskIndex}-has-attachment"
                onchange="toggleAttachment(${taskIndex}, this)"
            >
            <div
                id="${taskIndex}-attachment"
                class="attachment"
                data-enabled="false"
            ></div>
            <div>
                <div id="${taskIndex}-options"></div>
                <button
                    class="button-add-option"
                    type="button"
                    onclick="addOption(${taskIndex})"
                >
                    + Add option
                </button>
            </div>
        </div>`

    tasksSection.appendChild(newTask)
}

const addAttachment = (taskIndex) => {
    const attachment = document.getElementById(`${taskIndex}-attachment`)
    const attachmentFields = `
        <input
            class="input-text"
            type="text"
            name="${taskIndex}-attachment-name"
        >
        <select class="input-select" name="${taskIndex}-attachment-type">
            <option value="file">File</option>
            <option value="image">Image</option>
            <option value="video">Video</option>
            <option value="audio">Audio</option>
        </select>
        <input
            class="input-text"
            type="text"
            name="${taskIndex}-attachment-src"
        >`

    attachment.innerHTML = attachmentFields
    attachment.dataset.enabled = true
}

const removeAttachment = (taskIndex) => {
    const attachment = document.getElementById(`${taskIndex}-attachment`)
    attachment.innerHTML = ""
    attachment.dataset.enabled = false
}

const toggleAttachment = (taskIndex, checkbox) => {
    if (!checkbox.checked) {
        return removeAttachment(taskIndex)
    }

    addAttachment(taskIndex)
}

const addOption = (taskIndex) => {
    const taskOptions = document.getElementById(`${taskIndex}-options`)
    const newOption = document.createElement('input')

    newOption.className = "input-text"
    newOption.type = "text"
    newOption.name = `${taskIndex}-options`

    taskOptions.appendChild(newOption)
}

buttonAddTask.onclick = () => {
    addTask(taskIndex)
    numberOfTasksInput.value = ++taskIndex
}
