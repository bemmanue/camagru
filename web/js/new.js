$(document).ready(function() {
    const input = document.getElementById("input")
    const canvas = document.getElementById("canvas")
    const button = document.getElementById("button")

    button.disabled = true

    const ctx = canvas.getContext("2d")
    const img = new Image()


    button.addEventListener("click", function () {
        const multipart = new FormData();
        multipart.append('file', dataURItoBlob(canvas.toDataURL()));

        $.ajax({
            type: "POST",
            url: "/new",
            data: multipart,
            contentType: false,
            processData: false,
            success: function () {
                location.replace("/profile")
            },
            error: function () {
            }
        })
    })

    input.addEventListener("change", function () {
        img.src = URL.createObjectURL(this.files[0]);

        img.onload = function() {
            canvas.width = this.naturalWidth;
            canvas.height = this.naturalHeight;
            ctx.drawImage(this, 0, 0);
        };

        if (button.classList.contains("inactive")) {
            button.classList.remove("inactive")
            button.classList.add("active")
        }
        button.disabled = false
    })


    const ranges = document.getElementsByClassName("range")
    const brightness = document.querySelector("#brightness")
    const contrast = document.querySelector("#contrast")
    const grayscale = document.querySelector("#grayscale")
    const saturate = document.querySelector("#saturate")
    const sepia = document.querySelector("#sepia")

    Array.from(ranges).forEach(function(range) {
        range.addEventListener("input", () => {
            ctx.filter = `
            brightness(${brightness.value}) 
            contrast(${contrast.value}) 
            grayscale(${grayscale.value}) 
            saturate(${saturate.value})
            sepia(${sepia.value})
            `
            ctx.drawImage(img, 0, 0);
        })
    });
});

function dataURItoBlob(dataURI) {
    let byteString;
    if (dataURI.split(',')[0].indexOf('base64') >= 0) {
        byteString = atob(dataURI.split(',')[1])
    } else {
        byteString = unescape(dataURI.split(',')[1])
    }

    const mimeString = dataURI.split(',')[0].split(':')[1].split(';')[0]

    const ia = new Uint8Array(byteString.length);
    for (let i = 0; i < byteString.length; i++) {
        ia[i] = byteString.charCodeAt(i);
    }

    return new Blob([ia], {type:mimeString});
}
