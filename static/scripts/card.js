function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

let btnYes = document.getElementById("btn-yes")
let btnNo = document.getElementById("btn-no")
let gif = document.getElementById("gif")
let msg = document.getElementById("message")
let sender = document.getElementById("senderName")

btnYes.addEventListener("mouseenter", () => {
    gif.src = "static/resource/cat_excited.gif"
})
btnNo.addEventListener("mouseenter", () => {
    gif.src = "static/resource/cat_cry.gif"
})
btnYes.addEventListener("mouseleave", () => {
    gif.src = "static/resource/cat_heart.gif"
})
btnNo.addEventListener("mouseleave", () => {
    gif.src = "static/resource/cat_heart.gif"
})


btnYes.addEventListener("click", async() => {
    btnYes.style = "display: none"
    btnNo.style = "display: none"
    sender.style = "display: none"
    msg.textContent = "YAY THANK YOU!!! <333"
    await sleep(30)
    gif.src = "static/resource/cat_success.gif"

    let form = {
        "id": document.getElementById("card-id").textContent,
        "clickyesfirst": "yes"

    }

    let encoded = new URLSearchParams(form)

    fetch("card", {
        method: 'POST', 
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: encoded
    }).then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json(); // Or .text() if the server response is not JSON
        })
        .then(data => {
        console.log('Success:', data);
        })
        .catch(error => {
        console.error('Error:', error);
        });
})

var clickedNo = 100
var no = 0
var NoMessage = ["Wait", "Are you sure?", "please?", "No?", ":(", "last chance", "3", "2", "1"]
btnNo.addEventListener("click", async() => {
    if (no >= 9) {
        btnYes.style = "display: none"
        btnNo.style = "display: none"
        sender.style = "display: none"
        msg.textContent = "NOOOOO </3"
        await sleep(30)
        gif.src = "static/resource/cat_wiping.gif"

        let form = {
        "id": document.getElementById("card-id").textContent,
        "clickyesfirst": "no"

    }

    let encoded = new URLSearchParams(form)

    fetch("card", {
        method: 'POST', 
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: encoded
    }).then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json(); // Or .text() if the server response is not JSON
        })
        .then(data => {
        console.log('Success:', data);
        })
        .catch(error => {
        console.error('Error:', error);
        });
    } else {
        btnNo.textContent = NoMessage[no]
        btnNo.style = `width: ${clickedNo}px; height: ${clickedNo-30}px` 
    }
    clickedNo = clickedNo + 100
    no++
})
