let senderName = document.getElementById("senderName")
let message = document.getElementById("message")

let = ReceipientName = document.getElementById("receipient")
let = SenderName = document.getElementById("sender")

ReceipientName.addEventListener("input", (e) => {
    message.textContent = `${e.target.value}, will you be my Valentine?`
})
SenderName.addEventListener("input", (e) => {
    senderName.textContent = `-${e.target.value}`
})