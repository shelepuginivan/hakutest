const form = document.getElementById('test-form')
const modal = document.querySelector('.modal')
const modalBackdrop = document.querySelector('.modal__backdrop')

function cancel() {
    document.body.classList.remove('modal_open')
    modalBackdrop.classList.remove('visible')
}

function submit() {
    form.submit()
}

form.addEventListener('submit', function(event) {
    event.preventDefault()
    document.body.classList.add('modal_open')
    modalBackdrop.classList.add('visible')
})

modal.addEventListener('click', function(event) {
    event.stopPropagation()
})

document.body.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        cancel()
    }
})

Array.from(document.querySelectorAll('[data-cancel]')).forEach(function(el) {
    el.addEventListener('click', cancel)
})

Array.from(document.querySelectorAll('[data-confirm]')).forEach(function(el) {
    el.addEventListener('click', submit)
})
