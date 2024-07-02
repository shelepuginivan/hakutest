const form = document.getElementById('test-form')
const modalWrapper = document.querySelector('.modal__wrapper')
const modalBackdrop = document.querySelector('.modal__backdrop')
const modalConfirm = document.querySelector('.modal__confirm')
const modalClose = document.querySelector('.modal__close')

function cancel() {
    document.body.classList.remove('modal_open')
    modalWrapper.classList.remove('visible')
    modalBackdrop.classList.remove('visible')
}

function submit() {
    form.submit()
}

form.addEventListener('submit', function(event) {
    event.preventDefault()
    document.body.classList.add('modal_open')
    modalWrapper.classList.add('visible')
    modalBackdrop.classList.add('visible')
})

modalConfirm.addEventListener('click', submit)
modalBackdrop.addEventListener('click', cancel)
modalClose.addEventListener('click', cancel)

document.body.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        cancel()
    }
})
