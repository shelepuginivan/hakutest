const tasksSection = document.getElementById("tasks")
const buttonAddTask = document.getElementById("button-add-task")
const numberOfTasksInput = document.getElementById("number-of-tasks-input")

let taskIndex = Number(numberOfTasksInput.value)

const addTask = (taskIndex) => {
    const newTask = document.createElement('fieldset')

    newTask.classList.add("test-task", "task-${taskIndex}")
    newTask.innerHTML = `
        <legend>Task ${taskIndex + 1}</legend>
        <div class="input-wrapper">
            <label for="${taskIndex}-type">
                Type:
            </label>
            <select
                id="${taskIndex}-type"
                class="input-select"
                name="${taskIndex}-type"
            >
                <option value="single">Single answer</option>
                <option value="multiple">Multiple answers</option>
                <option value="open">Open question</option>
            </select>
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-text">
                Text:
            </label>
            <input
                id="${taskIndex}-text"
                class="input-text"
                type="text"
                name="${taskIndex}-text"
            >
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-answer">
                Answer:
            </label>
            <input
                class="input-text"
                type="text"
                name="${taskIndex}-answer"
            >
        </div>
        <div class="answer-options-wrapper">
            <p>Answer options:</p>
            <div class="answer-options" id="${taskIndex}-options"></div>
            <button
                class="button-add-option"
                type="button"
                onclick="addOption(${taskIndex})"
            >
                + Add option
            </button>
        </div>
        <div class="attachment-wrapper">
            <div>
                <input
                    id="${taskIndex}-has-attachment"
                    class="input-checkbox"
                    type="checkbox"
                    name="${taskIndex}-has-attachment"
                    onchange="toggleAttachment(${taskIndex}, this)"
                >
                <label for="${taskIndex}-has-attachment">
                    Add attachment
                </label>
            </div>
            <div
                id="${taskIndex}-attachment"
                class="attachment"
                data-enabled="false"
            ></div>
        </div>`

    tasksSection.appendChild(newTask)
}

const addAttachment = (taskIndex) => {
    const attachment = document.getElementById(`${taskIndex}-attachment`)
    const attachmentFields = `
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-name">
                Name:
            </label>
            <input
                id="${taskIndex}-attachment-name"
                class="input-text"
                type="text"
                name="${taskIndex}-attachment-name"
            >
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-type">
                Name:
            </label>
            <select
                id="${taskIndex}-attachment-type"
                class="input-select"
                name="${taskIndex}-attachment-type"
            >
                <option value="file">File</option>
                <option value="image">Image</option>
                <option value="video">Video</option>
                <option value="audio">Audio</option>
            </select>
        </div>
        <div class="input-wrapper">
            <label for="${taskIndex}-attachment-src">
                Name:
            </label>
            <input
                id="${taskIndex}-attachment-src"
                class="input-text"
                type="text"
                name="${taskIndex}-attachment-src"
            >
        </div>`

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
